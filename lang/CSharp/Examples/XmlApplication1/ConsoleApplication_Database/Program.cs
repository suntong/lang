using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

using System.Data;
using System.Data.SqlClient;



namespace ConsoleApplication_Database
{
    class Program
    {
        static void Main(string[] args)
        {


            string myServer = "ncdb10"; //"your_server_here";
            string myDB = "aeroperf1"; //"your_db_here";
            string UserId = "aeroperf1";
            string Password = "qtk7fqyg";
            //string queryString = "SELECT LoginId,UserId,IsLockedOut,IsApproved FROM AppUser WHERE IsLockedOut = 1 or IsApproved = 0;";
            string queryString = "UPDATE AppUser SET IsLockedOut = 0 WHERE IsLockedOut = 1;UPDATE AppUser SET IsApproved = 1 WHERE IsApproved = 0;Select LoginId,UserId,IsLockedOut,IsApproved from AppUser WHERE IsLockedOut = 1 and IsApproved = 0 ORDER BY LoginId;";

            string connectionString =
                       "Data Source=" + myServer + ";Initial Catalog=" + myDB + "; User ID=" + UserId + "; Password=" + Password;
                      // + "Integrated Security=true";



            // Create and open the connection in a using block. This
            // ensures that all resources will be closed and disposed
            // when the code exits.
            using (SqlConnection connection =
                new SqlConnection(connectionString))
            {
                // Create the Command and Parameter objects.
                SqlCommand command = new SqlCommand(queryString, connection);

                // Open the connection in a try/catch block. 
                // Create and execute the DataReader, writing the result
                // set to the console window.
                try
                {
                    connection.Open();
                    SqlDataReader reader = command.ExecuteReader();
                    while (reader.Read())
                    {
                        Console.WriteLine("\t{0}\t{1}\t{2}\t{3}",
                            reader[0], reader[1], reader[2], reader[3]);
                    }
                    reader.Close();
                }
                catch (Exception ex)
                {
                    Console.WriteLine(ex.Message);
                }
                Console.ReadLine();

            }
        }
    }
    }