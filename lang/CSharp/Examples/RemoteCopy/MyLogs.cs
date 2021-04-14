using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace RemoteCopy
{
    public static class MyLogs
    {
        public static void WriteLog(string methodName, string message, bool isError)
        {
            Console.ForegroundColor = ConsoleColor.DarkGreen;
            Console.WriteLine(methodName + " - " + message);

            if (isError == true)
            {
                //whatever
            }

            //write to a file
        }
    }
}
