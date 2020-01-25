/**
 *
 * @author: SUN, Tong
 *
 */

import java.net.URL;
import java.sql.*;
import java.lang.Runtime;

class jdbc_demo {

  // instance variables

  public static String query;

  public static Connection con;
  public static Statement  stmt;
  public static ResultSet  results;

  public static void CreateTable() {

    try {
      stmt = con.createStatement();
      int ResultCode;
      String SQLText =
	"create table tab1( keycol CHAR(3), col1 CHAR(20), col2 integer )";
      ResultCode = stmt.executeUpdate( SQLText );
      System.out.println( "ok: create table tab1");
      }
    catch (java.lang.Exception ex) {

      // Print description of the exception.
      System.out.println( "** Error on create table tab1. ** " );
      ex.printStackTrace ();

      }
    }

  // --------------------------------------------------------------------
  public static void CreateCustomersTable() {
    try  {
      int ResultCode;

      stmt = con.createStatement();
      // create the customer table
      String SQLText =
	"create table tb_customers (  firstname char(20), " +
	" lastname char(20),  " +
	" city      char(20), " +
	" address   char(20), " + 
	" state     char(2),  " +
	" zip       char(6)) ";
      ResultCode = stmt.executeUpdate( SQLText );

      System.out.println( "ok: create table tb_customers");
      }
    catch (java.lang.Exception ex) {
      // Print description of the exception.
      System.out.println( "** Error on create table tb_customers. ** " );
      ex.printStackTrace ();
      }
    }     

  public static void InsertData() {

    try {
      stmt = con.createStatement();
      int ResultCode;
      String SQLText = "insert into tab1 values ('AT','ABC', 10 )";
      ResultCode = stmt.executeUpdate( SQLText );
      System.out.println( "Inserted " + ResultCode + " rows." );

      SQLText = "insert into tab1 values ('AT1','ABC',11 )";
      System.out.println( "Inserted " + ResultCode + " rows." );
      ResultCode = stmt.executeUpdate( SQLText );

      SQLText = "insert into tab1 values ('CC1','数据库',11 )";
      System.out.println( "Inserted " + ResultCode + " rows." );
      ResultCode = stmt.executeUpdate( SQLText );

      SQLText = "insert into tab1 values ('CC2','中文访问',12 )";
      ResultCode = stmt.executeUpdate( SQLText );
      System.out.println( "Inserted " + ResultCode + " rows." );

      // insert into the customer table
      SQLText = "insert into tb_customers values ( " +
	" 'Paul', 'Mark', '2020 Main', 'Tampa', 'FL', '33210' ) ";
      ResultCode = stmt.executeUpdate( SQLText );

      SQLText = "insert into tb_customers values ( " +
	" 'John', 'Lyon', '3030 South St.', 'Strawbern', 'PA', '44210' ) ";
      ResultCode = stmt.executeUpdate( SQLText );

      SQLText = "insert into tb_customers values ( " +
	" 'George', 'Hammond', '4040  North St.', 'Gutawep', 'CA', '66210' ) ";
      ResultCode = stmt.executeUpdate( SQLText );

      SQLText = "insert into tb_customers values ( " +
	" 'Richard', 'Stars', '6060 West Ave.', 'Drummon', 'NY', '97210' ) ";
      ResultCode = stmt.executeUpdate( SQLText );

      System.out.println( "ok: insert into table");

      }
    catch (java.lang.Exception ex) {
      // Print description of the exception.
      System.out.println( "** Error on data insert. ** " );
      ex.printStackTrace ();
      }
    }

  // --------------------------------------------------------------------
  public static void SelectData0() {

    try {
      stmt = con.createStatement();
      String SQLText = "select * from tab1";
      results = stmt.executeQuery( SQLText );
      DisplayResults( results );
       
      //System.out.println( "ok: select * from tab1");
      }
    catch (java.lang.Exception ex) {
      // Print description of the exception.
      System.out.println( "** Error on data select. ** " );
      ex.printStackTrace ();
      }
    }

  // --------------------------------------------------------------------
  public static void SelectData(String fields, String table) {

    try {
      stmt = con.createStatement();
      String SQLText = "select " + fields + " from " + table;
      results = stmt.executeQuery( SQLText );
      DisplayResults( results );
       
      System.out.println
	("ok: select fields '" + fields + "' from table: "+ table);
      }
    catch (java.lang.Exception ex) {
      // Print description of the exception.
      System.out.println( "** Error on data select. ** " );
      ex.printStackTrace ();
      }
    }

  // --------------------------------------------------------------------
  public static void UpdateData() {

    try {
      stmt = con.createStatement();
      int ResultCode;

      String SQLText = "update tab1 set col1 = 'BLA BLA' where keycol = 'AT1'";

      ResultCode = stmt.executeUpdate( SQLText );
      System.out.println( "Updated " + ResultCode + " rows." );

      // show the updated row
      jdbc_demo.SelectData0();

      // print the number of rows updated 

       
      System.out.println( "ok: update tab1 ");
      }
    catch (java.lang.Exception ex) {
      // Print description of the exception.
      System.out.println( "** Error on data update. ** " );
      ex.printStackTrace ();
      }
    }


  // --------------------------------------------------------------------
  public static void DeleteData() {

    try {
      stmt = con.createStatement();
      int ResultCode;
      String SQLText = "delete from tab1 where keycol = 'AT1'";
      ResultCode = stmt.executeUpdate( SQLText );

      // print the number of rows deleted
      System.out.println( "Deleted " + ResultCode + " rows." );

      // show the deleted row
      jdbc_demo.SelectData0();
       
      System.out.println( "ok: delete from tab1");
      }
    catch (java.lang.Exception ex) {
      // Print description of the exception.
      System.out.println( "** Error on data delete. ** " );
      ex.printStackTrace ();
      }
    }

