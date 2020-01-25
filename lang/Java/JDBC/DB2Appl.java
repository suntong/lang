//  Source File Name: DB2Appl.java  %I%
//
//  Licensed Materials -- Property of IBM
//
//  (c) Copyright International Business Machines Corporation, 1996, 1997.
//      All Rights Reserved.
//
//  US Government Users Restricted Rights -
//  Use, duplication or disclosure restricted by
//  GSA ADP Schedule Contract with IBM Corp.

//  This sample program shows how to write a Java application using
//  the JDBC application driver to access a DB2 database.

// For more information about this sample, refer to the README file.

// For more information on programming in Java, refer to the 
// "Programming in Java" section of the Application Development Guide.

// For more information on building and running Java programs for DB2,
// refer to the "Building Java Applets and Applications" section of the 
// Application Building Guide.

// For more information on the SQL language, refer to the SQL Reference.

import java.sql.*;

class DB2Appl {

    static {
	try {
	    // register the driver with DriverManager
	    Class.forName("COM.ibm.db2.jdbc.app.DB2Driver");
	} catch (Exception e) {
	    System.out.println ("\n  Error loading DB2 Driver...\n");
	    e.printStackTrace();
            System.exit(1);
	}
    }

    public static void main(String argv[]) {
      Connection con = null;

      // URL is jdbc:db2:dbname
      String url = "jdbc:db2:sample";	// URL is jdbc:db2:dbname

      String newname = "SHILI";

      try {
         if (argv.length == 0) {
            // connect with default id/password
            con = DriverManager.getConnection(url);
            }
         else if (argv.length >= 2) {
            String userid = argv[0];
            String passwd = argv[1];
	    if (argv.length == 3) {
		newname = argv[2];
	    }
	    
            // connect with user-provided username and password
            con = DriverManager.getConnection(url, userid, passwd);
	 }
         else {
            System.out.println("\nUsage: java DB2Appl [username password newname]\n");
            System.exit(0);
         }

         // retrieve data from the database
         System.out.println("Retrieve some data from the database...");
         Statement stmt = con.createStatement();
         ResultSet rs = stmt.executeQuery("SELECT * from employee");

         System.out.println("Received results:");

         // display the result set
         // rs.next() returns false when there are no more rows
         while (rs.next()) {
            String a = rs.getString(1);
            String str = rs.getString(2);

            System.out.print(" empno= " + a);
            System.out.print(" firstname= " + str);
            System.out.print("\n");
         }

         rs.close();
         stmt.close();

         // update the database
         System.out.println("\n\nUpdate the database... ");
         stmt = con.createStatement();
         int rowsUpdated = 
	     stmt.executeUpdate("UPDATE employee set firstnme ='"
				+ newname + "' where empno = '000010'");

         System.out.print("Changed "+rowsUpdated);

         if (1 == rowsUpdated)
            System.out.println(" row.");
         else
            System.out.println(" rows.");

         stmt.close();
         con.close();
      } catch( Exception e ) {
         e.printStackTrace();
      }
   }
}
