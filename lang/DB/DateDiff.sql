----------------------------------------------------------------------------
-- Porgram: DateDiff
-- Purpose: Return date difference in Days,hours minutes and seconds
-- Authors: Tong Sun (c) 2013, All rights reserved
----------------------------------------------------------------------------

/*
Invoke:

  SELECT dbo.DateDiff_ToStr('2013-11-03 8:15:03.497', '2013-11-07 10:22:57.093') 
  SELECT MyDb.dbo.DateDiff_ToStr('2013-11-03 8:15:03.497', '2013-11-07 10:22:57.093') 
  -- 4d 2:7:53.596

Ref:

DECLARE @Start DATETIME, @END DATETIME
SET @Start = '2013-11-07 10:15:03.497'
set @End = '2013-11-07 10:22:57.093'
SELECT CONVERT(char(8), @end-@start, 108) ExecTime  -- 00:07:53
SELECT CONVERT(char(14), @end-@start, 114) ExecTime -- 00:07:53:597  

http://msdn.microsoft.com/en-us/library/ms187928.aspx

Adapted from:

http://social.msdn.microsoft.com/Forums/sqlserver/en-US/5e5a4474-31b8-4316-8a34-1e4a5572fb49/
by thilla (MSFT), Jin Chen - MSFT

*/

IF OBJECT_ID (N'dbo.DateDiff_ToStr', N'FN') IS NOT NULL
    DROP FUNCTION DateDiff_ToStr;
GO

CREATE FUNCTION DateDiff_ToStr (@StartTime DATETIME, @EndTime DATETIME)
RETURNS VARCHAR(15)
AS
BEGIN
    DECLARE @I INT
    SET @I = DATEDIFF(ms,@StartTime,@EndTime)

    DECLARE @R VARCHAR(15)
    SELECT @R = 
        convert(varchar(10), (@I/86400000)) + 'd ' + 
        convert(varchar(10), ((@I%86400000)/3600000)) + ':'+
        convert(varchar(10), (((@I%86400000)%3600000)/60000)) + ':'+
        convert(varchar(10), ((((@I%86400000)%3600000)%60000)/1000)) + '.' +
        convert(varchar(10), (((@I%86400000)%3600000)%1000))
    RETURN @R
END
