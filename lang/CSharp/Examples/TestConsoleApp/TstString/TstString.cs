using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

using TestTuple = System.Tuple<int, string>;

// Implement the Left/Right extension method to String, which returns the left/rightmost part of a string.
// http://www.dotnetperls.com/right
static class Extensions
{
    /// <summary>
    /// Get substring of specified number of characters on the right.
    /// </summary>
    public static string Right(this string value, int length)
    {
        length = value.Length - length;
        if (length < 0) length = 0;
        return value.Substring(length);
    }

    /// <summary>
    /// Get substring of specified number of characters on the left.
    /// </summary>
    public static string Left(this string value, int length)
    {
        return value.Substring(0, length > value.Length ? value.Length : length);
    }
}

namespace TstString
{
    class TstString
    {
        static void Main(string[] args)
        {
            TestString();
            TestStringAL();
            TestDictionary();
            TestDictionaryClass();
            TestTuple();
            TestSSN();

            // Keep the console window open in debug mode.
            Console.WriteLine("\nPress any key to exit.");
            Console.ReadKey();
        }

        static void TestString()
        {
            Console.WriteLine("\n== Test string operations");

            string sqlCommand = string.Format(
                "INSERT INTO event\n (TmStamp, UserName, Action, SvrName, LogFile, Commnt)\n VALUES  (CURRENT_TIMESTAMP, '{0}', '{1}', '{2}', '{3}.txt', '{4}' )", "User", "Task", "P_server", "tmpFileN", "taskDetail");
            Console.WriteLine(sqlCommand);

            String s = "Andrew was here";
            String b = "And";
            Console.WriteLine(s.Length);                // 15
            Console.WriteLine(s.Remove(0, b.Length));   // rew was here
            Console.WriteLine(s.IndexOf("drew"));       // 2
            Console.WriteLine(s.IndexOf("A"));          // 0
            Console.WriteLine(s.IndexOf("X"));          // -1
            Console.WriteLine(s.Substring(2, 3));       // dre

            String p = s.Substring(0, s.IndexOf("as")); // Andrew w;
            Console.WriteLine(p);
            p = s.Substring(p.Length);                  //as here
            Console.WriteLine(p);
            Console.WriteLine(p.Substring(0, p.IndexOf("ere"))); // as h
            Console.WriteLine(p.Right(30));                  // as here
            Console.WriteLine(p.Left(30));                  // as here
            Console.WriteLine();

            Console.WriteLine(s.Right(3));                  // ere
            Console.WriteLine(s.Left(5));                   // Andre
            Console.WriteLine(s.Substring(0, s.Length - 3));   // ... was h
            Console.WriteLine(s.Left(s.Length - 3));        // ... was h
            Console.WriteLine();

            // ReplaceTest 
            string errString = "This docment uses 3 other docments to docment the docmentation";
            Console.WriteLine("The original string is:{0}'{1}'{0}", Environment.NewLine, errString);
            string correctString = errString.Replace("docment", "document");
            Console.WriteLine("After correcting the string, the result is:{0}'{1}'",
                    Environment.NewLine, correctString);
        }

        static void TestStringAL()
        {
            Console.WriteLine("\n== Test string array and list");

            // == string array testing
            String[] x = new string[] { "item1", "item2", "item3" };

            // Appending strings to a string array
            // http://stackoverflow.com/questions/3693167/c-sharp-appending-strings-to-a-string-array

            var list = new List<string>();
            list.Add("any string 1");
            list.Add("any string 2");
            String[] y = list.ToArray();

            // concatenate two arrays 
            // http://stackoverflow.com/questions/1547252/how-do-i-concatenate-two-arrays-in-c
            // method 1
            var z = new string[x.Length + y.Length];
            x.CopyTo(z, 0);
            y.CopyTo(z, x.Length);
            // method 2
            //List<string> list = new List<string>();
            list = new List<string>();
            list.AddRange(x);
            list.AddRange(y);
            String[] z2 = list.ToArray();

            Console.WriteLine("\n== Test string list");

            // == string list testing
            // http://www.dotnetperls.com/list

            // Loop through List
            Console.WriteLine("foreach loop");
            foreach (string lc in list) // Loop through List with foreach
            {
                Console.WriteLine(lc);
            }

            // Reverse List in-place, no new variables required
            list.Reverse();

            Console.WriteLine("for loop");
            for (int i = 0; i < list.Count; i++) // Loop through List with for
            {
                Console.WriteLine(list[i]);
            }

            Console.WriteLine(list.Count); // 5
            list.Clear();
            Console.WriteLine(list.Count); // 0

            // Copy array to List
            list = new List<string>(x);
            Console.WriteLine(list.Count);

            // Join string List
            List<string> cities = new List<string>();
            cities.Add("New York");
            cities.Add("Mumbai");
            cities.Add("Berlin");
            cities.Insert(1, "Istanbul"); // index is 0 based, This makes it become the second element in the List. 

            // Join strings into one CSV line
            string line = string.Join(", ", cities.ToArray());
            Console.WriteLine(line);

            // Range of elements
            List<string> rivers = new List<string>(new string[]
	            {
	                "nile",
	                "amazon",     // River 2
	                "yangtze",    // River 3
	                "mississippi",
	                "yellow"
	            });

            // Get rivers 2 through 3
            List<string> range = rivers.GetRange(1, 2);
            foreach (string river in range)
            {
                Console.WriteLine(river);
            }

        }


