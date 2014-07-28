#################################################################
## Compare SQL Server Table schemas, Main
## http://powershell.org/wp/2013/04/28/comparing-sql-server-table-schemas-with-powershell/
## Enrique Puig Nouselles
## Epuig1984@gmail.com
#################################################################

#main program
##Define Variables
$srv1="(local)"
$srv2="(local)"

$bd1="TableCompare"
$bd2="TableCompare2"

$sch1="dbo"
$sch2="dbo"

$TableName1="TestTable"
$TableName2="TestTable"

#function call
Compare-SQLServerTables $srv1 $bd1 $sch1 $TableName1 $srv2 $bd2 $sch2 $TableName2 