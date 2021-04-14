using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Runtime.Serialization;
using System.IO;
using System.ServiceModel;
using System.Configuration;
using ICSharpCode.SharpZipLib;
using ICSharpCode.SharpZipLib.Checksums;
using ICSharpCode.SharpZipLib.Zip;
using ICSharpCode.SharpZipLib.BZip2;
using ICSharpCode.SharpZipLib.Zip.Compression.Streams;
using ICSharpCode.SharpZipLib.Zip.Compression;
using System.Reflection;
using System.Xml;

#if SILVERLIGHT
namespace SharpTop.Utility {
#else
namespace Dayforce.Common {
#endif

    /// <summary>
    /// This class provides utility methods to support serialization and deserialization with compression.
    /// </summary>
    public class Compressor {

        /// <summary>
        /// EncodingType. An enumeration specifying the types of XML encoding we support.
        /// </summary>
        /// 
        public enum EncodingType {
            PLAIN_XML = 0,                      // No encoding, text XML
            BINARY_XML = 1,                     // Binary XML Encoding
            BINARY_XML_DICTIONARY = 2,          // For Binary XML, use a shared dictionary with a session
            BINARY = 3                          // Binary payload. Contents are not XML and are determined by caller.

        }

        /// <summary>
        /// CompressionType. An enumeration specifying the types of compression we support.
        /// </summary>
        /// 
        public enum CompressionType {
            NO_COMPRESSION = 0,               // No compression
            DEFLATE = 1,                      // Deflate algorithm (level 9)
            BZIP2 = 2,                        // BZIP2 algorithm ( block size 9)
            DEFLATE_ARCHIVE = 3               // Legacy mode, deflate algorithm, wrapped in ZIP archive.
        }

        //
        // Encoding/Compression Format:
        // 
        // Field                    Length  Comments
        // CompressionType          1       Compression Type
        // Encoding Type            1       Encoding Type
        // XML Dictionary Length    var     Only present if EncodingType = BINARY_XML_DICTIONARY
        //                                  = Total Length of compressed dictionary (including length bytes)
        //                                  (Meaning you can position the stream to skip past this)
        // XML Dictionary           var     Only present if Encoding Type = BINARY_XML_DICTIONARY
        //                                  = Sequence of binary encoded strings, compressed with Deflate (see BinaryWriter)
        //                                  
        // Compressed Payload       var     Compressed payload, as per CompressionType
        //  
        // Note: PLAIN_XML + NO_COMPRESSION is very inefficient, as the base64 encoding bloats the payload.
        // Static initializer. Read our compression settings options from the AppSettings.

        static Compressor() {
#if !SILVERLIGHT
            string configValue = ConfigurationManager.AppSettings["compressionType"];
            if (!String.IsNullOrEmpty(configValue)) {
                try {
                    compressionType = (CompressionType)Enum.Parse(typeof(CompressionType), configValue);
                } catch (Exception ex) {
                    throw new ApplicationException("Invalid compression type specified in AppSettings: " + configValue, ex);
                }
            }

            configValue = ConfigurationManager.AppSettings["encodingType"];
            if (!String.IsNullOrEmpty(configValue)) {
                try {
                    encodingType = (EncodingType)Enum.Parse(typeof(EncodingType), configValue);
                } catch (Exception ex) {
                    throw new ApplicationException("Invalid encoding type specified in AppSettings: " + configValue, ex);
                }
            }
#endif
        }

        // Default compression type. There should be no reason to change this.
        private static CompressionType compressionType = CompressionType.DEFLATE;

        // Default encoding type. The default is BINARY_XML_DICTIONARY which has shown to 
        // reduce payload size by from 4-10% for the most part. (There are some cases where it 
        // doesn't help very much, but our compression of the session dictionary helps to ensure
        // that it doesn't make it worse).
        //
        // If the default is to be changed, change it to BINARY_XML.
        private static EncodingType encodingType = EncodingType.BINARY_XML_DICTIONARY;

        /// <summary>
        /// Set the compressors compression options. This is a static setting and controls the compressor
        /// globally (within this process).
        /// </summary>
        /// <param name="compressionOptions">The compression options to use.</param>
        public static void SetCompressionType(CompressionType compressionType) {
            Compressor.compressionType = compressionType;
        }

        /// <summary>
        /// Set the compressors encoding options. This is a static setting and controls the compressor
        /// globally (within this process).
        /// </summary>
        /// <param name="compressionOptions">The compression options to use.</param>
        public static void SetEncodingType(EncodingType encodingType) {
            if (encodingType == EncodingType.BINARY)
                throw new ArgumentException("The BINARY encoding type is for internal use only.");
            Compressor.encodingType = encodingType;
        }

