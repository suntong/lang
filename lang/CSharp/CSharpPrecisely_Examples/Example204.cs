// Example 204 from page 169 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

// Class constraint -- permits use of null

class C1<T> where T : class {
  T f = null;                           // Legal: T is a reference type
}

// class C2<T> {
//   T f = null;                        // Illegal: T could be a value type
// }


// Struct constraint -- permits use of U in U?, a nullable type

class D1<U> where U : struct {
  U? f;                                 // Legal: U is a non-nullable value type
}

// class D2<U> {
//   U? f;                              // Illegal: U could be a reference type
// }

class MyTest {
  public static void Main(String[] args) { 
  }
}
