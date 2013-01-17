// ReadDb.cs
// ------
// ------ From http://msdn.microsoft.com/en-us/library/ff965871.aspx
// ------------------------------------------------------------------


using System;
using System.Data;
using System.Data.OleDb;
using System.Data.SqlClient;
using System.Collections.Generic;

public class ReadDb
{
    public static void Main()
    {
        TestOleDB();
        TestSqlDBL();
        TestSqlDB2();
        TestSqlDBR();
        TestSqlAsOleDB();
        TestDictionary();
        TestSqlSP();

        // Keep the console window open in debug mode.
        Console.WriteLine("Press any key to exit.");
        Console.ReadKey();
    }

    /// <summary>
    /// TestDB, Test C# DB connection
    /// </summary>
    static void TestOleDB()
    {
        string DAM = "[ReadOleDb]";

        string myConnectionString =
                            "Provider=Microsoft.ACE.OLEDB.12.0;" +
                            "User Id=;Password=;" +
                        @"Data Source=D:\Projects\DBs\Northwind.accdb";

        Console.WriteLine("== Test OleDB\n");

        // Connection string for ADO.NET via OleDB
        OleDbConnection cn =
            new OleDbConnection(myConnectionString);

        // Prepare SQL query
        string query = "SELECT Customers.[Company], Customers.[First Name] FROM Customers ORDER BY Customers.[Company] ASC;";
        OleDbCommand cmd = new OleDbCommand(query, cn);

        try
        {
            cn.Open();
            Console.WriteLine("{0}: Successfully connected to database. Data source name:\n {1}",
                DAM, cn.DataSource);
            Console.WriteLine("{0}: SQL query:\n {1}", DAM, query);

            // Run the query and create a record set
            OleDbDataReader dr = cmd.ExecuteReader();
            Console.WriteLine("{0}: Retrieve schema info for the given result set:", DAM);
            for (int column = 0; column < dr.FieldCount; column++)
            {
                Console.Write(" | {0}", dr.GetName(column));
            }
            Console.WriteLine("\n{0}: Fetch the actual data: ", DAM);
            int row = 0;
            while (dr.Read())
            {
                Console.WriteLine(" | {0} | {1} ", dr.GetValue(0), dr.GetValue(1));
                row++;
            }
            Console.WriteLine("{0}: Total Row Count: {1}", DAM, row);
            dr.Close();
        }
        catch (OleDbException ex)
        {
            Console.WriteLine("{0}: OleDbException: Unable to connect or retrieve data from data source: {1}.",
                DAM, ex.ToString());
        }
        catch (Exception ex)
        {
            Console.WriteLine("{0}: Exception: Unable to connect or retrieve data from data source: {1}.",
                DAM, ex.ToString());
        }
        finally
        {
            cn.Close();
            Console.WriteLine("{0}: Cleanup. Done.", DAM);
        }
    }

    /// <summary>
    /// TestSqlDBL, Test C# Local DB connection
    /// </summary>
    static void TestSqlDBL()
    {
        string DAM = "[ReadSqlDb]";

        string myConnectionString =
                    "Data Source=(local);Initial Catalog=Backfill;Integrated Security=True";

        Console.WriteLine("\n\n== Test SqlDB\n");

        // Connection string for ADO.NET 
        SqlConnection cn =
            new SqlConnection(myConnectionString);

        // Prepare SQL query
        string query = "SELECT svr_name, commnt FROM db_db_svr;";
        SqlCommand cmd = new SqlCommand(query, cn);

        try
        {
            cn.Open();
            Console.WriteLine("{0}: Successfully connected to database. Data source name:\n {1}",
                DAM, cn.DataSource);
            Console.WriteLine("{0}: SQL query:\n {1}", DAM, query);

            // Run the query and create a record set
            SqlDataReader dr = cmd.ExecuteReader();
            Console.WriteLine("{0}: Retrieve schema info for the given result set:", DAM);
            for (int column = 0; column < dr.FieldCount; column++)
            {
                Console.Write(" | {0}", dr.GetName(column));
            }
            Console.WriteLine("\n{0}: Fetch the actual data: ", DAM);
            int row = 0;
            while (dr.Read())
            {
                Console.WriteLine(" | {0} | {1} ", dr.GetValue(0), dr.GetValue(1));
                row++;
            }
            Console.WriteLine("{0}: Total Row Count: {1}", DAM, row);
            dr.Close();
        }
        catch (SqlException ex)
        {
            Console.WriteLine("{0}: SqlException: Unable to connect or retrieve data from data source: {1}.",
                DAM, ex.ToString());
        }
        catch (Exception ex)
        {
            Console.WriteLine("{0}: Exception: Unable to connect or retrieve data from data source: {1}.",
                DAM, ex.ToString());
        }
        finally
        {
            cn.Close();
            Console.WriteLine("{0}: Cleanup. Done.", DAM);
        }
    }