        /// <summary>
        /// Given an arbitrary object, encode it using the DataContractSerializer, compress it and 
        /// return the compressed result as a byte array.
        /// </summary>
        /// <typeparam name="T">The type of the object to be compressed (which must be attributed with [DataContract]).</typeparam>
        /// <param name="obj">The object to be compressed.</param>
        /// <param name="knownTypes">An IEnumerable of known types that may appear in the object (optional).</param>
        /// <returns>The compressed object, as a byte array.</returns>
        public static Byte[] ObjectToXMLToCompressedBinary<T>(T obj, IEnumerable<Type> knownTypes = null) {
            return InternalObjectToXMLToBinary(obj, knownTypes);
        }

        /// <summary>
        /// Given an arbitrary object, encode it using the DataContractSerializer, compress it and 
        /// return the compressed result as a byte array.
        /// </summary>
        /// <param name="obj">The object to be compressed. (It's type must be attributed with [DataContract]).</param>
        /// <param name="knownTypes">An IEnumerable of known types that may appear in the object (optional).</param>
        /// <returns>The compressed object, as a byte array.</returns>
        public static Byte[] ObjectToXMLToCompressedBinary(object obj, IEnumerable<Type> knownTypes = null) {
            return InternalObjectToXMLToBinary(obj, knownTypes);
        }

        /// <summary>
        /// Given an arbitrary object, encode it using the DataContractSerializer and 
        /// return the compressed result as a byte array.
        /// </summary>
        /// <param name="obj">The object to be serialized. (It's type must be attributed with [DataContract]).</param>
        /// <param name="knownTypes">An IEnumerable of known types that may appear in the object (optional).</param>
        /// <returns>The compressed object, as a byte array.</returns>
        public static Byte[] ObjectToXMLToBinary(object obj, IEnumerable<Type> knownTypes = null) {
            return InternalObjectToXMLToBinary(obj, knownTypes, false);
        }

        /// <summary>
        /// Given an arbitrary object, encode it using the DataContractSerializer and 
        /// return the compressed result as a byte array.
        /// </summary>
        /// <typeparam name="T">The type of the object to be serialized (which must be attributed with [DataContract]).</typeparam>
        /// <param name="obj">The object to be compressed.</param>
        /// <param name="knownTypes">An IEnumerable of known types that may appear in the object (optional).</param>
        /// <returns>The compressed object, as a byte array.</returns>
        public static Byte[] ObjectToXMLToBinary<T>(T obj, IEnumerable<Type> knownTypes = null) {
            return InternalObjectToXMLToBinary(obj, knownTypes, false);
        }

