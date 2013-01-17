// A Hello World! program in C#.
using System;
namespace HelloWorld
{
    class Hello
    {
        static void Main()
        {
            Console.WriteLine("Hello World!");

            Console.WriteLine("== Test Dates");

            Console.WriteLine("Today: {0}", DateTime.Today);

            DateTime y = Common.DT.GetYesterday();
            Console.WriteLine("Yesterday: {0}", y);

            DateTime d1 = Common.DT.GetTomorrow();
            Console.WriteLine("Tomorrow: {0}", d1);

            // First day in year
            Console.WriteLine("First day: {0}",
            Common.DT.FirstDayOfYear());

            DateTime d2 = new DateTime(1999, 6, 1);
            Console.WriteLine("First day of 1999: {0}",
                Common.DT.FirstDayOfYear(d2));

            // Last day of the year
            Console.WriteLine("Last day: {0}",
            Common.DT.LastDayOfYear());

            DateTime d = new DateTime(1999, 6, 1);
            Console.WriteLine("Last day of 1999: {0}",
                Common.DT.LastDayOfYear(d));

            DateTime now = DateTime.Now;
            Console.WriteLine(String.Format("{0:yyyy-MM-ddThh:mm:ss.fff}", now));
            Console.WriteLine(String.Format("{0:yyyy-MM-ddThh:mm:ss.ffffff}", now));

            // Keep the console window open in debug mode.
            Console.WriteLine("Press any key to exit.");
            Console.ReadKey();
        }
    }
}


