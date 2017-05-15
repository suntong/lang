// Example 196 from page 163 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;     
using System.IO;           // BinaryReader, BinaryWriter
using System.Net;          // AddressFamily, Dns, IPAddress, ProtocolType, ...
using System.Net.Sockets;  // NetworkStream, Socket, SocketType, 

class Example196 {
  const int PortNo = 2357;
  
  public static void Main(String[] args) {
    bool server = (args.Length == 1 && args[0] == "server");
    bool client = (args.Length == 2 && args[0] == "client");
    if (server) {               // Server: accept questions about primality
      Socket serversocket = 
        new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
      serversocket.Bind(new IPEndPoint(IPAddress.Any, PortNo));
      serversocket.Listen(10);  // Max queue = 10 connections.
      for (;;) {                // For ever, accept connections
        NetworkStream s = new NetworkStream(serversocket.Accept());
        BinaryReader input  = new BinaryReader(s);
        BinaryWriter output = new BinaryWriter(s);
        int number = input.ReadInt32();
        output.Write(IsPrime(number));
        input.Close(); output.Close();
      }
    } else if (client) {        // Client: ask questions about primality
      IPAddress ipa = Dns.GetHostEntry(args[1]).AddressList[0];
      for (int i=1; i<100; i++) {
        Socket clientsocket = 
          new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
        clientsocket.Connect(new IPEndPoint(ipa, PortNo));
        NetworkStream n = new NetworkStream(clientsocket);
        BinaryWriter output = new BinaryWriter(n);
        BinaryReader input = new BinaryReader(n);
        output.Write(i);
        if (input.ReadBoolean())
          Console.Write(i + " ");
        output.Close(); input.Close();
      }
    } else {                // Neither server nor client
      Console.WriteLine("Start two copies of this program, possibly on different machines:");
      Console.WriteLine("   Example196 server");
      Console.WriteLine("   Example196 client <serverhostname>");
      Console.WriteLine("Use `Example196 client localhost' if the"); 
      Console.WriteLine("client and server run on the same machine.");
      Console.WriteLine("You may start several clients all talking to the same server.");
    }
  }
    
  static bool IsPrime(int p) {
    if (p == 2)
      return true;
    if (p == 1 || p % 2 == 0) 
      return false;
    for (int q=3; q*q<=p; q+=2)
      if (p % q == 0)
        return false;
    return true;
  }
}
