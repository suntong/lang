using System;
using System.IO;

namespace ConsoleApp
{
    class Program
    {
        static void Main(string[] args)
        {
            TextWriter stdout = Console.Out;
            stdout.WriteLine(DateTime.Today);
        }
    }
}
