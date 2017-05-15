// Example 198 from page 165 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Generic Arraylist: compile-time type check, no run-time checks needed

using System;
using System.Collections.Generic;

class MyTest {
  public static void Main(String[] args) {
    List<Person> cool = new List<Person>();
    cool.Add(new Person("Kristen"));
    cool.Add(new Person("Bjarne"));
    //    cool.Add(new Exception("Larry"));   // Wrong, detected at compile-time
    cool.Add(new Person("Anders"));
    Person p = (Person)cool[2];         // No run-time check needed
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
  