        static void TestDictionary()
        {
            Console.WriteLine("\n== Test Dictionary");
            // http://www.dotnetperls.com/dictionary

            Dictionary<string, int> dictionary = new Dictionary<string, int>();
            dictionary.Add("apple", 1);
            dictionary.Add("windows", 5);

            // See whether Dictionary contains this string.
            if (dictionary.ContainsKey("apple"))
            {
                int value = dictionary["apple"];
                Console.WriteLine(value);
            }
            // See whether Dictionary contains this string.
            if (!dictionary.ContainsKey("acorn"))
            {
                Console.WriteLine(false);
            }
            if (dictionary.ContainsValue(5))
            {
                Console.WriteLine(true); // true
            }
            dictionary.Remove("windows"); // Removes windows
            dictionary.Remove("nothing"); // Doesn't remove anything
            Console.WriteLine(dictionary.ContainsValue(5));
            Console.WriteLine();

            // Example Dictionary again
            Dictionary<string, int> d = new Dictionary<string, int>()
            {
	            {"llama", 0},
	            {"cat", 2},
	            {"iguana", -1},
	            {"dog", 1},
            };

            // Loop over pairs with foreach
            foreach (KeyValuePair<string, int> pair in d)
            {
                Console.WriteLine("KeyValuePair: {0}, {1}",
                pair.Key,
                pair.Value);
            }
            Console.WriteLine();

            // Use var keyword to enumerate dictionary
            foreach (var pair in d)
            {
                Console.WriteLine("var: {0}, {1}",
                pair.Key,
                pair.Value);
            }
            Console.WriteLine();

            // Store keys in a List
            {
                List<string> list = new List<string>(d.Keys);
                // Loop through list
                foreach (string k in list)
                {
                    Console.WriteLine("list: {0}, {1}",
                    k,
                    d[k]);
                }
            }
            Console.WriteLine();

            // Copy the Dictionary using the Copy constructor 
            dictionary = new Dictionary<string, int>(d);

            // Sort Dictionary 
            {
                // Acquire keys and sort them.
                var list = dictionary.Keys.ToList();
                list.Sort();

                // Loop through keys.
                foreach (var key in list)
                {
                    Console.WriteLine("sorted: {0}: {1}", key, dictionary[key]);
                }
            }
        }


        public class MySubClass
        {
            public string Value { get; set; }

            // Copy constructor. 
            public MySubClass(MySubClass _mySubClass)
            {
                Value = _mySubClass.Value;
            }

            // Instance constructor. 
            public MySubClass(string _Value)
            {
                Value = _Value;
            }
        }
        
        public class MyClass
        {
            private MySubClass mySubClass = new MySubClass("");

            private static Dictionary<string, MySubClass> dictionary;

            public void some_proc()
            {
                dictionary = new Dictionary<string, MySubClass>();
                mySubClass.Value = "foo";
                dictionary.Add("apple", mySubClass);
                dictionary.Add("apple1", new MySubClass(mySubClass));
                mySubClass.Value = "bar";
                dictionary.Add("orange", mySubClass);
                dictionary.Add("orange2", new MySubClass(mySubClass));
            }

