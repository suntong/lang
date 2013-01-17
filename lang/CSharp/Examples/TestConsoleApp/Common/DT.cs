//----------------------------------------------------------------
// DT.cs - Common C# DateTime functions
//----------------------------------------------------------------

using System;

namespace Common
{
    public class DT
    {
        // Uses C# DateTime to find yesterday by subtracting one day from the current day. 
        /// <summary>
        /// Gets the previous day to the current day.
        /// </summary>
        public static DateTime GetYesterday()
        {
            // Add -1 to now
            return DateTime.Today.AddDays(-1);
        }

        // Uses C# AddDays
        /// <summary>
        /// Gets the next day, tomorrow.
        /// </summary>
        public static DateTime GetTomorrow()
        {
            return DateTime.Today.AddDays(1);
        }

        // First day in year
        // use an overloaded method here, which can make usage of the methods easier as there are fewer method names.

        /// <summary>
        /// Gets the first day of the current year.
        /// </summary>
        public static DateTime FirstDayOfYear()
        {
            return FirstDayOfYear(DateTime.Today);
        }

        /// <summary>
        /// Finds the first day of year of the specified day.
        /// </summary>
        public static DateTime FirstDayOfYear(DateTime y)
        {
            return new DateTime(y.Year, 1, 1);
        }

        // Last day of the year

        /// <summary>
        /// Finds the last day of the year for today.
        /// </summary>
        public static DateTime LastDayOfYear()
        {
            return LastDayOfYear(DateTime.Today);
        }

        /// <summary>
        /// Finds the last day of the year for the selected day's year.
        /// </summary>
        public static DateTime LastDayOfYear(DateTime d)
        {
            // 1
            // Get first of next year
            DateTime n = new DateTime(d.Year + 1, 1, 1);
            // 2
            // Subtract 1 from it
            return n.AddDays(-1);
        }

    }
}
