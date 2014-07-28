# ==============================================================================================
# Microsoft PowerShell Source File -- Created with SAPIEN Technologies PrimalScript 2009
# NAME: GetTableColumnList_Enhanced.ps1
# AUTHOR: Max.Trinidad , 
# DATE  : 4/01/2011
# 
# COMMENT: This is the enhanced version not using the ForEach command. credit to David Moravec for 
#          fixing the "FT -Auto" command to include the "-GroupBy" parameter.
# T-SQL script Initial Credit to - Pinal Dave:
# - http://blog.sqlauthority.com/2007/08/09/sql-server-2005-list-all-the-column-with-specific-data-types/
# ==============================================================================================

## - Stored T-SQL modified script into a string variable
$sqlQuery = @"
SELECT
s.name+'.'+OBJECT_NAME(c.OBJECT_ID) as SchemaTableName
,c.name AS ColumnName
,SCHEMA_NAME(t.schema_id) AS SchemaName
,t.name AS TypeName
,c.max_length
FROM sys.columns AS c
JOIN sys.types AS t ON c.user_type_id=t.user_type_id
JOIN sys.tables as t2 on t2.object_id = c.object_id
JOIN sys.schemas as s on s.schema_id = t2.schema_id
ORDER BY c.OBJECT_ID;
"@

## - Load the SQLPS module for Denali ( for 2008/200R2 is SQLPS)
Import-Module SQLPSv2
$SavedResults1 = Invoke-SQLCmd -ServerInstance "MAX-PCWIN1" -Database "Adventureworks" -Query $sqlQuery
$SavedResults1 | ft -auto -GroupBy SchemaTableName

## - End of Script - ##