        /// <summary>
        /// Serialize and compress an object to a byte array.
        /// Done in one operation as a stream chain.
        /// </summary>
        /// <param name="input">The object to serialize (which must be attributed with [DataContract]).</param>
        /// <param name="knownTypes">An IEnumerable of known types tha tmay appear in the object (optional).</param>
        /// <returns>The serialized and compressed object, as a byte[].</returns>
        private static Byte[] InternalObjectToXMLToBinary(object obj, IEnumerable<Type> knownTypes = null, bool compress = true) {

            // Make local copies of the encoding type and compression type. This is used if
            // we activate heuristic based encoding and compression type selection.

            EncodingType actualEncodingType = encodingType;
            CompressionType actualCompressionType;
            if (compress)
                actualCompressionType = compressionType;
            else
                actualCompressionType = CompressionType.NO_COMPRESSION;


            // A first try at heuristic selection of compression. Fail. 
            // Don't enable this, as the average payload size (in aggregate) will increase.
            //
            // Heuristic encoding type selection. 
            // 1. If the admin has chosen binary xml with dictionary, or text xml, then we use those formats.
            // 2. If the admin has chosen binary xml, we initially ignore this and set plain xml, unless the
            // object being serialized is a list with greater than 10 elements, in which case we use binary xml.
            //
            /*
            if (encodingType != EncodingType.BINARY_XML_DICTIONARY && encodingType != EncodingType.PLAIN_XML) {
                // Try and get the count of elements in the list ( if it's a list).

                int listSize = -1;

                try {
                    // Lists have count properties.
                    PropertyInfo pi = obj.GetType().GetProperty("Count");
                    if (pi != null) {
                        Object valObj = pi.GetValue(obj, null);
                        listSize = (Int32) valObj;
                    }
                } catch (Exception) {
                }

                if (listSize != -1) {
                    if (listSize > 10) {
                        actualEncodingType = EncodingType.PLAIN_XML;
                        compressionType = CompressionType.BZIP2;
                    } 
                }
            } 
            */

            DFWriterSession xmlSession = new DFWriterSession();

            using (MemoryStream buffer = new MemoryStream()) {

                // Set up the compressor (default compressor is none).

                using (Stream compressor =
                        (actualCompressionType == CompressionType.DEFLATE_ARCHIVE) ?
                            (Stream)(new ZipOutputStream(buffer))
                            : ((actualCompressionType == CompressionType.DEFLATE) ?
                                    (Stream)(new DeflaterOutputStream(buffer, new Deflater(9, true)))
                                    : ((actualCompressionType == CompressionType.BZIP2) ?
                                    (Stream)(new BZip2OutputStream(buffer, 9))
                                    : buffer))) {

                    // Set up the serializer and encoder.                
                    DataContractSerializer serializer = (knownTypes == null) ? new DataContractSerializer(obj.GetType())
                                                                             : new DataContractSerializer(obj.GetType(), knownTypes);
                    XmlDictionaryWriter dicWriter;

                    if (actualEncodingType == EncodingType.BINARY_XML) {
                        dicWriter = XmlDictionaryWriter.CreateBinaryWriter(compressor);
                    } else if (actualEncodingType == EncodingType.BINARY_XML_DICTIONARY) {
                        dicWriter = XmlDictionaryWriter.CreateBinaryWriter(compressor, null, xmlSession);
                    } else {
                        dicWriter = XmlDictionaryWriter.CreateTextWriter(compressor, Encoding.UTF8);
                    }

                    if (actualCompressionType == CompressionType.DEFLATE_ARCHIVE) {
                        ZipOutputStream zipper = compressor as ZipOutputStream;
                        zipper.SetLevel(9);
                        ZipEntry entry = new ZipEntry(string.Empty);
                        entry.DateTime = DateTime.SpecifyKind(DateTime.Now, DateTimeKind.Unspecified);
                        zipper.PutNextEntry(entry);
                        serializer.WriteObject(dicWriter, obj);
                        dicWriter.Flush();
                        zipper.Flush();
                        zipper.Finish();
                    } else if (actualCompressionType == CompressionType.DEFLATE) {
                        DeflaterOutputStream deflator = compressor as DeflaterOutputStream;
                        deflator.IsStreamOwner = false;
                        serializer.WriteObject(dicWriter, obj);
                        dicWriter.Flush();
                        deflator.Close();
                    } else if (actualCompressionType == CompressionType.BZIP2) {
                        BZip2OutputStream zipper = compressor as BZip2OutputStream;
                        zipper.IsStreamOwner = false;
                        serializer.WriteObject(dicWriter, obj);
                        dicWriter.Flush();
                        zipper.Close();
                    } else if (actualCompressionType == CompressionType.NO_COMPRESSION) {
                        serializer.WriteObject(dicWriter, obj);
                        dicWriter.Flush();
                        compressor.Flush();
                    } else {
                        throw new Exception("Compressor: Invalid compression options: " + actualCompressionType);
                    }

                    // Compression and encoding are complete. Put in the lead bytes, the dictionary (if used) and return the result.

                    byte[] content;
                    if (actualEncodingType == EncodingType.BINARY_XML_DICTIONARY) {
                        MemoryStream tempMemStream = new MemoryStream();
                        xmlSession.SaveDictionaryStrings(tempMemStream);
                        content = new byte[tempMemStream.Length + buffer.Length + 2];
                        tempMemStream.Seek(0, SeekOrigin.Begin);
                        buffer.Seek(0, SeekOrigin.Begin);
                        tempMemStream.Read(content, 2, (int)tempMemStream.Length);
                        buffer.Read(content, 2 + (int)tempMemStream.Length, (int)buffer.Length);
                        tempMemStream.Close();
                    } else {
                        content = new byte[buffer.Length + 2];
                        buffer.Seek(0, SeekOrigin.Begin);
                        buffer.Read(content, 2, (int)buffer.Length);
                    }

                    content[0] = (byte)actualCompressionType;
                    content[1] = (byte)actualEncodingType;

                    return content;
                }
            }
        }

