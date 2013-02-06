////////////////////////////////////////////////////////////////////////////
// Porgram: StdIO
// Purpose: A demo of C# IO
// Authors: Tong Sun (c) 2013, All rights reserved
////////////////////////////////////////////////////////////////////////////

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

using System.Net.Mail;

namespace StdIO
{
    class Program
    {
        static void Main(string[] args)
        {
            TestStdIO();
            TestIOSeq1();
            TestIOSeq2();
            //TestSmtp();

            // Keep the console window open in debug mode.
            Console.WriteLine("Press any key to exit.");
            Console.ReadKey();
        }

        /// <summary>
        /// TestStdIO, Test C# stdout & stderr
        /// </summary>
        static void TestStdIO()
        {
            Console.WriteLine("stdout: " + Guid.NewGuid().ToString()
                + System.IO.Path.GetTempFileName() + ".bat");
            Console.Error.WriteLine("stderr: " + DateTime.Today);
        }

        /// <summary>
        /// TestIOSeq1, Test C# stdout & stderr output sequent 
        /// </summary>
        static void TestIOSeq1()
        {
            for (int i = 0; i < 5; i++)
            {
                Console.WriteLine("Hello");
            }
            Console.Error.WriteLine("Bah bah");
            for (int i = 0; i < 5; i++)
            {
                Console.WriteLine("2");
            }
        }

        /// <summary>
        /// TestIOSeq2, Test C# stdout & stderr output sequent, with flushing
        /// </summary>
        static void TestIOSeq2()
        {
            for (int i = 0; i < 5; i++)
            {
                Console.WriteLine("Hello");
            }
            Console.Out.Flush();
            Console.Error.WriteLine("Bah bah");
            Console.Error.Flush();
            for (int i = 0; i < 5; i++)
            {
                Console.WriteLine("2");
            }
            Console.Out.Flush();
        }

        /// <summary>
        /// TestSmtp, Test C# Smtp -- use SmtpClient class to send email
        /// </summary>
        static void TestSmtp()
        {
            MailMessage message = new MailMessage();

            message.From = new MailAddress(Properties.Settings.Default.email_noreply, 
                                            Properties.Settings.Default.email_name);
            message.To.Add(new MailAddress(Properties.Settings.Default.email_to));
            message.Subject = "Sending mail";
            message.Body = "Check sending email by Exchange from asp.net code <> ";

            SmtpClient client = new SmtpClient(Properties.Settings.Default.email_smtp, 25);

            try
            {
                client.Send(message);
            }
            catch (Exception exc)
            {
                Console.Error.WriteLine(exc.Message.ToString());
            }
        }

    }
}