    /// <summary>
    /// TestSqlDB2, Test C# DB SqlConnection
    /// </summary>
    static void TestSqlDB2()
    {

        string department = "DP3-WEB";
        string msConnectionString =
                    "Data Source=(local);Initial Catalog=Demo1;Integrated Security=True";
        
        const string sql =
        @"SELECT
              emp_name,
              department
          FROM
              employee
          WHERE
              department = @department
        ";

        using (SqlConnection con = new SqlConnection(msConnectionString))
        using (SqlCommand cmd = con.CreateCommand())
        {
            con.Open();
            cmd.CommandText = sql;

            cmd.Parameters.AddWithValue("@department", department);

            using (SqlDataReader dr = cmd.ExecuteReader())
            {
                int row = 0;
                while (dr.Read())
                {
                    Console.WriteLine(" | {0} | {1} ", dr.GetValue(0), dr.GetValue(1));
                    row++;
                }
                Console.WriteLine("Total Row Count: {0}", row);
            }
        }

    }

    /// <summary>
    /// TestSqlDBR, Test C# Remote DB connection
    /// </summary>
    static void TestSqlDBR()
    {
        string DAM = "[ReadSqlDb]";

        string myConnectionString =
                    "Data Source=Torsvdb04;Initial Catalog=QAPAYROLL5;uid=wbpoc;pwd=sql@tfs2008;";

        Console.WriteLine("\n\n== Test SqlDB\n");

        // Connection string for ADO.NET 
        SqlConnection cn =
            new SqlConnection(myConnectionString);

        // Prepare SQL query
        string query = "SELECT ClientId, Namespace FROM Client;";
        SqlCommand cmd = new SqlCommand(query, cn);

        try
        {
            cn.Open();
            Console.WriteLine("{0}: Successfully connected to database. Data source name:\n {1}",
                DAM, cn.DataSource);
            Console.WriteLine("{0}: SQL query:\n {1}", DAM, query);

            // Run the query and create a record set
            SqlDataReader dr = cmd.ExecuteReader();
            Console.WriteLine("{0}: Retrieve schema info for the given result set:", DAM);
            for (int column = 0; column < dr.FieldCount; column++)
            {
                Console.Write(" | {0}", dr.GetName(column));
            }
            Console.WriteLine("\n{0}: Fetch the actual data: ", DAM);
            int row = 0;
            while (dr.Read())
            {
                Console.WriteLine(" | {0} | {1} ", dr.GetValue(0), dr.GetValue(1));
                row++;
            }
            Console.WriteLine("{0}: Total Row Count: {1}", DAM, row);
            dr.Close();
        }
        catch (SqlException ex)
        {
            Console.WriteLine("{0}: SqlException: Unable to connect or retrieve data from data source: {1}.",
                DAM, ex.ToString());
        }
        catch (Exception ex)
        {
            Console.WriteLine("{0}: Exception: Unable to connect or retrieve data from data source: {1}.",
                DAM, ex.ToString());
        }
        finally
        {
            cn.Close();
            Console.WriteLine("{0}: Cleanup. Done.", DAM);
        }
    }

    /// <summary>
    /// TestSqlAsOleDB, Test C# DB connection as OleDB
    /// </summary>
    static void TestSqlAsOleDB()
    {
        string DAM = "[ReadOleDb]";

        string myConnectionString =
                            "Provider=SQLOLEDB;Data Source=(local);Initial Catalog=Backfill;Integrated Security=SSPI";

        Console.WriteLine("\n\n== Test Sql as OleDB\n");

        // Connection string for ADO.NET via OleDB
        OleDbConnection cn =
            new OleDbConnection(myConnectionString);

        // Prepare SQL query
        string query = "SELECT svr_name, commnt FROM db_db_svr;";
        OleDbCommand cmd = new OleDbCommand(query, cn);

        try
        {
            cn.Open();
            Console.WriteLine("{0}: Successfully connected to database. Data source name:\n {1}",
                DAM, cn.DataSource);
            Console.WriteLine("{0}: SQL query:\n {1}", DAM, query);

            // Run the query and create a record set
            OleDbDataReader dr = cmd.ExecuteReader();
            Console.WriteLine("{0}: Retrieve schema info for the given result set:", DAM);
            for (int column = 0; column < dr.FieldCount; column++)
            {
                Console.Write(" | {0}", dr.GetName(column));
            }
            Console.WriteLine("\n{0}: Fetch the actual data: ", DAM);
            int row = 0;
            while (dr.Read())
            {
                Console.WriteLine(" | {0} | {1} ", dr.GetValue(0), dr.GetValue(1));
                row++;
            }
            Console.WriteLine("{0}: Total Row Count: {1}", DAM, row);
            dr.Close();
        }
        catch (OleDbException ex)
        {
            Console.WriteLine("{0}: OleDbException: Unable to connect or retrieve data from data source: {1}.",
                DAM, ex.ToString());
        }
        catch (Exception ex)
        {
            Console.WriteLine("{0}: Exception: Unable to connect or retrieve data from data source: {1}.",
                DAM, ex.ToString());
        }
        finally
        {
            cn.Close();
            Console.WriteLine("{0}: Cleanup. Done.", DAM);
        }
    }

