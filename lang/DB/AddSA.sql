----------------------------------------------------------------------------
-- Porgram: AddSA
-- Purpose: Add the give domain user as sa for Sql Server
-- Authors: Tong Sun (c) 2013, All rights reserved
----------------------------------------------------------------------------

/*
Database: MS SQL Server 2008

Invoke:

 set dbserver=MySvr01
 sqlcmd -E -S %dbserver% -v du=DOMAIN\user -i AddSA.sql
 sqlcmd -U %user% -P %pwd% -S %dbserver% -v ...

*/

CREATE LOGIN [$(du)] FROM WINDOWS
GO
EXEC master..sp_addsrvrolemember @loginame = N'$(du)', @rolename = N'sysadmin'
GO
