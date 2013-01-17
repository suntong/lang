//----------------------------------------------------------------
// DateTime.cs - Program to test the C# DateTime functionalities
//----------------------------------------------------------------

using System;

class DateTimeTest
{
    static void Main()
    {
        TestDates();
        TestTimeSpan();
        TestFormats();

        // Keep the console window open in debug mode.
        Console.WriteLine("Press any key to exit.");
        Console.ReadKey();
    }

    // TestDates, test the C# Date functionalities

    /// <summary>
    /// Test C# Date related functionalities
    /// </summary>
    static void TestDates()
    {
        Console.WriteLine("== Test Dates");
        GetDate();

        Console.WriteLine("Today: {0}", DateTime.Today);

        DateTime y = GetYesterday();
        Console.WriteLine("Yesterday: {0}", y);

        DateTime d1 = GetTomorrow();
        Console.WriteLine("Tomorrow: {0}", d1);

        // First day in year
        Console.WriteLine("First day: {0}",
        FirstDayOfYear());

        DateTime d2 = new DateTime(1999, 6, 1);
        Console.WriteLine("First day of 1999: {0}",
            FirstDayOfYear(d2));

        // Last day of the year
        Console.WriteLine("Last day: {0}",
        LastDayOfYear());

        DateTime d = new DateTime(1999, 6, 1);
        Console.WriteLine("Last day of 1999: {0}",
            LastDayOfYear(d));

    }

    /// <summary>
    /// Uses C# DateTime constructor
    /// </summary>
    static void GetDate()
    {
        // This DateTime is constructed with an instance constructor.
        // ... We write it to the console.
        // ... If this is today, the second line will be "True".
        DateTime value = new DateTime(2010, 1, 18);
        Console.WriteLine(value);
        Console.WriteLine(value == DateTime.Today);
    }

    // Uses C# DateTime to find yesterday by subtracting one day from the current day. 
    /// <summary>
    /// Gets the previous day to the current day.
    /// </summary>
    static DateTime GetYesterday()
    {
        // Add -1 to now
        return DateTime.Today.AddDays(-1);
    }

    // Uses C# AddDays
    /// <summary>
    /// Gets the next day, tomorrow.
    /// </summary>
    static DateTime GetTomorrow()
    {
        return DateTime.Today.AddDays(1);
    }

    // First day in year
    // use an overloaded method here, which can make usage of the methods easier as there are fewer method names.

    /// <summary>
    /// Gets the first day of the current year.
    /// </summary>
    static DateTime FirstDayOfYear()
    {
        return FirstDayOfYear(DateTime.Today);
    }

    /// <summary>
    /// Finds the first day of year of the specified day.
    /// </summary>
    static DateTime FirstDayOfYear(DateTime y)
    {
        return new DateTime(y.Year, 1, 1);
    }

    // Last day of the year
    
    /// <summary>
    /// Finds the last day of the year for today.
    /// </summary>
    static DateTime LastDayOfYear()
    {
        return LastDayOfYear(DateTime.Today);
    }

    /// <summary>
    /// Finds the last day of the year for the selected day's year.
    /// </summary>
    static DateTime LastDayOfYear(DateTime d)
    {
        // 1
        // Get first of next year
        DateTime n = new DateTime(d.Year + 1, 1, 1);
        // 2
        // Subtract 1 from it
        return n.AddDays(-1);
    }

    // TestTimeSpan, Test TimeSpan

    static void TestTimeSpan()
    {
        Console.WriteLine("\n== Test TimeSpan");
        
        TimeSpan span1 = new TimeSpan(3, 0, 0);
        TimeSpan span2 = new TimeSpan(2, 0, 0);
        // Subtract two hours from three hours.
        TimeSpan span3 = span1.Subtract(span2);

        Console.WriteLine(span3);
    }

    // TestFormats, Test DateTime Formatting

