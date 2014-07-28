# ==============================================================================================
# Microsoft PowerShell Source File -- Created with SAPIEN Technologies PrimalScript 2009
# NAME: GetTableColumnList.ps1
# AUTHOR: Max.Trinidad , 
# DATE  : 4/01/2011
# 
# COMMENT: This script file is in reference to the following Blog post:
# http://www.maxtblog.com/index.php/2011/04/sql-powershell-list-your-db-table-columns-information/
#
# Initial Credit to - Pinal Dave:
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
$SavedResults1 | ft -auto

## - First create a list of all the tables in the stored object '$SavedResults1'
$TableName = $SavedResults1 | select -unique SchemaTableName | Sort SchemaTableName

## - Verify all tables selected and in order
$TableName

## - Display the each table columns information separately
foreach($t in $TableName){
	$SavedResults1 | where {$_.SchemaTableName -eq $t.SchemaTableName} | `
	select SchemaTableName,ColumnName,SchemaName,TypeName,max_length | FT -Auto
}
## - End of Script - ##