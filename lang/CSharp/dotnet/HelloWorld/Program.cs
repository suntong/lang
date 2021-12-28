// https://aka.ms/dotnet-hello-world
// See https://aka.ms/new-console-template for more information

Console.WriteLine("Hello, World!");
Console.WriteLine("The current time is " + DateTime.Now);

/*

$ dotnet run
Hello, World!
The current time is 12/28/2021 10:52:26


$ dotnet publish
Microsoft (R) Build Engine version 17.0.0+c9eb9dd64 for .NET
Copyright (C) Microsoft Corporation. All rights reserved.

  Determining projects to restore...
  All projects are up-to-date for restore.
  HelloWorld -> .../lang/CSharp/dotnet/HelloWorld/bin/Debug/net6.0/HelloWorld.dll
  HelloWorld -> .../lang/CSharp/dotnet/HelloWorld/bin/Debug/net6.0/publish/

$ bin/Debug/net6.0/publish/HelloWorld 
Hello, World!
The current time is 12/28/2021 11:12:29

*/