    static void TestFormats()
    {
        Console.WriteLine("\n== Test DateTime Formats");

        // create date time 2008-03-09 16:05:07.123
        DateTime dt = new DateTime(2008, 3, 9, 16, 5, 7, 123);

        Console.WriteLine(String.Format("{0:y yy yyy yyyy}", dt));  // "8 08 008 2008"   year
        Console.WriteLine(String.Format("{0:M MM MMM MMMM}", dt));  // "3 03 Mar March"  month
        Console.WriteLine(String.Format("{0:d dd ddd dddd}", dt));  // "9 09 Sun Sunday" day
        Console.WriteLine(String.Format("{0:h hh H HH}", dt));  // "4 04 16 16"      hour 12/24
        Console.WriteLine(String.Format("{0:m mm}", dt));  // "5 05"            minute
        Console.WriteLine(String.Format("{0:s ss}", dt));  // "7 07"            second
        Console.WriteLine(String.Format("{0:f ff fff ffff}", dt));  // "1 12 123 1230"   sec.fraction
        Console.WriteLine(String.Format("{0:F FF FFF FFFF}", dt));  // "1 12 123 123"    without zeroes
        Console.WriteLine(String.Format("{0:t tt}", dt));  // "P PM"            A.M. or P.M.
        Console.WriteLine(String.Format("{0:z zz zzz}", dt));  // "-6 -06 -06:00"   time zone

        Console.WriteLine(String.Format("{0:d/M/yyyy HH:mm:ss}", dt));  // "9/3/2008 16:05:07" - english (en-US)

        // month/day numbers without/with leading zeroes
        Console.WriteLine(String.Format("{0:M/d/yyyy}", dt));            // "3/9/2008"
        Console.WriteLine(String.Format("{0:MM/dd/yyyy}", dt));          // "03/09/2008"

        // day/month names
        Console.WriteLine(String.Format("{0:ddd, MMM d, yyyy}", dt));    // "Sun, Mar 9, 2008"
        Console.WriteLine(String.Format("{0:dddd, MMMM d, yyyy}", dt));  // "Sunday, March 9, 2008"

        // two/four digit year
        Console.WriteLine(String.Format("{0:MM/dd/yy}", dt));            // "03/09/08"
        Console.WriteLine(String.Format("{0:MM/dd/yyyy}", dt));          // "03/09/2008"

        // standard format specifiers in String.Format method 
        Console.WriteLine("-- Standard formats");
        Console.WriteLine(String.Format("{0:t}", dt));  // "4:05 PM"                         ShortTime
        Console.WriteLine(String.Format("{0:d}", dt));  // "3/9/2008"                        ShortDate
        Console.WriteLine(String.Format("{0:T}", dt));  // "4:05:07 PM"                      LongTime
        Console.WriteLine(String.Format("{0:D}", dt));  // "Sunday, March 09, 2008"          LongDate
        Console.WriteLine(String.Format("{0:f}", dt));  // "Sunday, March 09, 2008 4:05 PM"  LongDate+ShortTime
        Console.WriteLine(String.Format("{0:F}", dt));  // "Sunday, March 09, 2008 4:05:07 PM" FullDateTime
        Console.WriteLine(String.Format("{0:g}", dt));  // "3/9/2008 4:05 PM"                ShortDate+ShortTime
        Console.WriteLine(String.Format("{0:G}", dt));  // "3/9/2008 4:05:07 PM"             ShortDate+LongTime
        Console.WriteLine(String.Format("{0:m}", dt));  // "March 09"                        MonthDay
        Console.WriteLine(String.Format("{0:y}", dt));  // "March, 2008"                     YearMonth
        Console.WriteLine(String.Format("{0:r}", dt));  // "Sun, 09 Mar 2008 16:05:07 GMT"   RFC1123
        Console.WriteLine(String.Format("{0:s}", dt));  // "2008-03-09T16:05:07"             SortableDateTime
        Console.WriteLine(String.Format("{0:u}", dt));  // "2008-03-09 16:05:07Z"            UniversalSortableDateTime

        // DateTime format strings
        Console.WriteLine("-- DateTime format strings");
        DateTime time = DateTime.Now;              // Use current time
        string format = "MMM ddd d HH:mm yyyy";    // Use this format
        Console.WriteLine(time.ToString(format));  // Write to console
        format = "M d h:mm yy";            // Use this format
        Console.WriteLine(time.ToString(format)); // Write to console

        DateTime now = DateTime.Now;
        Console.WriteLine(now.ToString("d"));
        Console.WriteLine(now.ToString("D"));
        Console.WriteLine(now.ToString("f"));
        Console.WriteLine(now.ToString("F"));
        Console.WriteLine(now.ToString("g"));
        Console.WriteLine(now.ToString("G"));
        Console.WriteLine(now.ToString("m"));
        Console.WriteLine(now.ToString("M"));
        Console.WriteLine(now.ToString("o"));
        Console.WriteLine(now.ToString("O"));
        Console.WriteLine(now.ToString("s"));
        Console.WriteLine(now.ToString("t"));
        Console.WriteLine(now.ToString("T"));
        Console.WriteLine(now.ToString("u"));
        Console.WriteLine(now.ToString("U"));
        Console.WriteLine(now.ToString("y"));
        Console.WriteLine(now.ToString("Y"));

        Console.WriteLine(now.ToLongDateString());  // Equivalent to D
        Console.WriteLine(now.ToLongTimeString());  // Equivalent to T
        Console.WriteLine(now.ToShortDateString()); // Equivalent to d
        Console.WriteLine(now.ToShortTimeString()); // Equivalent to t
        Console.WriteLine(now.ToString());

        now = DateTime.Today;
        for (int i = 0; i < 7; i++)
        {
            Console.WriteLine(now.ToString("ddd"));
            Console.WriteLine(now.ToString("dddd"));
            now = now.AddDays(1);
        }

        Console.WriteLine("-- Display AM and PM ");
        for (int i = 0; i < 2; i++)
        {
            Console.WriteLine(now.ToString("tt "));
            now = now.AddHours(12);
        }

        Console.WriteLine("-- Display year ");
        Console.WriteLine(now.ToString("y "));
        Console.WriteLine(now.ToString("yy"));
        now = new DateTime(2009, 1, 18);
        Console.WriteLine(now.ToString("y "));
        Console.WriteLine(now.ToString("yy"));
        Console.WriteLine(now.ToString("yyy"));   // <-- Don't use this
        Console.WriteLine(now.ToString("yyyy"));
        Console.WriteLine(now.ToString("yyyyy")); // <-- Don't use this
    }
}
