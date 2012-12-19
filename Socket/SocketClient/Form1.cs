using System;
using System.Windows.Forms;
using System.Net.Sockets;
using System.Text;

using System.Net; 

namespace SocketClient
{
    public partial class Form1 : Form
    {
        const int port = 8888;
        const int testCase = 2;

        System.Net.Sockets.TcpClient clientSocket = new System.Net.Sockets.TcpClient();
        Socket client;

        public Form1()
        {
            InitializeComponent();
        }

        public void msg(string mesg)
        {
            textBox1.Text = textBox1.Text + Environment.NewLine + " >> " + mesg;
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            msg("Client Started");
            string theServer = "127.0.0.1";
            //theServer = "TorsvPerf01";
            if (testCase == 1)
                clientSocket.Connect(theServer, 8888);
            else
            {
                IPAddress ipAddress = IPAddress.Parse("127.0.0.1");
                IPEndPoint localEndPoint = new IPEndPoint(ipAddress, port);

                client = new Socket(localEndPoint.AddressFamily, SocketType.Stream, ProtocolType.Tcp);
                client.Connect(localEndPoint);
            }

            label1.Text = "Client Socket Program - Server Connected ...";
        }

        private void button1_Click(object sender, EventArgs e)
        {

            if (testCase == 1)
                NetworkStreamTest();
            else
                SocketSendTest();
        }

        private void NetworkStreamTest()
        {
            NetworkStream serverStream = clientSocket.GetStream();
            byte[] outStream = System.Text.Encoding.ASCII.GetBytes("Message from Client$");
            serverStream.Write(outStream, 0, outStream.Length);
            serverStream.Flush();

            byte[] inStream = new byte[10025];
            serverStream.Read(inStream, 0, (int)clientSocket.ReceiveBufferSize);
            string returndata = System.Text.Encoding.ASCII.GetString(inStream);
            msg("Data from Server : " + returndata);
        }


        private void SocketSendTest()
        {
            // http://msdn.microsoft.com/en-us/library/w93yy28a.aspx

            Byte[] sendBytes = Encoding.ASCII.GetBytes("Message from Client$");

            // Blocks until send returns. 
            int i = client.Send(sendBytes);

            // Get reply from the server.
            byte[] bytes = new byte[256];
            i = client.Receive(bytes);
            msg("Data from Server : " + Encoding.UTF8.GetString(bytes));
        }

    }
}
