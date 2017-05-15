// Example 70 from page 57 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Using events to pick up readings from a simulated thermometer.

using System;
using System.Threading;

delegate void Handler(double temperature);

class Thermometer {
  public event Handler Reading;
  private int temperature = 80;
  private static Random rnd = new Random();

  public Thermometer() {
    new Thread(new ThreadStart(Run)).Start();
  }

  private void Run() {
    for (;;) {                          // Forever simulate new readings
      temperature += rnd.Next(-5, 6);   // Random number in range -5..5
      if (Reading != null)              // If there are any handlers,
        Reading(temperature);           // call them with the new reading
      Thread.Sleep(rnd.Next(2000));
    }
  }
}

class MyTest {
  public static void Main(String[] args) {
    Thermometer t = new Thermometer();
    t.Reading += new Handler(PrintReading);
    t.Reading += new Handler(CountReading);
  }

  public static void PrintReading(double temperature) {
    Console.WriteLine(temperature);
  }

  public static void CountReading(double temperature) {
    if (++readCount % 10 == 0) 
      Console.WriteLine("Now {0} readings", readCount);
  }

  private static int readCount = 0;
}
