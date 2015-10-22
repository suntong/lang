----------------------------------------------------------------------------
-- Porgram: GetView.sql
-- Purpose: Get View Definition
-- Authors: Tong Sun (c) 2014, All rights reserved
----------------------------------------------------------------------------

/*
Invoke:

 set dbserver=MyDb05
 sqlcmd -E -S %dbserver% -d MyDatabase -v viewName=vwMyView -h-1 -y 0 -i GetView.sql
 sqlcmd -U %user% -P %pwd% -S %dbserver% -v ...

*/

SET NOCOUNT ON

SELECT Definition
FROM sys.objects o
JOIN sys.sql_modules m ON m.object_id = o.object_id
WHERE o.object_id = object_id('$(viewName)')
    AND o.type = 'V'
