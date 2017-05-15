// Example 151 from page 121 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Using enums in calendrical calculations

using System;

public enum Day {
  Mon, Tue, Wed, Thu, Fri, Sat, Sun
}

public enum Month {
  Jan=1, Feb, Mar, Apr, May, Jun, Jul, Aug, Sep, Oct, Nov, Dec
}

public class Date {
  readonly int yy /* 0-9999 */, dd /* 1-31 */;
  readonly Month mm;

  public Date(int yy, Month mm, int dd) {
    if (Ok(yy, mm, dd)) {
      this.yy = yy; this.mm = mm; this.dd = dd;
    } else
      throw new Exception("Illegal date ("+yy+","+mm+","+dd+")");
  }

  public static bool LeapYear(int y) {
    return y % 4 == 0 && y % 100 != 0 || y % 400 == 0;
  }

  public bool LeapYear() {
    return LeapYear(yy);
  }

  public static int MonthDays(int y, Month m) {
    switch (m) {
    case Month.Apr: case Month.Jun: case Month.Sep: case Month.Nov:
      return 30;
    case Month.Feb: 
      return LeapYear(y) ? 29 : 28;
    default:
      return 31;
    }
  }

  public int MonthDays() {
    return MonthDays(yy, mm);
  }

  public static int YearDays(int y) {
    return LeapYear(y) ? 366 : 365;
  }

  public int YearDays() {
    return YearDays(yy);
  }

  public static bool Ok(int y, Month m, int d) {
    return 1 <= d && d <= MonthDays(y, m);
  }

  // ISO week numbers: the week is from Monday to Sunday.  Week 1 is
  // the first week having a Thursday.

  public static int WeekNumber(int y, Month m, int d) {
    int yday = DayInYear(y, m, d);
    int wday = (int)Weekday(y, m, d);
    int week = (yday - wday + 10)/7;
    if (week == 0) 
      return (yday + YearDays(y-1) - wday + 10)/7;
    else
      return week;
  }

  public int WeekNumber() {
    return WeekNumber(yy, mm, dd);
  }

  // Translated from Emacs's calendar.el:
  // Reingold: Number of the day within the year: 

  public static int DayInYear(int y, Month m, int d) {
    int monthno = (int)m - 1;
    int monthadjust = 
      monthno > 1 ? (27 + 4 * monthno) / 10 - (LeapYear(y) ? 1 : 0) : 0;
    return d - 1 + 31 * monthno - monthadjust;
  }

  public int DayInYear() {
    return DayInYear(yy, mm, dd);
  }

  // Reingold: Find the number of days elapsed from the (imagined)
  // Gregorian date Sunday, December 31, 1 BC to the given date.
        
  public static int ToDaynumber(int y, Month m, int d) {
    int prioryears = y - 1;
    return 
      DayInYear(y, m, d) 
      + 1 + 365 * prioryears
      + prioryears / 4 - prioryears / 100 + prioryears / 400;
  }

  public int ToDaynumber() {  
    return ToDaynumber(yy, mm, dd);
  }

  // Reingold et al: from absolute day number to year, month, date: 

  public static Date FromDaynumber(int n) { 
    int d0 = n - 1;
    int n400 = d0 / 146097;
    int d1 = d0 % 146097;
    int n100 = d1 / 36524;
    int d2 = d1 % 36524;
    int n4 = d2 / 1461;
    int d3 = d2 % 1461;
    int n1 = d3 / 365;
    int d = 1 + d3 % 365;
    int y = 400 * n400 + 100 * n100 + n4 * 4 + n1 + 1;
    if (n100 == 4 || n1 == 4) {
      return new Date(y-1, Month.Dec, 31);
    } else {
      Month m = Month.Jan;
      int mdays;
      while ((mdays = MonthDays(y, m)) < d) {
        d -= mdays;
        m++;
      }
      return new Date(y, m, d);
    }
  }
  
  // Day of the week: 0=Mon, 1=Tue, ..., 6=Sun

  public static Day Weekday(int y, Month m, int d) {
    return (Day)((ToDaynumber(y, m, d)+6) % 7);
  }

  public Day Weekday() {
    return Weekday(yy, mm, dd);
  }

  public override String ToString() { // ISO format such as 2003-05-31
    return String.Format("{0:D4}-{1:D2}-{2:D2}", yy, (int)mm, dd);
  }
}

class Example151 {
  public static void Main(String[] args) {
    if (args.Length != 3) 
      Console.WriteLine("Usage: Example151 yyyy mm dd\n");
    else {
      Date d = new Date(int.Parse(args[0]), 
                        (Month)int.Parse(args[1]), 
                        int.Parse(args[2]));
      Console.WriteLine(d + " is " + d.Weekday() + " in week " + d.WeekNumber());
    }
  }
}