        /// <summary>
        /// Compresses an Xml document (in a string) with PLAIN_XML encoding and return the
        /// result as a base64 encoded string,suitable for including in a SOAP envelope.
        /// This method is used by performance test code.
        /// </summary>
        /// <param name="xml">The XML to compress.</param>
        /// <returns>The compressed XML, as a base64 string.</returns>
        public static string CompressXmlToBase64String( string xml) {

            EncodingType actualEncodingType = EncodingType.PLAIN_XML;
            CompressionType actualCompressionType = compressionType;

            using (MemoryStream buffer = new MemoryStream()) {

                // Set up the compressor (default compressor is none).

                using (Stream compressor =
                        (actualCompressionType == CompressionType.DEFLATE_ARCHIVE) ?
                            (Stream)(new ZipOutputStream(buffer))
                            : ((actualCompressionType == CompressionType.DEFLATE) ?
                                    (Stream)(new DeflaterOutputStream(buffer, new Deflater(9, true)))
                                    : ((actualCompressionType == CompressionType.BZIP2) ?
                                    (Stream)(new BZip2OutputStream(buffer, 9))
                                    : buffer))) {

                    byte[] encodedXml = System.Text.Encoding.UTF8.GetBytes(xml);

                    if (actualCompressionType == CompressionType.DEFLATE_ARCHIVE) {
                        ZipOutputStream zipper = compressor as ZipOutputStream;
                        zipper.SetLevel(9);
                        ZipEntry entry = new ZipEntry(string.Empty);
                        entry.DateTime = DateTime.SpecifyKind(DateTime.Now, DateTimeKind.Unspecified);
                        zipper.PutNextEntry(entry);
                        zipper.Write(encodedXml, 0, encodedXml.Length);
                        zipper.Flush();
                        zipper.Finish();
                    } else if (actualCompressionType == CompressionType.DEFLATE) {
                        DeflaterOutputStream deflator = compressor as DeflaterOutputStream;
                        deflator.IsStreamOwner = false;
                        deflator.Write(encodedXml, 0, encodedXml.Length);
                        deflator.Close();
                    } else if (actualCompressionType == CompressionType.BZIP2) {
                        BZip2OutputStream zipper = compressor as BZip2OutputStream;
                        zipper.IsStreamOwner = false;
                        zipper.Write(encodedXml, 0, encodedXml.Length);
                        zipper.Close();
                    } else if (actualCompressionType == CompressionType.NO_COMPRESSION) {
                        compressor.Write(encodedXml, 0, encodedXml.Length);
                        compressor.Flush();
                    } else {
                        throw new Exception("Compressor: Invalid compression options: " + actualCompressionType);
                    }

                    // Compression and encoding are complete. Put in the lead bytes, the dictionary (if used) and return the result.

                    byte[] content;
                    content = new byte[buffer.Length + 2];
                    buffer.Seek(0, SeekOrigin.Begin);
                    buffer.Read(content, 2, (int)buffer.Length);
                    content[0] = (byte)actualCompressionType;
                    content[1] = (byte)actualEncodingType;

                    return System.Convert.ToBase64String(content);
                }
            }
        }

        /// <summary>
        /// Deserialize an object from it's XML representation.
        /// </summary>
        /// <typeparam name="T">The type of the object.</typeparam>
        /// <param name="objAsString">The object as an XML document, previously serialized with the DataContractSerializer.</param>
        /// <returns></returns>
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Naming", "CA1709:IdentifiersShouldBeCasedCorrectly", MessageId = "XML"), System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "CA1004:GenericMethodsShouldProvideTypeParameter")]
        public static T XMLToObject<T>(string objAsString) {

            DataContractSerializer serializer = new DataContractSerializer(typeof(T));
            using (StringReader backing = new StringReader(objAsString)) {
                XmlReaderSettings settings = new XmlReaderSettings();
                settings.CheckCharacters = false;
                return (T)serializer.ReadObject(XmlReader.Create(backing, settings));
            }
        }

#if !SILVERLIGHT
        /// <summary>
        /// Serialize an object to XML and return the result.
        /// Server-side only as this uses XmlTextWriter
        /// </summary>
        /// <param name="input">The object to serialize (which must be attributed with [DataContract]).</param>
        /// <returns>The serialized object as a string.</returns>
        public static string ObjectToXML(object obj) {
            // Set up the serializer and encoder.                
            DataContractSerializer serializer = new DataContractSerializer(obj.GetType());

            using (StringWriter backingString = new StringWriter()) {
                using (XmlTextWriter writer = new XmlTextWriter(backingString)) {
                    serializer.WriteObject(writer, obj);
                    return backingString.ToString();
                }
            }
        }
#endif

        /// <summary>
        /// Decompress and deserialize a binary payload into an object. This version expects to have our leads bytes in 
        /// the payload that identify the compression and encoding techniques used, but falls back to the old
        /// school technique if unrecognized.
        /// </summary>
        /// <typeparam name="T">The type of object expected.</typeparam>
        /// <param name="obj">The compressed binary of the object.</param>
        /// <param name="knownTypes">An IEnumerable of known types that may appear in the object (optional).</param>
        /// <returns>The decompressed and deserialized object.</returns>
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Naming", "CA1709:IdentifiersShouldBeCasedCorrectly", MessageId = "XML"), System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "CA1004:GenericMethodsShouldProvideTypeParameter")]
        public static T CompressedBinaryToXMLToObject<T>(byte[] obj, IEnumerable<Type> knownTypes = null) where T : class {
            return CompressedBinaryToXMLToObject(typeof(T), obj, knownTypes) as T;
        }

        /// <summary>
        /// Decompress and deserialize a binary payload into an object. This version expects to have our leads bytes in 
        /// the payload that identify the compression and encoding techniques used, but falls back to the old
        /// school technique if unrecognized.
        /// </summary>
        /// <typeparam name="T">The type of object expected.</typeparam>
        /// <param name="obj">The compressed binary of the object.</param>
        /// <param name="knownTypes">An IEnumerable of known types that may appear in the object (optional).</param>
        /// <returns>The decompressed and deserialized object.</returns>
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Naming", "CA1709:IdentifiersShouldBeCasedCorrectly", MessageId = "XML"), System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "CA1004:GenericMethodsShouldProvideTypeParameter")]
        public static object CompressedBinaryToXMLToObject( Type T, byte[] obj, IEnumerable<Type> knownTypes = null) {

