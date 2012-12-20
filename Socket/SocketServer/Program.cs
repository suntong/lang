using System;
using System.Net.Sockets;
using System.Text;

using System.Net;

namespace ConsoleApplication1
{
    class Program
    {
        public const int port = 8888;

        static void Main(string[] args)
        {
            int testCase = 2;

            if (testCase == 1)
                NetworkStreamTest();
            else
                SocketSendTest();
        }

        static void NetworkStreamTest()
        {
            TcpListener serverSocket = new TcpListener(port);
            TcpClient clientSocket = default(TcpClient);
            serverSocket.Start();
            Console.WriteLine(" >> Server Started");
            clientSocket = serverSocket.AcceptTcpClient();
            Console.WriteLine(" >> Accept connection from client");

            int requestCount = 0;
            while ((true))
            {
                try
                {
                    requestCount = requestCount + 1;
                    NetworkStream networkStream = clientSocket.GetStream();
                    byte[] bytesFrom = new byte[10025];
                    networkStream.Read(bytesFrom, 0, (int)clientSocket.ReceiveBufferSize);
                    string dataFromClient = System.Text.Encoding.ASCII.GetString(bytesFrom);
                    dataFromClient = dataFromClient.Substring(0, dataFromClient.IndexOf("$"));
                    Console.WriteLine(" >> Data from client - " + dataFromClient);
                    string serverResponse = "Server response " + Convert.ToString(requestCount);
                    Byte[] sendBytes = Encoding.ASCII.GetBytes(serverResponse);
                    networkStream.Write(sendBytes, 0, sendBytes.Length);
                    networkStream.Flush();
                    Console.WriteLine(" >> " + serverResponse);
                }
                catch (Exception ex)
                {
                    Console.WriteLine(ex.ToString());
                }
            }

            clientSocket.Close();
            serverSocket.Stop();
            Console.WriteLine(" >> exit");
            Console.ReadLine();
        }


        static void SocketSendTest()
        {
            // http://msdn.microsoft.com/en-us/library/dz10xcwh
            // http://msdn.microsoft.com/en-us/library/80z2essb

            IPHostEntry ipHostInfo = Dns.Resolve(Dns.GetHostName());
            IPAddress ipAddress = ipHostInfo.AddressList[0];
            ipAddress = IPAddress.Parse("127.0.0.1");
            IPEndPoint localEndPoint = new IPEndPoint(ipAddress, port);

            Socket listener =
                    new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
            listener.Bind(localEndPoint);
            listener.Listen(port);

            Console.WriteLine(" >> Server Started, waiting for a connection...");
            Socket clientSocket = listener.Accept();
            string remoteIP = ((IPEndPoint)clientSocket.RemoteEndPoint).ToString(); //IP:port
            Console.WriteLine(" >> Accepted connection from client " + remoteIP);

            // http://msdn.microsoft.com/en-us/library/w93yy28a.aspx
            int requestCount = 0;
            while ((true))
            {
                try
                {
                    byte[] bytesFrom = new byte[10025];

                    requestCount = requestCount + 1;
                    clientSocket.Receive(bytesFrom);

                    string dataFromClient = System.Text.Encoding.ASCII.GetString(bytesFrom);
                    dataFromClient = dataFromClient.Substring(0, dataFromClient.IndexOf("$"));
                    // must have above substring extraction, but why not for server-send/client-receive?
                    Console.WriteLine(" >> Data from client - " + dataFromClient);

                    string serverResponse = "Server response " + Convert.ToString(requestCount);
                    Byte[] sendBytes = Encoding.ASCII.GetBytes(serverResponse);

                    // Synchronous call. Blocks until send returns. 
                    int i = clientSocket.Send(sendBytes);
                    Console.WriteLine(" >> " + serverResponse);
                    Console.WriteLine("    Sent {0} bytes.", i);
                }
                catch (Exception ex)
                {
                    Console.WriteLine(ex.ToString());
                }
            }

        }

    }
}
