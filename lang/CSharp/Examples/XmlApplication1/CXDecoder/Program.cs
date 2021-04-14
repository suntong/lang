using System;
using System.IO;

using System.Text;
using System.Text.RegularExpressions;

using System.Xml;
using System.Xml.XPath;

using ICSharpCode.SharpZipLib;
using ICSharpCode.SharpZipLib.Zip.Compression.Streams;
using ICSharpCode.SharpZipLib.Zip.Compression;

namespace CXDecoder
{
    class Program
    {
        static void Main(string[] args)
        {
            string text = "";
            Stream st = Console.OpenStandardInput();
            using (StreamReader reader = new StreamReader(st))
            {
               text = reader.ReadToEnd();
            }

            text = "PABzADoARQBuAHYAZQBsAG8AcABlACAAeABtAGwAbgBzADoAcwA9ACIAaAB0AHQAcAA6AC8ALwBzAGMAaABlAG0AYQBzAC4AeABtAGwAcwBvAGEAcAAuAG8AcgBnAC8AcwBvAGEAcAAvAGUAbgB2AGUAbABvAHAAZQAvACIAPgA8AHMAOgBCAG8AZAB5AD4APABFAHgAZQBjAHUAdABlACAAeABtAGwAbgBzAD0AIgBoAHQAdABwADoALwAvAEQAYQB5AGYAbwByAGMAZQAvAFMAZQByAHYAaQBjAGUAcwAvAEMAbwByAGUAUwBlAHIAdgBpAGMAZQAiAD4APABiACAAeABtAGwAbgBzADoAZAA0AHAAMQA9ACIAaAB0AHQAcAA6AC8ALwBEAGEAeQBmAG8AcgBjAGUALwBEAGEAdABhAC8AQwBvAHIAZQBTAGUAcgB2AGkAYwBlACIAIAB4AG0AbABuAHMAOgBpAD0AIgBoAHQAdABwADoALwAvAHcAdwB3AC4AdwAzAC4AbwByAGcALwAyADAAMAAxAC8AWABNAEwAUwBjAGgAZQBtAGEALQBpAG4AcwB0AGEAbgBjAGUAIgA+ADwAZAA0AHAAMQA6AFMAZQBzAHMAaQBvAG4AVABpAGMAawBlAHQAPgBxAGEAcABhAHkAcgBvAGwAbAA1ADoAMwBjAGUAYgA2AGYAOABiAC0AMQA1ADcAYQAtADQANwBmADcALQA5ADgAOQAwAC0AMgAzAGEAMwAxADAAMAA4ADMANAA4ADYAPAAvAGQANABwADEAOgBTAGUAcwBzAGkAbwBuAFQAaQBjAGsAZQB0AD4APABkADQAcAAxADoASQBuAGYAbwBzAD4APABkADQAcAAxADoAQwBvAHIAZQBTAGUAcgB2AGkAYwBlAFIAZQBxAHUAZQBzAHQASQBuAGYAbwA+ADwAZAA0AHAAMQA6AFIAZQBxAHUAZQBzAHQATgBhAG0AZQA+AEcAZQB0ADwALwBkADQAcAAxADoAUgBlAHEAdQBlAHMAdABOAGEAbQBlAD4APAAvAGQANABwADEAOgBDAG8AcgBlAFMAZQByAHYAaQBjAGUAUgBlAHEAdQBlAHMAdABJAG4AZgBvAD4APAAvAGQANABwADEAOgBJAG4AZgBvAHMAPgA8AGQANABwADEAOgBQAGEAeQBsAG8AYQBkAHMAIAB4AG0AbABuAHMAOgBkADUAcAAxAD0AIgBoAHQAdABwADoALwAvAHMAYwBoAGUAbQBhAHMALgBtAGkAYwByAG8AcwBvAGYAdAAuAGMAbwBtAC8AMgAwADAAMwAvADEAMAAvAFMAZQByAGkAYQBsAGkAegBhAHQAaQBvAG4ALwBBAHIAcgBhAHkAcwAiAD4APABkADUAcAAxADoAYgBhAHMAZQA2ADQAQgBpAG4AYQByAHkAPgBBAFEASgB5AEEAQQBBAEEATABjAG8ANwBEAHMASQB3AEUAQQBYAEEASABvAG0AQwBRADEAQgBRADQAQQAzAFEAYwBZAEUAMABKAEUASQA0AEIAZQAzAEsAUABMAEMAVgArAE0ATgA2AHcAZABlAG4AWQBlAHIAWgA5AEUAZwBRAFgAbgByAG8ARABlADgAUABxAG0ANgA5AGEAagBrAFQAVwBjADkAUwBwAGwAegBNAEMAQwBVAEwAKwBRAGEASABTAHQAZQBGADkAWgBrAGwANwB2ADYAcgB0AFcAYgBhAHkAVwBSADUAMABiAEgAcgBEAG4AUQBmAEwAdABaADUAUgBOADYASABWAEoAVwBUAHcAMgBxAEEAKwB2AHcAWQBPAFcASgB0AFUAVwB2AEkAYQBRAHAAdQBoAHYANABBAEIAYwBKAFIAQwBrAEEAdwBHAEEARABnAGwAaQBTAGMAdwBSAEcAVwBzAGIASAB4AFoAaQBsADUAYwA0AFYALwAyAGsAbwB0AHcAMwBqAFkAVQBYAFoAYgBmAFoAOQBFAFIAVgBLAGkASQA1AFYAWgByAEIAYgA5AFQAdABaAHUARQBKAGIASABmAFoAZgA4AC8ASABGAHEANwAyAGMASQA2ADIAbQBjAHoARwBOADkAdwB3AFgAaABjAGQAWgAyAEkAOQAyADEANgBvADEAUQB1AE8AawA0AFkATQBZAE4AeAA0AE0AWQBDAEcANABwADAASQBZAFEAUQBaAG4AbwAwAFEAOAA9ADwALwBkADUAcAAxADoAYgBhAHMAZQA2ADQAQgBpAG4AYQByAHkAPgA8AC8AZAA0AHAAMQA6AFAAYQB5AGwAbwBhAGQAcwA+ADwALwBiAD4APAAvAEUAeABlAGMAdQB0AGUAPgA8AC8AcwA6AEIAbwBkAHkAPgA8AC8AcwA6AEUAbgB2AGUAbABvAHAAZQA+AA==";
            //text = "PABzADoARQBuAHYAZQBsAG8AcABlACAAeABtAGwAbgBzADoAcwA9ACIAaAB0AHQAcAA6AC8ALwBzAGMAaABlAG0AYQBzAC4AeABtAGwAcwBvAGEAcAAuAG8AcgBnAC8AcwBvAGEAcAAvAGUAbgB2AGUAbABvAHAAZQAvACIAPgA8AHMAOgBCAG8AZAB5AD4APABHAGUAdABBAGMAYwBlAHMAcwBQAGUAcgBtAGkAcwBzAGkAbwBuAHMAQwBvAG0AcAByAGUAcwBzAGUAZAAgAHgAbQBsAG4AcwA9ACIAaAB0AHQAcAA6AC8ALwB0AGUAbQBwAHUAcgBpAC4AbwByAGcALwAiAD4APABzAGUAcwBzAGkAbwBuAEkAZAA+AHEAYQBwAGEAeQByAG8AbABsADUAOgAzAGMAZQBiADYAZgA4AGIALQAxADUANwBhAC0ANAA3AGYANwAtADkAOAA5ADAALQAyADMAYQAzADEAMAAwADgAMwA0ADgANgA8AC8AcwBlAHMAcwBpAG8AbgBJAGQAPgA8AHIAbwBsAGUASQBkAD4AMQAwADAAMQA8AC8AcgBvAGwAZQBJAGQAPgA8AC8ARwBlAHQAQQBjAGMAZQBzAHMAUABlAHIAbQBpAHMAcwBpAG8AbgBzAEMAbwBtAHAAcgBlAHMAcwBlAGQAPgA8AC8AcwA6AEIAbwBkAHkAPgA8AC8AcwA6AEUAbgB2AGUAbABvAHAAZQA+AA==";
            Console.WriteLine(DecodeReq(text)); // Write to console.
        }


