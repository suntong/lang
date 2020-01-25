import java.sql.*;
import java.io.*;
import java.util.*;

class CCTest{

  public static void main(String args[]) throws Exception {
    System.out.println("$)AJdHk TNO:");
    BufferedReader in=new BufferedReader(new InputStreamReader(System.in));
    String t1=in.readLine();
    System.out.println("$)AJdHk TNAME:");
    String t2=in.readLine();

    Class.forName("org.postgresql.Driver");

    Properties info=new Properties();
    info.put("user","pguser");
    info.put("password","Soon2Chng");
    info.put("charSet","GB2312");
    Connection dbconn=DriverManager.getConnection("jdbc:postgresql:test",info);

    PreparedStatement st = dbconn.prepareStatement("insert into test values(?,?)");
    st.setString(1,t1);
    st.setString(2,t2);
    st.executeUpdate();

    String str_sql="SELECT * FROM test";
    System.out.println(str_sql);
    ResultSet rs = st.executeQuery(str_sql);
    while(rs.next()) {
      System.out.println(rs.getString(1)+"t"+rs.getString(2)+"n");
      }

    rs.close();
    st.close();
    dbconn.close();
    }
  }

