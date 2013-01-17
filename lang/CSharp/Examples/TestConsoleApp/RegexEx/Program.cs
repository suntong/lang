using System;
using System.Text.RegularExpressions;

namespace RegexEx
{
    class Program
    {
        static void Main(string[] args)
        {
            Test1();
            Test2();
            Test3();
            Test_Replace();

            Console.WriteLine("\n== ");
            // Keep the console window open in debug mode.
            Console.WriteLine("Press any key to exit.");
            Console.ReadKey();
        }

        static void Test1()
        {
            // From http://www.knowdotnet.com/articles/regereplacementstrings.html
            
            /*
             * $&            - matched text
             * $_            - original source string
             * $`            - text before match 
             * $'            - text after match 
             * ${group_name} - text matched by named group 
             * $1, $2        - text matched by numbered group 
             * $$            - the literal "$"
             * 
             */

            Console.WriteLine("== Using Replacement Strings with Regex.Replace");
            string sInput = "Number 1.23 and 4.56";
            Console.WriteLine(sInput);

            // Put $ in front of monetary values
            sInput  = Regex.Replace(sInput,@"\d+\.\d{2}","$$$&");
            Console.WriteLine(sInput);

            sInput = "<d4p1:Payloads xmlns:d5p1=\"http://schemas.microsoft.com/2003/10/Serialization/Arrays\"><d5p1:base64Binary><GeneralGetRequest xmlns=\"http://SharpTop.Net/Services/Platform\"><MethodName>GetAllPayGroupBusinessDayInfo</MethodName><SessionTicket>qapayroll5:3ceb6f8b-157a-47f7-9890-23a310083486</SessionTicket></GeneralGetRequest></d5p1:base64Binary></d4p1:Payloads>";
            sInput = Regex.Replace(sInput, @"<(CoreServiceGenericParameters xmlns=""http://Dayforce/Data/CoreService""|GeneralGetRequest xmlns=""http://SharpTop.Net/Services/Platform"")",
                            @"$& xmlns:i=""http://www.w3.org/2001/XMLSchema-instance""");
            Console.WriteLine(sInput);

        }

        static void Test2()
        {
            // From http://www.dotnetperls.com/regex-replace

            Console.WriteLine("\n== Program that uses Regex.Replace method");
            
            // This is the input string we are replacing parts from.
            string input = "Dot Net Not Perls";

            // Use Regex.Replace to replace the pattern in the input.
            // ... The pattern N.t indicates three letters, N, any character, and t.
            string output = Regex.Replace(input, "N.t", "NET");

            // Write the output.
            Console.WriteLine(input);
            Console.WriteLine(output);

            Console.WriteLine("\n== Program that capitalizes strings");
            // Input strings.
            const string s1 = "samuel allen";
            const string s2 = "dot net perls";
            const string s3 = "Mother teresa";

            // Write output strings.
            Console.WriteLine(TextTools.UpperFirst(s1));
            Console.WriteLine(TextTools.UpperFirst(s2));
            Console.WriteLine(TextTools.UpperFirst(s3));
        }

        public static class TextTools
        {
            /// <summary>
            /// Uppercase first letters of all words in the string.
            /// </summary>
            /// 

            /*
             * use the delegate(Match match) syntax for a private method that alters strings to have an uppercase first letter. 
             * Delegate methods are methods you can use as variables and parameters.
             * Delegate Tutorial
             *  http://www.dotnetperls.com/delegate
             */

            public static string UpperFirst(string s)
            {
                return Regex.Replace(s, @"\b[a-z]\w+", delegate(Match match)
                {
                    string v = match.ToString();
                    return char.ToUpper(v[0]) + v.Substring(1);
                });
            }
        }


        static void Test3()
        {
            // From http://www.dijksterhuis.org/regular-expressions-csharp-practical-use/

            Console.WriteLine("\n== String Comparison – finding valid HTML tags");
            {
                string Input = "apples make for great party accessories";
                Regex FindA = new Regex("a");

                foreach (Match Tag in FindA.Matches(Input))
                {
                    Console.WriteLine("Found 'a' at {0}", Tag.Index);
                }
            }

            Console.WriteLine("\n== Searches for all valid HTML tags ");
            {
                Regex HTMLTag = new Regex(@"(<\/?[^>]+>)");

                string Input = "<b><i><a href='http://apple.com'>Ipod News</a></b></i>";

                foreach (Match Tag in HTMLTag.Matches(Input))
                {
                    Console.WriteLine("Found {0}", Tag.Value);
                }
                /*
                 * Result:
                 * 
                 * Found <b>
                 * Found <i>
                 * Found <a href=’http://apple.com’>
                 * Found </a>
                 * Found </b>
                 * Found </i>
                 *                  
                 */
            }

            Console.WriteLine("\n== Splitting a string into parts");
            {
                string IPMatchExp = @"(\d{1,3})\.(\d{1,3})\.(\d{1,3})\.(\d{1,3})";
                Match theMatch = Regex.Match("10.0.0.6", IPMatchExp);

                if (theMatch.Success)
                {
                    Console.WriteLine("{0}.{1}.{2}.{3}", theMatch.Groups[1].Value,
                                                          theMatch.Groups[2].Value,
                                                          theMatch.Groups[3].Value,
                                                          theMatch.Groups[4].Value);
                }
            }

            Console.WriteLine("\n== String Replacement");
            {
                Regex Replacer = new Regex(@"\w "); // Single [a-zA-Z] followed by a space
                string Input = "ax bx sax dam pom";
                string Output = Replacer.Replace(Input, "b_"); // Replace all items found with a b and underscore
                Console.WriteLine(Output);
            }

            Console.WriteLine("\n== Substitution Patterns");
            {
                Regex Replacer = new Regex(@"(\w*) (\w*)");
                string Input = "Molly Mallone";
                string Output = Replacer.Replace(Input, "$2 $1");
                Console.WriteLine(Output);
            }

        }

        static void Test_Replace()
        {
            Console.WriteLine("\n== Advanced Regex Replace");
            
            string IP = "Server:  UnKnown\r\nAddress:  192.168.25.3\r\n\r\nName:    CANWS02.dayforce.com\r\nAddress:  192.168.17.25\r\n\r\n"; // ExecuteCommand("nslookup", strHostName);
            IP = Regex.Replace(IP, @".*Name:.*?\.com", "", RegexOptions.Singleline);
            // to use multiple RegexOptions, just OR them together
            IP = Regex.Replace(IP, @".*address: *", "", RegexOptions.Singleline | RegexOptions.IgnoreCase);
            IP = Regex.Replace(IP, @"[\r\n]", ""); // globle is default
            Console.WriteLine("IP: " + IP);


            Console.WriteLine("\n== Program that uses Regex to Replace with an expression");

            // This is the input string we are replacing parts from.
            string input = "Dot 0 1 2 Perls";
            string[] lookup = new string[] { "Net", "Not", "Nut" };

            // Use Regex.Replace to replace the pattern in the input.
            // ... The pattern N.t indicates three letters, N, any character, and t.
            string output = Regex.Replace(input, @"( )(\d)( )", match =>
                match.Groups[1].ToString() + lookup[Convert.ToInt32(match.Groups[2].ToString())] + match.Groups[3].ToString());

            // Write the output.
            Console.WriteLine(input);
            Console.WriteLine(output);
        }

    }
}