            // Our result object. We'll return this.
            object resultObj = null;

            using (MemoryStream memStrm = new MemoryStream(obj)) {
                //
                // The compression options are the first byte of the stream.
                //
                CompressionType compressionType = (CompressionType)memStrm.ReadByte();
                EncodingType encodingType = (EncodingType)memStrm.ReadByte();

                bool recognized = true;
                if (compressionType != CompressionType.DEFLATE && compressionType != CompressionType.BZIP2 && compressionType != CompressionType.DEFLATE_ARCHIVE && compressionType != CompressionType.NO_COMPRESSION) {
                    recognized = false;
                }

                if (encodingType != EncodingType.BINARY_XML && encodingType != EncodingType.BINARY_XML_DICTIONARY && encodingType != EncodingType.PLAIN_XML) {
                    recognized = false;
                }

                if (!recognized) {
                    // The lead byte for a ZIP file (old technique) is not one of our valid lead bytes, so we're OK
                    // to assume that an unrecognized lead byte is probably a ZIP archive. If it's not, decompression 
                    // with fail anyway and throw.
                    memStrm.Seek(0, SeekOrigin.Begin);
                    compressionType = CompressionType.DEFLATE_ARCHIVE;
                    encodingType = EncodingType.PLAIN_XML;
                }

                XmlBinaryReaderSession xmlSession;
                if (encodingType == EncodingType.BINARY_XML_DICTIONARY) {
                    xmlSession = DFReaderSession.CreateReaderSession(memStrm);
                } else {
                    xmlSession = new XmlBinaryReaderSession();
                }

                // Construct the appropriate decompression stream.
                using (Stream decompressor =
                        (compressionType == CompressionType.DEFLATE_ARCHIVE) ?
                            (Stream)(new ZipInputStream(memStrm))
                            : ((compressionType == CompressionType.DEFLATE) ?
                                    (Stream)(new InflaterInputStream(memStrm, new Inflater(true)))
                                    : ((compressionType == CompressionType.BZIP2) ?
                                    (Stream)(new BZip2InputStream(memStrm))
                                    : memStrm))) {

                    // If this is a ZIP archive format, advance to the first entry
                    if (compressionType == CompressionType.DEFLATE_ARCHIVE) {
                        ((ZipInputStream)decompressor).GetNextEntry();
                    }

                    // Set up the serializer and decoder.                
                    DataContractSerializer serializer = (knownTypes == null) ? new DataContractSerializer(T)
                                                                             : new DataContractSerializer(T, knownTypes);
                    XmlDictionaryReader dicReader;

                    if (encodingType == EncodingType.BINARY_XML) {
                        dicReader = XmlDictionaryReader.CreateBinaryReader(decompressor, XmlDictionaryReaderQuotas.Max);
                        resultObj = serializer.ReadObject(dicReader);
                    } else if (encodingType == EncodingType.BINARY_XML_DICTIONARY) {
                        dicReader = XmlDictionaryReader.CreateBinaryReader(decompressor, null, XmlDictionaryReaderQuotas.Max, xmlSession);
                        resultObj = serializer.ReadObject(dicReader);
                    } else { // Text XML
                        using (StreamReader textStream = new StreamReader(decompressor, Encoding.UTF8)) {
                            XmlReaderSettings settings = new XmlReaderSettings();
                            settings.CheckCharacters = false;
                            resultObj = serializer.ReadObject(XmlReader.Create(textStream, settings));
                        }
                    }
                }
            }
            return resultObj;
        }

#if !SILVERLIGHT

