// Example 255 from page 217 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Attribute arguments are evaluated at compiletime; they must be
// constant expressions of a limited repertoire of types.

// Attribute constructors are executed at runtime, by applying them to
// the pre-evaluated argument values.  This happens at every call to
// GetCustomAttributes (in MS .Net 2.0 as well as Mono 1.0).

using System;                   // Attribute, AttributeUsage, AttributeTargets
using System.Reflection;        // MemberInfo

public enum Month {
  Jan=1, Feb, Mar, Apr, May, Jun, Jul, Aug, Sep, Oct, Nov, Dec
}

[AttributeUsage(AttributeTargets.Class | AttributeTargets.Method,
                AllowMultiple = true)]
class AuthorAttribute : Attribute { 
  public readonly String name;
  public readonly Month mm;

  public AuthorAttribute(String name, Month mm) { 
    this.name = name; this.mm = mm;
    Console.WriteLine("Creating AuthorAttribute: {0}", this);
  }

  public override String ToString() { 
    return String.Format("{0} ({1})", name, mm);
  }
}

class TestAttributes {
  [Author("Donald", Month.May)]
  public void MyMethod1() { }

  [Author("Andrzej", Month.Jul)]
  [Author("Andreas", Month.Mar)]
  public void MyMethod2() { }

  public static void Main(String[] args) {
    Type ty = typeof(TestAttributes);
    foreach (MemberInfo mif in ty.GetMembers()) {
      if (mif.Name.StartsWith("MyMethod")) {
        Console.WriteLine("\nGetting attributes of {0}:", mif.Name);
        Object[] attrs = mif.GetCustomAttributes(false);
        Console.WriteLine("\nThe attributes of {0} are:", mif.Name);
        foreach (Attribute attr in attrs) 
          Console.Write("{0} ", attr);
        Console.WriteLine();
        Console.WriteLine("\nGetting attributes of {0} again:", mif.Name);
        mif.GetCustomAttributes(false);
      }
    }
  }
}
