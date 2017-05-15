// Example 218 from page 179 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;       // IEnumerable<T>, IComparer<T>
using System.Drawing;                   // Color

// delegate R Foo<in R, out A>(A x);    // Invalid: R is an output type, A is an input type

delegate R Foo<out R, in A>(A x);  // OK

class MyTest {
  public static void Main(String[] args) {
    // Interface examples:
    Student[] students = { new Student("Anders"), new Student("Kasper"), new Student("Vincens") };
    IEnumerable<Student> ss = students;
    PrintPersons(ss);
    PrintPersons(students);
    bool sorted = Sorted(new PersonComparer(), students);
    // Delegate type examples: 
    Func<Person,ColoredPoint> pc = (Person p) => new ColoredPoint(2, p.name.Length, Color.Red);
    Func<Student,Point> sp = pc;
    Func<Func<Student,Point>,int> fspi 
      = (Func<Student,Point> fsp) => fsp(new Student("Lise")).y;
    Func<Func<Person,ColoredPoint>,int> fpci = fspi;
    Console.WriteLine(sp(new Student("Morten")));
    Console.WriteLine(fpci(pc));
  }

  static void PrintPersons(IEnumerable<Person> ps) {
    foreach (Person p in ps)
      Console.WriteLine(p.name);
  }

  static bool Sorted(IComparer<Student> cmp, Student[] a) {
    for (int i=1; i<a.Length; i++)
      if (cmp.Compare(a[i-1],a[i]) > 0)
        return false;
    return true;
  }
}

class Person { 
  public readonly String name;
  public Person(String name) { 
    this.name = name;
  }
}

class Student : Person { 
  public Student(String name) : base(name) { }
}

class PersonComparer : IComparer<Person> {
  public int Compare(Person p1, Person p2) {
    return p1.name.CompareTo(p2.name);
  }
}

public class Point {
  protected internal int x, y;
  public Point(int x, int y) { this.x = x; this.y = y; }
  public void Move(int dx, int dy) { x += dx; y += dy; }
  public override String ToString() { return "(" + x + ", " + y + ")"; }
}

class ColoredPoint : Point {
  protected Color c;
  public ColoredPoint(int x, int y, Color c) : base(x, y) { this.c = c; }
  public Color GetColor { get { return c; } }
}

namespace VarianceDemo {
  interface IEnumerator<out T> { T Current { get; } }
  interface IEnumerable<out T> { IEnumerator<T> GetEnumerator(); }
  interface IComparer<in T> { int Compare(T v1, T v2); }
  interface IComparable<in T> { int CompareTo(T v); }
  interface IEqualityComparer<in T> { bool Equals(T v1, T v2); int GetHashCode(T v); }
  interface IEquatable<in T> { bool Equals(T v); }
  delegate R Func<out R>();
  delegate R Func<in A1,out R>(A1 x1);
  delegate R Func<in A1,in A2,out R>(A1 x1, A2 x2);
  delegate void Action<out R>();
  delegate void Action<in A1,out R>(A1 x1);
  delegate void Action<in A1,in A2,out R>(A1 x1, A2 x2);
}

