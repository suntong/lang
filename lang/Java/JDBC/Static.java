/*@lineinfo:filename=Static*//*@lineinfo:user-code*//*@lineinfo:1^1*/// -*- Java -*-
//
//   Source File Name: [1]Static.sqlj  1.2
//
//   Licensed Materials -- Property of IBM
//
//   (c) Copyright International Business Machines Corporation, 1998.
//       All Rights Reserved.
//
//   US Government Users Restricted Rights -
//   Use, duplication or disclosure restricted by
//   GSA ADP Schedule Contract with IBM Corp.

//  PURPOSE:
//   This sample program shows how to write a basic SQLJ application.

//  For more information about this sample, refer to the [2]README file.

//  For more information on [3]Programming in Java, refer to the
//  "[4]Programming in Java" section of the [5]Application Development Guide.


//  For more information on building and running Java programs for DB2,
//  refer to the "[6]Building Java Applets and Applications" section of the
//  [7]Application Building Guide.

//  For more information on the SQL language, refer to the [8]SQL Reference.

import java.sql.*;
import sqlj.runtime.*;
import sqlj.runtime.ref.*;

class Static
{   static
    {   try
        {   Class.forName ("COM.ibm.db2.jdbc.app.DB2Driver").newInstance ();
        }
        catch (Exception e)
        {   System.out.println ("\n  Error loading DB2 Driver...\n");
            System.out.println (e);
            System.exit(1);
        }
    }

    public static void main(String argv[])
    {   try
        {   System.out.println ("  Java Static Sample");

            String url = "jdbc:db2:sample";       //  URL is jdbc:db2:dbname
            Connection con = null;

            //  Set the connection                 /* :rk.3:erk. */
            if (argv.length == 0)
            {   //  connect with default id/password
                con = DriverManager.getConnection(url);
            }
            else if (argv.length == 2)
            {   String userid = argv[0];
                String passwd = argv[1];

                //  connect with user-provided username and password
                con = DriverManager.getConnection(url, userid, passwd);
            }
            else
            {throw new Exception("\nUsage: java Static [username password]\n");
            }

            //  Set the default context
            DefaultContext ctx = new DefaultContext(con);
            DefaultContext.setDefaultContext(ctx);


            String firstname = null;

            /*@lineinfo:generated-code*//*@lineinfo:75^13*/

//  ************************************************************
//  #sql { SELECT FIRSTNME 
//                     FROM employee
//                     WHERE LASTNAME = 'JOHNSON'  };
//  ************************************************************

{
  sqlj.runtime.profile.RTResultSet __sJT_rtRs;
  sqlj.runtime.ConnectionContext __sJT_connCtx = sqlj.runtime.ref.DefaultContext.getDefaultContext();
  if (__sJT_connCtx == null) sqlj.runtime.error.RuntimeRefErrors.raise_NULL_CONN_CTX();
  sqlj.runtime.ExecutionContext __sJT_execCtx = __sJT_connCtx.getExecutionContext();
  if (__sJT_execCtx == null) sqlj.runtime.error.RuntimeRefErrors.raise_NULL_EXEC_CTX();
  synchronized (__sJT_execCtx) {
    sqlj.runtime.profile.RTStatement __sJT_stmt = __sJT_execCtx.registerStatement(__sJT_connCtx, Static_SJProfileKeys.getKey(0), 0);
    try 
    {
      sqlj.runtime.profile.RTResultSet __sJT_result = __sJT_execCtx.executeQuery();
      __sJT_rtRs = __sJT_result;
    }
    finally 
    {
      __sJT_execCtx.releaseStatement();
    }
  }
  try 
  {
    sqlj.runtime.ref.ResultSetIterImpl.checkColumns(__sJT_rtRs, 1);
    if (!__sJT_rtRs.next())
    {
      sqlj.runtime.error.RuntimeRefErrors.raise_NO_ROW_SELECT_INTO();
    }
    firstname = __sJT_rtRs.getString(1);
    if (__sJT_rtRs.next())
    {
      sqlj.runtime.error.RuntimeRefErrors.raise_MULTI_ROW_SELECT_INTO();
    }
  }
  finally 
  {
    __sJT_rtRs.close();
  }
}


//  ************************************************************

/*@lineinfo:user-code*//*@lineinfo:77^47*/    /* :rk.4:erk. */


            System.out.println ("First name = " + firstname);
        }
        catch( Exception e )
        {   System.out.println (e);
        }
    }
}/*@lineinfo:generated-code*/class Static_SJProfileKeys 
{
  private static Static_SJProfileKeys inst = null;
  public static java.lang.Object getKey(int keyNum) 
    throws java.sql.SQLException 
  {
    if (inst == null)
    {
      inst = new Static_SJProfileKeys();
    }
    return inst.keys[keyNum];
  }
  private final sqlj.runtime.profile.Loader loader = sqlj.runtime.RuntimeContext.getRuntime().getLoaderForClass(getClass());
  private java.lang.Object[] keys;
  private Static_SJProfileKeys() 
    throws java.sql.SQLException 
  {
    keys = new java.lang.Object[1];
    keys[0] = sqlj.runtime.ref.DefaultContext.getProfileKey(loader, "Static_SJProfile0");
  }
}
