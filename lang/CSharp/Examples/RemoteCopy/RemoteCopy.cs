using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Net.Sockets;
using System.Net;

using System.Text.RegularExpressions;

namespace RemoteCopy
{
    public class RemoteCopyService
    {
        public const int portCmmd = 9999;
        public const int portData = 9998;

        public static IPAddress ipAddress;

        static RemoteCopyService()
        {
            ipAddress = GetIpFromHostName(string.Empty);
        }

        public void Start()
        {
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

            SocketsFileTransfer sockets = new SocketsFileTransfer();
            
            //The file names are received from 9999 port
            int port = 9999;
            string filename = sockets.ReceiveFilename(ipAddress, port);

            //The file data is received from 9998 port
            port = 9998;
            sockets.ReceiveFile(port, filename);

            MyLogs.WriteLog("Start","Process finished",false);

            //Start the process again
            this.Start();

        }

        public static IPAddress GetIpFromHostName(string servername)
        {
            IPAddress ip = null;

            String strHostName = string.Empty;
            if (servername == string.Empty)
            {
                strHostName = Dns.GetHostName();
                MyLogs.WriteLog("GetIpFromHostName", "Machine name: " + strHostName, false);
            }
            else
            {
                strHostName = servername;
            }

            String IP = ExecuteCommand("nslookup", strHostName);
            IP = Regex.Replace(IP, @".*Name:.*?\.com", "", RegexOptions.Singleline);
            IP = Regex.Replace(IP, @".*Address: *", "", RegexOptions.Singleline);
            IP = Regex.Replace(IP, @"[\r\n]", "");
            ip = IPAddress.Parse(IP);

            MyLogs.WriteLog("File Transfer Server started", "Watching on IP: " + ip, false);
            return ip;
        }

        /// <summary>
        /// ExecuteCommand, Executes a shell command synchronously.
        /// </summary>
        /// <param name="command">string command</param>
        /// <param name="arguments">string command arguments</param>
        /// <returns>string, as output of the command.</returns>

        static string ExecuteCommand(string command, string arguments)
        {
            // create the ProcessStartInfo 
            // E.g., using "cmd" as the program to be run, and "/c " as the parameters.
            // Incidentally, /c tells cmd that we want it to execute the command that follows,
            // and then exit.
            System.Diagnostics.ProcessStartInfo procStartInfo =
                new System.Diagnostics.ProcessStartInfo(command, arguments);

            // The following commands are needed to redirect the standard output.
            // This means that it will be redirected to the Process.StandardOutput StreamReader.
            procStartInfo.RedirectStandardOutput = true;
            procStartInfo.UseShellExecute = false;
            // Do not create the black window.
            procStartInfo.CreateNoWindow = true;
            // Now we create a process, assign its ProcessStartInfo and start it
            System.Diagnostics.Process proc = new System.Diagnostics.Process();
            proc.StartInfo = procStartInfo;
            proc.Start();
            // Get the output into a string
            string result = proc.StandardOutput.ReadToEnd();
            // Return the command output.
            return result;
        }

    }
}