        /// <summary>
        /// DecodeReq: Decode the web request encoded in base64 stored in the MS webtest file
        /// </summary>

        private static string DecodeReq(string input)
        {
            byte[] bytes = Convert.FromBase64String(input);
            string xmlStr = System.Text.Encoding.Unicode.GetString(bytes);

            XmlDocument xmldoc = new XmlDocument();
            xmldoc.LoadXml(xmlStr);
            XmlNode root = xmldoc.DocumentElement;

            XmlNode compressedNode = root.SelectSingleNode(@"//*[local-name() = 'base64Binary']");

            if (compressedNode != null)
            {
                string xmlData = DecodeToXml(compressedNode.InnerXml);
                compressedNode.InnerXml = xmlData;
            }

            TextWriter errorWriter = Console.Error;
            xmldoc.Save(errorWriter);
            errorWriter.WriteLine("\n\n"); // double-line space

            String theXml = root.OuterXml;
            theXml = Regex.Replace(theXml,
                            @"<(CoreServiceGenericParameters xmlns=""http://Dayforce/Data/CoreService""" +
                             @"|GeneralGetRequest xmlns=""http://SharpTop.Net/Services/Platform"")",
                            @"$& xmlns:i=""http://www.w3.org/2001/XMLSchema-instance""");
            byte[] encodedXml = System.Text.Encoding.Unicode.GetBytes(theXml);
            return Convert.ToBase64String(encodedXml, Base64FormattingOptions.None);

        }

