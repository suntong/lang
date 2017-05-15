// Example 103 from page 85 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

// A log of Strings that retains only the last SIZE logged Strings

public class Log {
  private const int SIZE = 5;
  private static int instanceCount = 0;
  private int count = 0;
  private String[] log = new String[SIZE];

  public Log() {
    instanceCount++;
  }
  
  // The number of Logs created

  public static int InstanceCount {
    get { return instanceCount; }
  }

  // Add a String to this Log

  public void Add(String msg) {
    log[count++ % SIZE] = msg;
  }

  // Property giving the number of strings inserted in this Log

  public int Count {
    get { return count; }
  }

  // The most recently logged string, if any
  
  public String Last {
    get { // Return the last log entry, or null if nothing logged yet
      return count==0 ? null : log[(count-1)%SIZE];
    }
    set { // Update the last log entry, or create one if nothing logged yet 
      if (count==0)
        log[count++] = value;
      else
        log[(count-1)%SIZE] = value;
    }
  }    

  // Return all log entries

  public String[] All {
    get {
      int size = Math.Min(count, SIZE);
      String[] res = new String[size];
      for (int i=0; i<size; i++)
        res[i] = log[(count-size+i) % SIZE];
      return res;
    }
  }
}

class TestLog {
  public static void Main(String[] args) {
    Log log1 = new Log(), log2 = new Log();
    Console.WriteLine("Number of logs = " + Log.InstanceCount);
    log1.Add("Alarm"); log1.Add("Shower"); log1.Add("Coffee"); 
    log1.Add("Bus"); log1.Add("Work"); log1.Add("Lunch"); 
    Console.WriteLine(log1.Last);
    log1.Last += " nap";
    Console.WriteLine(log1.Last);
    log1.Add("More work");
    Console.WriteLine("Logged entries = " + log1.Count);
    foreach (String s in log1.All)
      Console.WriteLine(s);
  }
}
