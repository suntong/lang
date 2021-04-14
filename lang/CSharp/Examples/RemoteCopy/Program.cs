using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Net.Sockets;
using System.Net;

namespace RemoteCopy
{
    class Program
    {
        static void Main(string[] args)
        {
            RemoteCopyService rcs = new RemoteCopyService();
            SocketsFileTransfer sft = new SocketsFileTransfer();
            rcs.Start();
        }
                
    }
}
