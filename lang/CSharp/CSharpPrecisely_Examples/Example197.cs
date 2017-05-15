// Example 197 from page 165 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Object Arraylist: no compile-time type check

using System;
using SC = System.Collections;

class MyTest {
  public static void Main(String[] args) {
    SC.ArrayList cool = new SC.ArrayList();    // Needs: using SC = System.Collections;
    cool.Add(new Person("Kristen"));
    cool.Add(new Person("Bjarne"));
    cool.Add(new Exception("Larry"));          // Wrong, but no compiletime check
    cool.Add(new Person("Anders"));
    Person p = (Person)cool[2];                // Compiles OK, but throws at runtime
  }
}

class Person {
  private static int counter = 0;
  private readonly String name;
  private readonly int serialNumber;

  public Person(String name) {
    this.name = name;
    this.serialNumber = counter++;
  }
}
  
