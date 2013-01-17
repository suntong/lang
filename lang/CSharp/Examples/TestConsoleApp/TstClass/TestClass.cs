using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace TstClass
{
    // Auto-Implemented Properties
    // http://msdn.microsoft.com/en-us/library/bb384054.aspx

    // This class is mutable. Its data can be modified from outside the class. 
    class Customer
    {
        // Auto-Impl Properties for trivial get and set 
        public double TotalPurchases { get; set; }
        public string Name { get; set; }
        public int CustomerID { get; set; }

        // Constructor 
        public Customer(double purchases, string name, int ID)
        {
            TotalPurchases = purchases;
            Name = name;
            CustomerID = ID;
        }
        // Methods 
        public string GetContactInfo() { return "ContactInfo"; }
        public string GetTransactionHistory() { return "History"; }

        // .. Additional methods, events, etc.
    }

    class StaticCustomer
    {
        // Auto-Impl Properties for trivial get and set 
        public static double TotalPurchases { get; set; }
        public static string Name { get; set; }
        public static int CustomerID { get; set; }

        // Constructor 
        static StaticCustomer()
        {
            TotalPurchases = 4987.63;
            Name = "Northwind";
            CustomerID = 90108;
        }

        // .. Additional methods, events, etc.
    }

    // How to: Write a Copy Constructor (C# Programming Guide)
    // http://msdn.microsoft.com/en-us/library/ms173116.aspx

    class Person
    {
        private string name;
        public int age;
        public B b;

        // Copy constructor. 
        public Person(Person previousPerson)
        {
            name = previousPerson.name;
            age = previousPerson.age;
            b = previousPerson.b;
        }

        //// Alternate copy contructor calls the instance constructor. 
        //public Person(Person previousPerson) 
        //    : this(previousPerson.name, previousPerson.age) 
        //{ 
        //} 

        // Also, ref
        // Copy a class in C#
        // http://stackoverflow.com/questions/1031023/copy-a-class-c-sharp

        // Instance constructor. 
        public Person(string name, int age, string thebase)
        {
            this.name = name;
            this.age = age;
            b = new B(); //instance for b which is part of object a. This is reference type that is in A
            this.b.j = thebase;
        }

        // Get accessor. 
        public string Details
        {
            get
            {
                return name + " is " + age.ToString() + " at " + b.j;
            }
        }
    }

    class B
    {
        public string j = "";

        public object Clone()
        {
            return this.MemberwiseClone();

        }
    }

    class TestPerson
    {
        public static void TestIt()
        {
            // Create a new person object.
            Person person1 = new Person("George", 40, "Camp1");
            Person p1 = person1;

            // Create another new object, copying person1.
            Person person2 = new Person(person1);
            Console.WriteLine(person2.Details);
            // Output: George is 40 at Camp1

            // Change person2, will it affect person1?
            person2.age = 50;
            person2.b.j = "Camp2";
            Console.WriteLine(p1.Details);
        }
    }

    /*
     * Generic C# Copy Constructor
     * http://stackoverflow.com/questions/433926/generic-c-sharp-copy-constructor
     * 
     * > What would be the best way to write a generic copy constructor function for my c# classes? 
     * 
     * Here's a constructor that I'm using. Note that this is a shallow constructor, and rather simplistic, due to the nature of my base
     * class. Should be good enough to get you started.
     * 

    public partial class LocationView : Location
    {
        public LocationView() { }

        // base class copy constructor
        public LocationView(Location value)
        {
            Type t = typeof(Location);
            PropertyInfo[] properties = t.GetProperties();
            foreach (PropertyInfo pi in properties)
            {
                pi.SetValue(this, pi.GetValue(value, null), null);
            }
        }
        public Quote quote { get; set; }
    }

     */

    // .NET Concepts Simplified
    // http://seesharpconcepts.blogspot.ca/2012/05/shallow-copy-and-deep-copy-in-c.html

    /*
     * Direct assignment
     * 
     * When copying one instance to another there are several ways of doing it. If we try to assign an object to another object, only 
     * reference gets copied and hence both objects will point to same memory.

Obj a = new  Obj();
Obj b = a; //a and b will have same reference

     * Hence if you modify anything in one object will modify another. It is because in memory it is same memory, 
     * same reference assigned to two variables.
     * 
     * Shallow Copy: 
     * 
     * Shallow copy is the way copying an object's value type fields bit by bit into target object and object's reference types are copied as
     * references into the target object but not the referenced object itself. This can be done in C# using MemberwiseClone() method on an
     * object. As in MSDN, "The MemberwiseClone method creates a shallow copy by creating a new object, and then copying the nonstatic fields
     * of the current object to the new object."

     * This cloned object will contain all the values of value types that are in source object and references are copied for reference types 
     * that are in source object. i.e. copying all fields that value type bit by bit copy and for all reference type fields only references 
     * are copied but not the referred object. Lets look into below example:
     * 
     */

    class A : ICloneable
    {
        public string i = "";
        public B b;

        #region ICloneable Members

        public object Clone()
        {
            return this.MemberwiseClone();

        }

        #endregion
    }

    class Program
    {

        static void Main(string[] args)
        {
            // Intialize a new object.
            Customer cust1 = new Customer(4987.63, "Northwind", 90108);

            //Modify a property
            cust1.TotalPurchases += 499.99;

            Console.WriteLine("StaticCustomer's name is '" + StaticCustomer.Name + "'");

            TestPerson.TestIt();

            // With above two classes let's try to do a shallow copy

            A a = new A(); //new instance creation of type A
            a.b = new B(); //instance for b which is part of object a. This is reference type that is in A
            a.b.j = "test"; //assign values for variables inside instance b which is contained in instance a
            a.i = "test"; //i is the value type here, assigning the value to the value type inside A

            A a1 = a.Clone() as A;        //Shallow copy using this.MemberwiseClone() in A's instance 

            a1.i = "asas"; //Change the value in cloned copy for value type
            a1.b.j = "changed"; //change value in cloned copy for a variable inside reference type of A's instance

            // To prove references are copied for reference type not just value copy, we will change the value a1.b.j to some other string
            // "changed". This changes the value of a.b.j as well. This is because when clone is created , all members are cloned and 
            // references are copied as it is. Hence a.b has reference as a1.b.

            // Ref: Shallow Copy vs. Deep Copy in .NET
            // http://www.codeproject.com/Articles/28952/Shallow-Copy-vs-Deep-Copy-in-NET

            // Keep the console window open in debug mode.
            Console.WriteLine("Press any key to exit.");
            Console.ReadKey();
        }
    }
}
