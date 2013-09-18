using System;
using RJH.CommandLineHelper;

// By Ray Hayes, 1 May 2003
// http://www.codeproject.com/Articles/3852/Automatic-Command-Line-Parsing-in-C

namespace CommandLineSwitch
{
    /// <summary>The application call acts as a tester for the command line
    /// parser.  It demonstrates using switch attributes on properties
    /// meaning that the coder does not have to implement anything except
    /// instantiating the parser in the most basic way.</summary>

    class Application
    {
        #region Enumerations
        public enum DaysOfWeek
        {
            Sun,
            Mon,
            Tue,
            Wed,
            Thu,
            Fri,
            Sat
        };
        #endregion

        #region Private Variables
        private bool m_showHelp = true;
        private bool m_SomethingElse = false;
        private string m_SomeName = "My Name";
        private int m_Age = 999;
        private string m_test = "XXXX";
        private DaysOfWeek m_DoW = DaysOfWeek.Sun;
        #endregion

        #region Command Line Switches
        /// <summary>Simple example of a Boolean switch.</summary>
        [CommandLineSwitch("SomeHelp", "Show some additional help")]
        public bool ShowSomeHelp
        {
            get { return m_showHelp; }
            set { m_showHelp = value; }
        }

        /// <summary>Simple example of a Boolean switch.</summary>
        /// <remark>There is no get value set, so this value is in effect
        /// a write only one.  This can affect the implementation of the toggle
        /// for Boolean values.</remark>
        [CommandLineSwitch("SomethingElse", "Do something else")]
        public bool Wibble
        {
            set { m_SomethingElse = value; }
        }

        /// <summary>Simple example of a string switch with an alias.</summary>
        [CommandLineSwitch("Name", "User Name")]
        [CommandLineAlias("User")]
        public string UserName
        {
            get { return m_SomeName; }
            set { m_SomeName = value; }
        }

        /// <summary>Simple example of an integer switch.</summary>
        [CommandLineSwitch("Age", "User age")]
        public int Age
        {
            get { return m_Age; }
            set { m_Age = value; }
        }

        /// <summary>Simple example of a read-only, e.g. no writeback, Boolean
        /// command line switch.</summary>
        [CommandLineSwitch("Test", "Test switch")]
        public string Test
        {
            get { return m_test; }
        }

        [CommandLineSwitch("Day", "Day of the week selection")]
        [CommandLineAlias("DoW")]
        public DaysOfWeek DoW
        {
            get { return m_DoW; }
            set { m_DoW = value; }
        }
        #endregion

        #region Private Utility Functions
        private int Run(string[] cmdLine)
        {
            // Initialise the command line parser, passing in a reference to this
            // class so that we can look for any attributes that implement
            // command line switches.
            Parser parser = new Parser(System.Environment.CommandLine, this);

            // Programmatically add some switches to the command line parser.
            parser.AddSwitch("Wibble", "Do something silly");

            // Add a switches with lots of aliases for the first name, "help" and "a".
            parser.AddSwitch(new string[] { "help", @"\?" }, "show help");
            parser.AddSwitch(new string[] { "a", "b", "c", "d", "e", "f" }, "Early alphabet");

            // Parse the command line.
            parser.Parse();

            // ----------------------- DEBUG OUTPUT -------------------------------
            Console.WriteLine("Program Name      : {0}", parser.ApplicationName);
            Console.WriteLine("Non-switch Params : {0}", parser.Parameters.Length);
            for (int j = 0; j < parser.Parameters.Length; j++)
                Console.WriteLine("                {0} : {1}", j, parser.Parameters[j]);
            Console.WriteLine("----");
            Console.WriteLine("Value of ShowSomeHelp    : {0}", ShowSomeHelp);
            Console.WriteLine("Value of m_SomethingElse : {0}", m_SomethingElse);
            Console.WriteLine("Value of UserName        : {0}", UserName);
            Console.WriteLine("----");

            // Walk through all of the registered switches getting the available
            // information back out.
            Parser.SwitchInfo[] si = parser.Switches;
            if (si != null)
            {
                Console.WriteLine("There are {0} registered switches:", si.Length);
                foreach (Parser.SwitchInfo s in si)
                {
                    Console.WriteLine("Command : {0} - [{1}]", s.Name, s.Description);
                    Console.Write("Type    : {0} ", s.Type);

                    if (s.IsEnum)
                    {
                        Console.Write("- Enums allowed (");
                        foreach (string e in s.Enumerations)
                            Console.Write("{0} ", e);
                        Console.Write(")");
                    }
                    Console.WriteLine();

                    if (s.Aliases != null)
                    {
                        Console.Write("Aliases : [{0}] - ", s.Aliases.Length);
                        foreach (string alias in s.Aliases)
                            Console.Write(" {0}", alias);
                        Console.WriteLine();
                    }

                    Console.WriteLine("------> Value is : {0} (Without any callbacks {1})\n",
                        s.Value != null ? s.Value : "(Unknown)",
                        s.InternalValue != null ? s.InternalValue : "(Unknown)");
                }
            }
            else
                Console.WriteLine("There are no registered switches.");

            // Test looking for a specificly named values.
            Console.WriteLine("----");
            if (parser["help"] != null)
                Console.WriteLine("Request for help = {0}", parser["help"]);
            else
                Console.WriteLine("Request for help has no associated value.");
            Console.WriteLine("User Name is {0}", parser["name"]);

            // Note the difference between the parser and a callback value.
            Console.WriteLine("The property of test (/test) is internally is read-only, " +
                                    "e.g. no update can be made by the parser:\n" +
                                    "   -- The indexer gives a value of : {0}\n" +
                                    "   -- Internally the parser has    : {1}",
                                    parser["test"],
                                    parser.InternalValue("test"));

            // Test if the enumeration value has changed to Friday!
            if (DoW == DaysOfWeek.Fri)
                Console.WriteLine("\nYeah Friday.... PUB LUNCH TODAY...");

            // For error handling, were any switches handled?
            string[] unhandled = parser.UnhandledSwitches;
            if (unhandled != null)
            {
                Console.WriteLine("\nThe following switches were not handled.");
                foreach (string s in unhandled)
                    Console.WriteLine("  - {0}", s);
            }

            return 0;
        }
        #endregion

        private static int Main(string[] cmdLine)
        {
            Application app = new Application();
            return app.Run(cmdLine);
        }
    }
}