            public void some_other_proc()
            {
                mySubClass = dictionary["apple"];
                dictionary.Remove("apple");
                string key = "apple1";
                mySubClass = dictionary[key];
                dictionary.Remove(key);
            }
        }

        static void TestDictionaryClass()
        {
            Console.WriteLine("\n== Test Dictionary Class");
            // http://social.msdn.microsoft.com/Forums/en-US/csharpgeneral/thread/1bb039f5-b09f-4bd4-814a-5de9ed2af934

            MyClass myClass = new MyClass();
            myClass.some_proc();
            myClass.some_other_proc();
        }

        static void TestTuple()
        {
            Console.WriteLine("\n== Test Tuple");

            /*
             * http://www.dotnetperls.com/tuple
             * 
             * A Tuple has many items. Each item can have any type. The Tuple type in the C# language provides 
             * a unified syntax for creating objects with typed fields. Once created, the fields in the Tuple 
             * cannot be mutated. This makes the Tuple similar to a value type.
             */

            {
                // Create three-item tuple.
                Tuple<int, string, bool> tuple = new Tuple<int, string, bool>(1,
                    "cat", true);
                // Access tuple properties.
                if (tuple.Item1 == 1)
                {
                    Console.WriteLine(tuple.Item1);
                }
                if (tuple.Item2 == "dog")
                {
                    Console.WriteLine(tuple.Item2);
                }
                if (tuple.Item3)
                {
                    Console.WriteLine(tuple.Item3);
                }
            }
            Console.WriteLine();

            {
                // Create four-item tuple; use var implicit type.
                var tuple = new Tuple<string, string[], int, int[]>("perl",
                    new string[] { "java", "c#" },
                    1,
                    new int[] { 2, 3 });
                // Pass tuple as argument.
                M(tuple);
            }

            // Sort Tuple List
            // http://www.dotnetperls.com/sort-tuple
            {
                //List<Tuple<int, string>> list = new List<Tuple<int, string>>();
                //list.Add(new Tuple<int, string>(1, "cat"));
                //list.Add(new Tuple<int, string>(100, "apple"));
                //list.Add(new Tuple<int, string>(2, "zebra"));

                // Using C#'s typedef instead via the above "using TestTuple = System.Tuple<int, string>;"
                List<TestTuple> list = new List<TestTuple>();
                list.Add(new TestTuple(1, "cat"));
                list.Add(new TestTuple(100, "apple"));
                list.Add(new TestTuple(2, "zebra"));

                // Use Sort method with Comparison delegate.
                // ... Has two parameters; return comparison of Item2 on each.
                list.Sort((a, b) => a.Item2.CompareTo(b.Item2));

                foreach (var element in list)
                {
                    Console.WriteLine(element);
                }
            }

        }

        
        static void M(Tuple<string, string[], int, int[]> tuple)
        {
            // Evaluate the tuple's items.
            Console.WriteLine(tuple.Item1);
            foreach (string value in tuple.Item2)
            {
                Console.WriteLine(value);
            }
            Console.WriteLine(tuple.Item3);
            foreach (int value in tuple.Item4)
            {
                Console.WriteLine(value);
            }

        }

        //==========================================================================

                
        static void TestSSN()
        {
            Console.WriteLine("\n== Test SSN");
             for (int x = 1; x <= 3; x++)
                Console.WriteLine(RandomSSN("9"));
            for (int x = 1; x <= 3; x++)
                Console.WriteLine(RandomSSN("98"));
            for (int x = 1; x <= 3; x++)
                Console.WriteLine(RandomSSN("987"));
            Console.WriteLine(RandomSSN("77","-"));
            Console.WriteLine(GenerateSSN("-"));
        }

        public static string RandomSSN(string thePrefix, string delimiter = "")
        {
            string generatedSSN = GenerateSSN(delimiter);
            return thePrefix + generatedSSN.Substring(thePrefix.Length);
        }

        public static string GenerateSSN(string delimiter)
        {
            int iThree = GetRandomNumber(132, 921);
            int iTwo = GetRandomNumber(12, 83);
            int iFour = GetRandomNumber(1423, 9211);
            return iThree.ToString() + delimiter + iTwo.ToString() + delimiter + iFour.ToString();
        }

        //Function to get random number
        private static readonly Random getrandom = new Random();
        public static int GetRandomNumber(int min, int max)
        {
            return getrandom.Next(min, max);
        }
    }
}
