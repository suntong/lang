using System;
using System.Collections.Generic;
using System.Text;
using System.Net;
using System.Net.Sockets;
using System.IO;
using System.Threading;

namespace RemoteCopy
{
    public class SocketsFileTransfer
    {
        private const string STRING_TO_CLOSE = "CLOSE APP";

        string filenameSend, filenameRecv;

        // http://msdn.microsoft.com/en-us/library/system.net.sockets.socket.aspx

        public static Socket ConnectSocket(string server, int port)
        {
            // System.Net.Sockets.Socket Class
            // http://msdn.microsoft.com/en-us/library/system.net.sockets.socket.aspx
            
            Socket s = null;
            IPHostEntry hostEntry = null;

            // Get host related information.
            hostEntry = Dns.GetHostEntry(server);

            // Loop through the AddressList to obtain the supported AddressFamily. This is to avoid 
            // an exception that occurs when the host IP Address is not compatible with the address family 
            // (typical in the IPv6 case). 
            foreach(IPAddress address in hostEntry.AddressList)
            {
                IPEndPoint ipe = new IPEndPoint(address, port);
                Socket tempSocket = 
                    new Socket(ipe.AddressFamily, SocketType.Stream, ProtocolType.Tcp);

                tempSocket.Connect(ipe);

                if(tempSocket.Connected)
                {
                    s = tempSocket;
                    break;
                }
                else
                {
                    continue;
                }
            }
            return s;
        }

        public int SendFileName(Socket client, string filename)
        {
            // http://msdn.microsoft.com/en-us/library/w93yy28a.aspx

            try
            {
                Byte[] sendBytes = Encoding.ASCII.GetBytes(filename);

                // Blocks until send returns. 
                int i = client.Send(sendBytes);
                Console.WriteLine("Sent {0} bytes.", i);

                // Get reply from the server.
                byte[] bytes = new byte[256];
                i = client.Receive(bytes);
                Console.WriteLine(Encoding.UTF8.GetString(bytes));
            }
            catch (SocketException e)
            {
                Console.WriteLine("{0} Error code: {1}.", e.Message, e.ErrorCode);
                return (e.ErrorCode);
            }
            return 0;
        }

        public void SendFile(IPAddress ipAddress, int port, string filenameToUse)
        {
            // http://msdn.microsoft.com/en-us/library/sx0a40c2.aspx

            IPEndPoint ipEndPoint = new IPEndPoint(ipAddress, port);

            // Create a TCP socket.
            Socket client = new Socket(AddressFamily.InterNetwork,
                    SocketType.Stream, ProtocolType.Tcp);

            // Connect the socket to the remote endpoint.
            client.Connect(ipEndPoint);

            // Send file to remote device
            filenameToUse = @"D:\Tmp\run.bat";
            Console.WriteLine("Sending {0} to the host.", filenameToUse);
            client.SendFile(filenameToUse);

            // Release the socket.
            client.Shutdown(SocketShutdown.Both);
            client.Close();
        }

        public void ReceiveFile(int port, string filepathandname)
        {
            string methodname = "ReceiveFile";

            try
            {
                IPEndPoint ipEnd = new IPEndPoint(IPAddress.Any, port);
                Socket sock = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.IP);
                sock.Bind(ipEnd);
                sock.Listen(1);
                Socket serverSocket = sock.Accept();
                byte[] data = new byte[1000000];
                int received = serverSocket.Receive(data);
                int filenameLength = BitConverter.ToInt32(data, 0);
                string filename = Encoding.ASCII.GetString(data, 4, filenameLength);

                this.CreateDirectoryFromPath(filepathandname);

                BinaryWriter bWrite = new BinaryWriter(File.Open(filepathandname, FileMode.Create));
                bWrite.Write(data, filenameLength + 4, received - filenameLength - 4);
                int received2 = serverSocket.Receive(data);
                while (received2 > 0)
                {
                    bWrite.Write(data, 0, received2);
                    received2 = serverSocket.Receive(data);
                }
                bWrite.Close();
                serverSocket.Close();
                sock.Close();
                MyLogs.WriteLog(methodname, "File copied ok: " + filepathandname, false);
            }
            catch (Exception ex)
            {
                MyLogs.WriteLog(methodname, "Port: " + port + " file: " + filepathandname + " " + ex.ToString(), true);
            }
        }

        

        public string ReceiveFilename(IPAddress ipAddress, int port)
        {
            string filename = string.Empty;
            string methodname = "ReceiveFilename";

            try
            {
                
                TcpListener tcpListener = new TcpListener(ipAddress, port);
                //Starting the listening
                tcpListener.Start();
                //Wait for a client
                Socket socketForClient = tcpListener.AcceptSocket();

                if (socketForClient.Connected)
                {       
                    // If connected
                    MyLogs.WriteLog(methodname,"Connected",false);
                    NetworkStream networkStream = new NetworkStream(socketForClient);
                    StreamWriter streamWriter = new StreamWriter(networkStream);
                    StreamReader streamReader = new StreamReader(networkStream);            
                    string theString = "Ok so far";

                    //Wait for the client request
                    filename = streamReader.ReadLine();
                    MyLogs.WriteLog(methodname,theString + " " + filename,false);

                    //Answering the client
                    streamWriter.WriteLine(theString);
                    streamWriter.Flush();

                    //Close
                    streamReader.Close();
                    streamWriter.Close();
                    networkStream.Close();
                    socketForClient.Close();
                    tcpListener.Stop();

                    //If the client has requested the close of the server
                    if (filename == STRING_TO_CLOSE)
                    {
                        MyLogs.WriteLog(methodname, "Closing requested", false);
                        Environment.Exit(999);
                    }
                }
            }
            catch (Exception ex)
            {
                MyLogs.WriteLog("ReceiveFilename", ex.ToString(), true);
            }

            return filename;
        }

        private void CreateDirectoryFromPath(string path)
        {
            string directoryPath = System.IO.Path.GetDirectoryName(path);
            if (Directory.Exists(directoryPath) == false)
            {
                try
                {
                    Directory.CreateDirectory(directoryPath);
                    MyLogs.WriteLog("CreateDirectoryFromPath", "Directory: " + directoryPath + " created", false);
                }
                catch (Exception ex)
                {
                    MyLogs.WriteLog("CreateDirectoryFromPath", "Cant create directory for: " + path + " " + ex.ToString(), true);
                }
            }
        }
    }
}