        private static byte[] base64Decode(string data)
        {
            byte[] todecode_byte = Convert.FromBase64String(data);
            return todecode_byte;
        }

        private static string Decompress(byte[] byteInput)
        {
            string result = "";
            if (byteInput != null)
            {
                using (MemoryStream ms = new MemoryStream(byteInput, 0, byteInput.Length))
                {
                    // skip the first 2 bytes of the stream.
                    ms.ReadByte(); ms.ReadByte();

                    XmlBinaryReaderSession xmlSession;
                    xmlSession = CreateReaderSession(ms);

                    // Set up the decoder.                
                    Stream s2 = new InflaterInputStream(ms, new Inflater(true));
                    XmlDictionaryReader dicReader = XmlDictionaryReader.CreateBinaryReader(s2, null, XmlDictionaryReaderQuotas.Max, xmlSession);
                    dicReader.MoveToStartElement();
                    result = dicReader.ReadOuterXml();
                }
            }

            return result;
        }

        public static string DecodeToXml(string encodedText)
        {
            return Decompress(base64Decode(encodedText));
        }

        public static System.Xml.XmlBinaryReaderSession CreateReaderSession(Stream stream)
        {
            XmlBinaryReaderSession result = new XmlBinaryReaderSession();

            // Read the length of the dictionary from the stream.
            BinaryReader reader = new BinaryReader(stream, Encoding.UTF8);
            int nextId = 0;
            int totalSessionLength = reader.ReadInt32();        // May be zero if no dictionary

            if (totalSessionLength > 0)
            {
                //
                // Decompress dictionary and store it in a byte array.
                //
                byte[] compressedDictionary = new byte[totalSessionLength];
                stream.Read(compressedDictionary, 0, totalSessionLength);
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
                while (reader.BaseStream.Position < uncompressedStream.Length)
                {
                    string dicString = reader.ReadString();
                    result.Add(nextId++, dicString);
                }
                uncompressedStream.Close();
            }

            return result;
        }

    }
}
