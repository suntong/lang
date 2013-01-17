//----------------------------------------------------------------
// Sizes.cs--Program to tell the size of the C# variable types
//----------------------------------------------------------------
using System;
class Sizes
{
    public static void Main()
    {
        Console.WriteLine("\nA byte     is {0} byte(s)", sizeof(byte));
        Console.WriteLine("A sbyte    is {0} byte(s)", sizeof(sbyte));
        Console.WriteLine("A char     is {0} byte(s)", sizeof(char));
        Console.WriteLine("\nA short    is {0} byte(s)", sizeof(short));
        Console.WriteLine("An ushort is {0} byte(s)", sizeof(ushort));
        Console.WriteLine("\nAn int     is {0} byte(s)", sizeof(int));
        Console.WriteLine("An uint    is {0} byte(s)", sizeof(uint));
        Console.WriteLine("\nA long     is {0} byte(s)", sizeof(long));
        Console.WriteLine("An ulong   is {0} byte(s)", sizeof(ulong));
        Console.WriteLine("\nA float    is {0} byte(s)", sizeof(float));
        Console.WriteLine("A double   is {0} byte(s)", sizeof(double));
        Console.WriteLine("\nA decimal is {0} byte(s)", sizeof(decimal));
        Console.WriteLine("\nA boolean is {0} byte(s)", sizeof(bool));

        // Keep the console window open in debug mode.
        Console.WriteLine("Press any key to exit.");
        Console.ReadKey();
    }
}
