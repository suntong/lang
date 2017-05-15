// Example 39 from page 31 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

class MyTest {
  public static void Main(String[] args) {
    double[,][] rate = new double[10,12][];
    rate[0, 0] = new double[31];        // Jan 2000 has 31 days
    rate[0, 1] = new double[29];        // Feb 2000 has 29 days
    rate[0, 2] = new double[31];        // Mar 2000 has 31 days
    rate[0, 3] = new double[30];        // Apr 2000 has 30 days
    rate[0,11] = new double[31];        // Dec 2000 has 31 days
    rate[1, 0] = new double[31];        // Jan 2001 has 31 days
    rate[1, 1] = new double[28];        // Feb 2001 has 28 days
    rate[2, 1] = new double[28];        // Feb 2002 has 28 days
    rate[3, 1] = new double[28];        // Feb 2003 has 28 days
    rate[3,11] = new double[31];        // Dec 2003 has 31 days
    rate[4, 0] = new double[31];        // Jan 2004 has 31 days

    // USD per EUR daily Interbank exchange rates (www.oanda.com)
    rate[0, 1][27] = 0.9748;            // 28 Feb 2000
    rate[0, 1][28] = 0.9723;            // 29 Feb 2000
    rate[0, 2][ 0] = 0.9651;            //  1 Mar 2000
    rate[0,11][30] = 0.9421;            // 31 Dec 2000
    rate[1, 0][ 0] = 0.9421;            //  1 Jan 2001
    rate[1, 1][27] = 0.9180;            // 28 Feb 2001
    rate[2, 1][27] = 0.8641;            // 28 Feb 2002
    rate[3, 1][27] = 1.0759;            // 28 Feb 2003
    rate[3,11][ 1] = 1.1983;            //  1 Dec 2003
    rate[3,11][30] = 1.2557;            // 31 Dec 2003
    rate[4, 0][ 6] = 1.2741;            //  7 Jan 2004

    for (int y=0; y<rate.GetLength(0); y++)
      for (int m=0; m<rate.GetLength(1); m++)
        if (rate[y,m] != null) 
          for (int d=0; d<rate[y,m].Length; d++)
            if (rate[y,m][d] != 0.0)
              Console.WriteLine("{0:D4}-{1:D2}-{2:D2}: {3:F4} $US/Euro", 
                                y+2000, m+1, d+1, rate[y,m][d]);
  }
}
