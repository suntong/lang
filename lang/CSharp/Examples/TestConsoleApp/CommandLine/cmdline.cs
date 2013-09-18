// cmdline.cs
// http://msdn.microsoft.com/en-us/library/aa288457.aspx

using System;

namespace CommandLine
{
    public class cmdline
    {
        public static void Main(string[] args)
        {
            // The Length property is used to obtain the length of the array. 
            // Notice that Length is a read-only property:
            Console.WriteLine("Number of command line parameters = {0}",
               args.Length);
            for (int i = 0; i < args.Length; i++)
            {
                Console.WriteLine("Arg[{0}] = [{1}]", i, args[i]);
            }

            // Another approach to iterating over the array is to use the foreach statement: 
            foreach (string s in args)
            {
                Console.WriteLine(s);
            }

            // Keep the console window open in debug mode.
            Console.WriteLine("Press any key to exit.");
            Console.ReadKey();
        }
    }
}