        /// <summary>
        /// Decompress and deserialize a binary payload and return the actual XML. This is only used for diagnostics and perf testing
        /// (and for this reason absolutely must not be compiled onto the Silverlight client as its a security risk).
        /// Also, be careful about refactoring this to deal with the code that this method has in common
        /// with CompressedBinaryToXMLToObject<T> because that might lead to more code on the client that exposes
        /// our encoding technique (I recognize this is just obfuscation, but every layer of this helps). 
        /// 
        /// This version expects to have our leads bytes in the payload that identify the compression and encoding 
        /// techniques used, but falls back to the old school technique if unrecognized.
        /// </summary>
        /// <typeparam name="T">The type of object expected.</typeparam>
        /// <param name="obj">The compressed binary of the object.</param>
        /// <returns>The decompressed XML.</returns>
        public static string CompressedBinaryToXml(byte[] obj) {

            string result = "";

            using (MemoryStream memStrm = new MemoryStream(obj)) {
                //
                // The compression options are the first byte of the stream.
                //
                CompressionType compressionType = (CompressionType)memStrm.ReadByte();
                EncodingType encodingType = (EncodingType)memStrm.ReadByte();

                bool recognized = true;
                if (compressionType != CompressionType.DEFLATE && compressionType != CompressionType.BZIP2 && compressionType != CompressionType.DEFLATE_ARCHIVE) {
                    recognized = false;
                }

                if (encodingType != EncodingType.BINARY_XML && encodingType != EncodingType.BINARY_XML_DICTIONARY && encodingType != EncodingType.PLAIN_XML) {
                    recognized = false;
                }

                if (!recognized) {
                    // The lead byte for a ZIP file (old technique) is not one of our valid lead bytes, so we're OK
                    // to assume that an unrecognized lead byte is probably a ZIP archive. If it's not, decompression 
                    // with fail anyway and throw.
                    memStrm.Seek(0, SeekOrigin.Begin);
                    compressionType = CompressionType.DEFLATE_ARCHIVE;
                    encodingType = EncodingType.PLAIN_XML;
                }

                XmlBinaryReaderSession xmlSession;
                if (encodingType == EncodingType.BINARY_XML_DICTIONARY) {
                    xmlSession = DFReaderSession.CreateReaderSession(memStrm);
                } else {
                    xmlSession = new XmlBinaryReaderSession();
                }

                // Construct the appropriate decompression stream.
                using (Stream decompressor =
                        (compressionType == CompressionType.DEFLATE_ARCHIVE) ?
                            (Stream)(new ZipInputStream(memStrm))
                            : ((compressionType == CompressionType.DEFLATE) ?
                                    (Stream)(new InflaterInputStream(memStrm, new Inflater(true)))
                                    : ((compressionType == CompressionType.BZIP2) ?
                                    (Stream)(new BZip2InputStream(memStrm))
                                    : memStrm))) {

                    // If this is a ZIP archive format, advance to the first entry
                    if (compressionType == CompressionType.DEFLATE_ARCHIVE) {
                        ((ZipInputStream)decompressor).GetNextEntry();
                    }

                    // Set up the decoder.                
                    XmlDictionaryReader dicReader;

                    if (encodingType == EncodingType.BINARY_XML) {
                        dicReader = XmlDictionaryReader.CreateBinaryReader(decompressor, XmlDictionaryReaderQuotas.Max);
                        dicReader.MoveToStartElement();
                        result = dicReader.ReadOuterXml();
                    } else if (encodingType == EncodingType.BINARY_XML_DICTIONARY) {
                        dicReader = XmlDictionaryReader.CreateBinaryReader(decompressor, null, XmlDictionaryReaderQuotas.Max, xmlSession);
                        dicReader.MoveToStartElement();
                        result = dicReader.ReadOuterXml();
                    } else { // Text XML
                        using (StreamReader textStream = new StreamReader(decompressor, Encoding.UTF8))
                        {
                            XmlReaderSettings settings = new XmlReaderSettings();
                            settings.CheckCharacters = false;
                            result = textStream.ReadToEnd();
                        }
                    }
                }
            }

            return result;
        }

        /// Decompress and deserialize a binary payload (encoded in base64) and return the actual XML. 
        /// This is only used for diagnostics and perf testing (and for this reason absolutely must not be compiled 
        /// onto the Silverlight client as its a security risk).
        /// Also, be careful about refactoring this to deal with the code that this method has in common
        /// with CompressedBinaryToXMLToObject<T> because that might lead to more code on the client that exposes
        /// our encoding technique (I recognize this is just obfuscation, but every layer of this helps). 
        /// 
        /// This version expects to have our leads bytes in the payload that identify the compression and encoding 
        /// techniques used, but falls back to the old school technique if unrecognized.
        /// <param name="base64String">The compressed payload (a byte[]) encoded in base 64.</param>
        /// <returns>The Xml decompressed XML.</returns>
        public static string CompressedBinaryToXml(string base64String) {
            return CompressedBinaryToXml(System.Convert.FromBase64String(base64String));
        }

#endif

