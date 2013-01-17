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
            TestSmtp();

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
        /// TestSmtp, Test C# Smtp -- use SmtpClient class to send email
        /// </summary>
        static void TestSmtp()
        {
            MailMessage message = new MailMessage();

            message.From = new MailAddress("noreply@dayforce.com", "Backfill Wizard");
            message.To.Add(new MailAddress("asun@Dayforce.com"));
            message.Subject = "Sending mail";
            message.Body = "Check sending email by Exchange from asp.net code <> ";

            SmtpClient client = new SmtpClient("dffs1", 25);

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