  public static void DropTable() {

    try {
      stmt = con.createStatement();
      int ResultCode;
      String SQLText =
	"drop table ";
      ResultCode = stmt.executeUpdate( SQLText + "tab1" );
      ResultCode = stmt.executeUpdate( SQLText + "tb_customers" );
      System.out.println( "ok: drop table tables");
      }
    catch (java.lang.Exception ex) {
      // Print description of the exception.
      System.out.println( "** Error on dropping tables. ** " );
      ex.printStackTrace ();
      }
    }



  // --------------------------------------------------------------------
  public static void DisplayResults( ResultSet results ) throws SQLException {
    int i;

    // Get the ResultSetMetaData and use this for
    // the column headings

    ResultSetMetaData rsmd = results.getMetaData ();

    // Get the number of columns in the result set

    int numCols = rsmd.getColumnCount ();

    // Display column headings

    System.out.println("=================================================");
    for (i=1; i<=numCols; i++) {
      if (i > 1) System.out.print(" | ");
      System.out.print(rsmd.getColumnLabel(i));
      }
    System.out.println("");
    System.out.println("-------------------------------------------------");
		
    // Display data, fetching until end of the result set

    boolean more = results.next ();
    while (more) {

      // Loop through each column, getting the
      // column data and displaying

      for (i=1; i<=numCols; i++) {
	if (i > 1) System.out.print(" | ");
	System.out.print(results.getString(i));
	}
      System.out.println("");

      // Fetch the next result set row

      more = results.next ();
      }
    System.out.println("=================================================");
    }
  }


  /* --------------------------------------------------------------------- */

  class JDBCExample {

    static String dbdriver;
    static String url;

    static String userid = "";
    static String passwd = "";

    public static void fulltest () {
      // Create the table
      jdbc_demo.CreateTable();
      jdbc_demo.CreateCustomersTable();

      // Insert Data
      jdbc_demo.InsertData();

      // Select Data
      jdbc_demo.SelectData0();
      jdbc_demo.SelectData("*", "tab1");


	// Update Data
      jdbc_demo.UpdateData();

      // Delete data
      jdbc_demo.DeleteData();

      // Drop tables
      jdbc_demo.DropTable();
      }

    //-------------------------------------------------------------------
    // checkForWarning
    // Checks for and displays warnings.  Returns true if a warning
    // existed
    //-------------------------------------------------------------------

    private static boolean checkForWarning (SQLWarning warn)
      throws SQLException {
      boolean rc = false;

      // If a SQLWarning object was given, display the
      // warning messages.  Note that there could be
      // multiple warnings chained together

      if (warn != null) {
	System.out.println ("\n *** Warning ***\n");
	rc = true;
	while (warn != null) {
	  System.out.println ("SQLState: " +
			      warn.getSQLState ());
	  System.out.println ("Message:  " +
			      warn.getMessage ());
	  System.out.println ("Vendor:   " +
			      warn.getErrorCode ());
	  System.out.println ("");
	  warn = warn.getNextWarning ();
	  }
	}
      return rc;
      }

    public static void main (String argv[]) {

      dbdriver = "COM.ibm.db2.jdbc.app.DB2Driver";
      url = "jdbc:db2:sample";	// URL is jdbc:db:dbname

      dbdriver = "org.postgresql.Driver";
      url = "jdbc:postgresql:template1";	// URL is jdbc:db:dbname
      url = "jdbc:postgresql://localhost/template1";
      url = "jdbc:postgresql://localhost:5432/template1";

      if (argv.length == 0) {
	System.err.println("\nUsage: java DB2Appl username password [url [driver]]\n");
	System.exit(0);
	}

      if (argv.length >= 2) {
	userid = argv[0];
	passwd = argv[1];

	if (argv.length >= 3) { url = argv[2]; }
	if (argv.length >= 4) { dbdriver = argv[3]; }
	}
                
      System.err.println("\nProgram started\n");

      // Load the jdbc driver
      try {
	Class.forName (dbdriver);
	} 
      catch(java.lang.ClassNotFoundException e) {
	System.err.print("Error when loading the jdbc driver\n" + 
			 " ClassNotFoundException: ");
	System.err.println(e.getMessage());
	}
	
      try { 
	// Attempt to connect to a driver.  
	jdbc_demo.con = DriverManager.getConnection (url, userid, passwd);

	// If we get here, we are successfully
	// connected to the URL

	fulltest();

	// Close the statement
	jdbc_demo.stmt.close();

	// Close the connection
	jdbc_demo.con.close();
	}
      catch (SQLException ex) {

	// A SQLException was generated.  Catch it and
	// display the error information.  Note that there
	// could be multiple error objects chained
	// together
	System.out.println ("\n*** SQLException caught ***\n");

	while (ex != null) {
	  System.out.println("SQLState: " +
			     ex.getSQLState ());
	  System.out.println("Message:  " +
			     ex.getMessage ());
	  System.out.println("Vendor:   " +
			     ex.getErrorCode ());
	  ex = ex.getNextException ();
	  System.out.println("");
	  }
	}
      catch (java.lang.Exception ex) {
	// Got some other type of exception.  Dump it.
	ex.printStackTrace();
	}
      }

    }