        /// <summary>
        /// Decompress a binary payload into a byte array. This version expects to have our leads bytes in 
        /// the payload that identify the compression and encoding techniques used, but falls back to the old
        /// school technique if unrecognized. You can safely pass a ZIP archive (in a byte array) and this
        /// method will retrieve the contents of the first entry in the archive (uncompressed).
        /// </summary>
        /// <param name="input">The compressed binary of the object.</param>
        /// <returns>The decompressed bytes.</returns>
        public static byte[] DecompressBytes(byte[] input) {

            byte[] result;

            using (MemoryStream memStrm = new MemoryStream(input)) {
                //
                // The compression and encoding options are the first 2 bytes of the stream.
                //
                CompressionType compressionType = (CompressionType)memStrm.ReadByte();
                EncodingType encodingType = (EncodingType)memStrm.ReadByte();

                bool recognized = true;
                if (compressionType != CompressionType.DEFLATE && compressionType != CompressionType.BZIP2 && compressionType != CompressionType.DEFLATE_ARCHIVE) {
                    recognized = false;
                }

                // We only support binary encoding (obviously, since we make no interpretation of the payload).
                if (encodingType != EncodingType.BINARY) {
                    recognized = false;
                }

                if (!recognized) {
                    // The lead byte for a ZIP file (old technique) is not one of our valid lead bytes, so we're OK
                    // to assume that an unrecognized lead byte is probably a ZIP archive. If it's not, decompression 
                    // with fail anyway and throw.
                    memStrm.Seek(0, SeekOrigin.Begin);
                    compressionType = CompressionType.DEFLATE_ARCHIVE;
                    encodingType = EncodingType.BINARY;
                }

                // Construct the appropriate decompression stream.
                using (Stream decompressor =
                        (compressionType == CompressionType.DEFLATE_ARCHIVE) ?
                            (Stream)(new ZipInputStream(memStrm))
                            : ((compressionType == CompressionType.DEFLATE) ?
                                    (Stream)(new InflaterInputStream(memStrm, new Inflater(true)))
                                    : ((compressionType == CompressionType.BZIP2) ?
                                    (Stream)(new BZip2InputStream(memStrm))
                                    : memStrm))) {

                    // If this is a ZIP archive format, advance to the first entry
                    if (compressionType == CompressionType.DEFLATE_ARCHIVE) {
                        ((ZipInputStream)decompressor).GetNextEntry();
                    }

                    // Decompress the array, storing the outputs in a new memory stream.
                    using (MemoryStream outputStream = new MemoryStream()) {
                        byte[] data = new byte[4096];

                        while (true) {
                            int size = decompressor.Read(data, 0, data.Length);
                            if (size > 0) {
                                outputStream.Write(data, 0, size);
                            } else
                                break;
                        }

                        // Copy the output stream to the response
                        outputStream.Flush();
                        result = new byte[outputStream.Length];
                        outputStream.Seek(0, SeekOrigin.Begin);
                        outputStream.Read(result, 0, (int)outputStream.Length);
                    }
                }
            }
            return result;
        }

        /// <summary>
        /// Compress a byte array. The current compression type is supported, with BINARY encoding (meaning no encoding).
        /// The compressed byte array will contain our standard lead bytes.
        /// </summary>
        /// <param name="input">The binary data to compress.</param>
        /// <returns>The serialized and compressed object, as a byte[].</returns>
        public static Byte[] CompressBytes(byte[] input) {

            // Make local copies of the encoding type and compression type. This is used if
            // we activate heuristic based compression type selection.

            EncodingType actualEncodingType = EncodingType.BINARY;
            CompressionType actualCompressionType = compressionType;

            using (MemoryStream buffer = new MemoryStream()) {

                // Set up the compressor (default compressor is none).

                using (Stream compressor =
                        (actualCompressionType == CompressionType.DEFLATE_ARCHIVE) ?
                            (Stream)(new ZipOutputStream(buffer))
                            : ((actualCompressionType == CompressionType.DEFLATE) ?
                                    (Stream)(new DeflaterOutputStream(buffer, new Deflater(9, true)))
                                    : ((actualCompressionType == CompressionType.BZIP2) ?
                                    (Stream)(new BZip2OutputStream(buffer, 9))
                                    : buffer))) {

                    if (actualCompressionType == CompressionType.DEFLATE_ARCHIVE) {
                        ZipOutputStream zipper = compressor as ZipOutputStream;
                        zipper.SetLevel(9);
                        ZipEntry entry = new ZipEntry(string.Empty);
                        entry.DateTime = DateTime.SpecifyKind(DateTime.Now, DateTimeKind.Unspecified);
                        zipper.PutNextEntry(entry);
                        compressor.Write(input, 0, input.Length);
                        compressor.Flush();
                        zipper.Flush();
                        zipper.Finish();
                    } else if (actualCompressionType == CompressionType.DEFLATE) {
                        DeflaterOutputStream deflator = compressor as DeflaterOutputStream;
                        deflator.IsStreamOwner = false;
                        deflator.Write(input, 0, input.Length);
                        deflator.Flush();
                        deflator.Close();
                    } else if (actualCompressionType == CompressionType.BZIP2) {
                        BZip2OutputStream zipper = compressor as BZip2OutputStream;
                        zipper.IsStreamOwner = false;
                        zipper.Write(input, 0, input.Length);
                        zipper.Flush();
                        zipper.Close();
                    } else if (actualCompressionType == CompressionType.NO_COMPRESSION) {
                        compressor.Write(input, 0, input.Length);
                        compressor.Flush();
                    } else {
                        throw new Exception("Compressor: Invalid compression options: " + actualCompressionType);
                    }

                    // Compression is complete. Put in the lead bytes and return the result.

                    byte[] content = new byte[buffer.Length + 2];
                    content[0] = (byte)actualCompressionType;
                    content[1] = (byte)actualEncodingType;
                    buffer.Seek(0, SeekOrigin.Begin);
                    buffer.Read(content, 2, (int)buffer.Length);

                    return content;
                }
            }
        }

