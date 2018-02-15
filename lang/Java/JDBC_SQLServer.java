// https://docs.microsoft.com/en-us/sql/connect/jdbc/connection-url-sample

import java.sql.*;  

public class JDBC_SQLServer {  

   public static void main(String[] args) {  

      // Create a variable for the connection string.  
      String connectionUrl = "jdbc:sqlserver://localhost;integratedSecurity=true;databaseName=master";  

      // Declare the JDBC objects.  
      Connection con = null;  
      Statement stmt = null;  
      ResultSet rs = null;  

      try {  
         // Establish the connection.  
         Class.forName("com.microsoft.sqlserver.jdbc.SQLServerDriver");  
         con = DriverManager.getConnection(connectionUrl);  

	 DatabaseMetaData dm = (DatabaseMetaData) con.getMetaData();
	 System.out.println("Driver name: " + dm.getDriverName());
	 System.out.println("Driver version: " + dm.getDriverVersion());
	 System.out.println("Product name: " + dm.getDatabaseProductName());
	 System.out.println("Product version: " + dm.getDatabaseProductVersion());
	 System.out.println();

         // Create and execute an SQL statement that returns some data.  
         String SQL = "select top 5 database_id, name from sys.databases WHERE database_id >= 5";  
         stmt = con.createStatement();  
         rs = stmt.executeQuery(SQL);  

         // Iterate through the data in the result set and display it.  
         while (rs.next()) {  
            System.out.println(rs.getInt(1) + " " + rs.getString(2));  
         }  
      }  

      // Handle any errors that may have occurred.  
      catch (Exception e) {  
         e.printStackTrace();  
      }  
      finally {  
         if (rs != null) try { rs.close(); } catch(Exception e) {}  
         if (stmt != null) try { stmt.close(); } catch(Exception e) {}  
         if (con != null) try { con.close(); } catch(Exception e) {}  
      }  
   }  
}  