    /// <summary>
    /// TestDictionary, Test C# Dictionary Class
    /// </summary>
    static void TestDictionary()
    {
        Console.WriteLine("\n\n== Test C# Dictionary Class\n");

        // Example Dictionary again
        Dictionary<string, int> d = new Dictionary<string, int>()
        {
	        {"cat", 2},
	        {"dog", 1},
	        {"llama", 0},
	        {"iguana", -1}
        };

        // can also use Dictionary Add method 
        d.Add("apple", 1);
        d.Add("windows", 5);

        // Loop over pairs with foreach
        foreach (KeyValuePair<string, int> pair in d)
        {
            Console.WriteLine("{0}, {1}",
            pair.Key,
            pair.Value);
        }
        Console.WriteLine();

        // Use var keyword to enumerate dictionary
        foreach (var pair in d)
        {
            Console.WriteLine("{0}, {1}",
            pair.Key,
            pair.Value);
        }
    }

    /// <summary>
    /// TestSqlSP, Test C# to get SQL Server stored procedure’s return value and output value
    /// </summary>
    static void TestSqlSP()
    {
        string DAM = "[ReadSqlSP]";

        string myConnectionString =
                    "Data Source=perfdb01;Initial Catalog=perfdb;uid=wbpoc;pwd=sql@tfs2008;";

        Console.WriteLine("\n\n== Test SqlSP\n");

        string _SqlCommand = "Select 1";

        using (SqlConnection connection = new SqlConnection(myConnectionString))
        {
            connection.Open();
            using (SqlCommand sqlCommand = new SqlCommand(_SqlCommand, connection))
            {
                sqlCommand.ExecuteNonQuery();
            }
        }

        using (SqlConnection conn = new SqlConnection(myConnectionString))
        using (SqlCommand cmd = new SqlCommand("rvtest", conn))
        {
            //cmd.CommandText = "rvtest";
            cmd.CommandType = CommandType.StoredProcedure;
            cmd.Parameters.AddWithValue("@a", "23");

            var bParameter = cmd.Parameters.Add("@b", SqlDbType.Int);
            bParameter.Direction = ParameterDirection.Output;
            var returnParameter = cmd.Parameters.Add("returnParameter", SqlDbType.Int);
            returnParameter.Direction = ParameterDirection.ReturnValue;

            conn.Open();
            cmd.ExecuteNonQuery();
            var result = returnParameter.Value;
            var bval = bParameter.Value;

            Console.WriteLine(" | {0} | {1} ", result, bval);
        }

        // Best approach
        //using (var conn = new SqlConnection(connectionString))
        //using (var command = new SqlCommand("ProcedureName", conn)
        //{
        //    CommandType = CommandType.StoredProcedure
        //})
        //{
        //    conn.Open();
        //    command.ExecuteNonQuery();
        //    conn.Close();
        //}

        Console.WriteLine("--");
        {

            string procedureName = "rvtest";
            Dictionary<string, string> cmdParameters = new Dictionary<string, string>()
	        {
	            {"@a", "12"},
	            {"@b", ""}
	        };
            string connectionString = myConnectionString;

            using (var conn = new SqlConnection(connectionString))
            using (var command = new SqlCommand(procedureName, conn)
            {
                CommandType = CommandType.StoredProcedure
            })
            {

                // fill cmd parameters
                foreach (var pair in cmdParameters)
                {
                    command.Parameters.AddWithValue(pair.Key, pair.Value);
                }

                var returnParameter = command.Parameters.Add("returnParameter", SqlDbType.Int);
                returnParameter.Direction = ParameterDirection.ReturnValue;

                conn.Open();
                command.ExecuteNonQuery();
                conn.Close();

                var result = returnParameter.Value;
                Console.WriteLine(Convert.ToInt32(result));
            }

        }

    }

    /// <summary>
    /// TestMSBin1, Test MSBin1 
    /// </summary>
    static void TestMSBin1()
    {
        Console.WriteLine("== Test MSBin1");

    }
}