        /// <summary>
        /// Compress a string to a byte array. 
        /// The string is not interpreted as XML (straight binary compression is used).
        /// </summary>
        /// <param name="input">The string to be compressed.</param>
        /// <returns>The compressed string as a byte[].</returns>
        public static byte[] CompressString(string input) {
            return CompressBytes(Encoding.UTF8.GetBytes(input));
        }

        /// <summary>
        /// Decompress a byte array and convert to a .Net string. The string is assumed not to be
        /// an Xml document, so no attempt at Xml deserialization will be attempted. (The string 
        /// must have been encoded as EncodingType.BINARY).
        /// </summary>
        /// <param name="input">The byte array containing the compressed string.</param>
        /// <returns>The decompressed string.</returns>
        public static string DecompressString(byte[] input) {
            byte[] theLoad = DecompressBytes(input);
            return Encoding.UTF8.GetString(theLoad, 0, theLoad.Length);
        }

        class DFWriterSession : XmlBinaryWriterSession {
            List<string> dictionaryStrings;
            public override bool TryAdd(XmlDictionaryString value, out int key) {
                if (dictionaryStrings == null)
                    dictionaryStrings = new List<string>();

                bool result = base.TryAdd(value, out key);
                if (result) {
                    this.dictionaryStrings.Add(value.Value);
                }

                return result;
            }

            public void SaveDictionaryStrings(Stream stream) {
                if (dictionaryStrings != null) {
                    // Create a deflation output stream and write each of the strings in the dictionary to it.
                    int totalLength = 0;
                    MemoryStream ms = new MemoryStream();
                    DeflaterOutputStream dos = new DeflaterOutputStream(ms, new Deflater(9, true));
                    dos.IsStreamOwner = false;
                    BinaryWriter binWriter = new BinaryWriter(dos, Encoding.UTF8);
                    foreach (string dicString in this.dictionaryStrings) {
                        binWriter.Write(dicString);
                        totalLength += dicString.Length;
                    }

                    dos.Close();

                    // Write the compressed length of the dictionary to the output stream.
                    BinaryWriter binWriter2 = new BinaryWriter(stream, Encoding.UTF8);
                    binWriter2.Write((int) ms.Position);

                    // Copy the compressed content to the output stream. 
                    // Avoid using CopyTo so that we can compile this in 3.5 (for Fiddler).
                    ms.Position = 0;
                    int bytesRead;
                    byte[] buffer = new byte[4096];
                    while ((bytesRead = ms.Read(buffer, 0, buffer.Length)) != 0) 
                        stream.Write(buffer, 0, bytesRead);
                    
                    ms.Close();

                } else {
                    // No dictionary. Just write a 0 (length).
                    BinaryWriter binWriter2 = new BinaryWriter(stream, Encoding.UTF8);
                    binWriter2.Write((int)0);
                }
            }
        }


        class DFReaderSession {

            public static XmlBinaryReaderSession CreateReaderSession(Stream stream) {
                XmlBinaryReaderSession result = new XmlBinaryReaderSession();

                // Read the length of the dictionary from the stream.
                BinaryReader reader = new BinaryReader(stream, Encoding.UTF8);
                int nextId = 0;
                int totalSessionLength = reader.ReadInt32();        // May be zero if no dictionary

                if (totalSessionLength > 0) {
                    //
                    // Decompress dictionary and store it in a byte array.
                    //
                    byte[] compressedDictionary = new byte[totalSessionLength];
                    stream.Read( compressedDictionary, 0, totalSessionLength);
                    InflaterInputStream iis = new InflaterInputStream(new MemoryStream(compressedDictionary), new Inflater(true));
                    MemoryStream uncompressedStream = new MemoryStream();

                    // Decompress and copy. Avoid CopyTo as it's .Net 4 and this file needs to be compiled for Fiddler.
                    byte[] buffer = new byte[2048];
                    int bytesRead;
                    while ((bytesRead = iis.Read(buffer, 0, buffer.Length)) != 0)
                        uncompressedStream.Write(buffer, 0, bytesRead);

                    iis.Close();

                    // Reposition the uncompressed stream and read all the dictionary strings from it.
                    uncompressedStream.Position = 0;
                    reader = new BinaryReader(uncompressedStream, Encoding.UTF8);
                    while ( reader.BaseStream.Position < uncompressedStream.Length) {
                        string dicString = reader.ReadString();
                        result.Add(nextId++, dicString);
                    }
                    uncompressedStream.Close();
                }

                return result;
            }
        }
    }
}
