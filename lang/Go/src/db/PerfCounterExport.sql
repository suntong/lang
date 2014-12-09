-- To obtain Performance Counters to view within Perfmon
-- From http://blogs.msdn.com/b/geoffgr/archive/2013/09/09/how-to-export-perfmon-counter-values-from-the-visual-studio-load-test-results-database.aspx

--START CODE--
CREATE PROCEDURE TSL_prc_PerfCounterCollectionInCsvFormat
       @RunId nvarchar(10),
       @InstanceName nvarchar(1024)
AS
       DECLARE @CounterName nvarchar(max), @CounterNameColumns nvarchar(max)
       --Get List of columns to use in query. Shove them into a single long XML string
       SELECT @CounterNameColumns = (
                                  SELECT ', [' + REPLACE(InstanceName, ']', ']]') + ']' FROM MTSL_View_PerfmonInstanceNamesAndIds
                                  WHERE LoadTestRunId = @RunId
                                  AND InstanceName LIKE @InstanceName
                     FOR XML PATH(''))
       --Make a copy of the list WITHOUT the comma at the very beginning of the string
       SELECT @CounterName = RIGHT(@CounterNameColumns, LEN(@CounterNameColumns) - 1)
       -- Use the previous strings to build the query string that can be pivoted
       DECLARE @SQL nvarchar(max)
        SELECT @SQL = N'
              select
              IntervalStartTime AS [(PDH-CSV 4.0) (Eastern Daylight Time)(240)]' +
          --IntervalStartTime' +
              @CounterNameColumns + '
        from (
                     select
                           interval.IntervalStartTime,
                           MTSL_View_PerfmonInstanceNamesAndIds.InstanceName,
                           countersample.ComputedValue
                     FROM
                           MTSL_View_PerfmonInstanceNamesAndIds
                           INNER JOIN LoadTestPerformanceCounterSample AS countersample
                           ON countersample.InstanceId = MTSL_View_PerfmonInstanceNamesAndIds.InstanceId
                           AND countersample.LoadTestRunId = MTSL_View_PerfmonInstanceNamesAndIds.LoadTestRunId
                           INNER JOIN LoadTestRunInterval AS interval
                           ON interval.LoadTestRunId = countersample.LoadTestRunId
                           AND interval.TestRunIntervalId = countersample.TestRunIntervalId
                     WHERE
                           MTSL_View_PerfmonInstanceNamesAndIds.LoadTestRunId = ' + @RunId + '
                           AND
                           MTSL_View_PerfmonInstanceNamesAndIds.InstanceName LIKE '''+@InstanceName+'''
        ) Data
        PIVOT (
          SUM(ComputedValue)
          FOR InstanceName
          IN (
               ' + @CounterName + '
          )
        ) PivotTable
        ORDER BY IntervalStartTime ASC
        '
       -- print @SQL
       -- Execute the generated query
       exec sp_executesql @SQL
GO
--START CODE--

GRANT EXECUTE ON TSL_prc_PerfCounterCollectionInCsvFormat TO PUBLIC
GO

/*===============================================================================
MTSL_View_PerfmonInstanceNamesAndIds
===============================================================================*/

CREATE VIEW MTSL_View_PerfmonInstanceNamesAndIds AS
       SELECT
              instance.LoadTestRunId
              ,instance.InstanceId
              ,(
                     '\\' + category.MachineName
                     + '\' + category.CategoryName
                     +      case instance.InstanceName when 'systemdiagnosticsperfcounterlibsingleinstance'
                                   then ''
                                   else '(' + instance.InstanceName  + ')'
                                   end
                     + '\' + counter.CounterName
              ) AS InstanceName
       FROM LoadTestPerformanceCounterCategory AS category
       INNER JOIN LoadTestPerformanceCounter AS counter
              ON category.LoadTestRunId = counter.LoadTestRunId
              AND category.CounterCategoryId = counter.CounterCategoryId
       INNER JOIN LoadTestPerformanceCounterInstance AS instance
              ON counter.CounterId = instance.CounterId
              AND counter.LoadTestRunId = instance.LoadTestRunId
GO

/*===============================================================================
LoadTestRuns
Describe: Returns LoadTestRuns stored in the database
Example: SELECT * FROM LoadTestRuns ORDER BY LoadTestRunId DESC
===============================================================================*/

IF OBJECT_ID ('LoadTestRuns', 'V') IS NOT NULL
DROP VIEW LoadTestRuns
GO

CREATE VIEW LoadTestRuns
AS
SELECT  LoadTestRunId,
        LoadTestName
        ,StartTime
        ,EndTime
        ,RunDuration
        ,Outcome
  FROM  LoadTestRun LTR
GROUP  BY LoadTestRunId,
        LoadTestName
        ,StartTime
        ,EndTime
        ,RunDuration
        ,Outcome
--ORDER  BY StartTime DESC
