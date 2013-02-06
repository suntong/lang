# install SQL Server 2008 tools to use Invoke-Sqlcmd
add-pssnapin sqlserverprovidersnapin100
add-pssnapin sqlservercmdletsnapin100

Write-Host
Write-Host '== Plan SQL'
Write-Host

Write-Host '-- 1'
# Powershell bug: switch on/off the following two lines you'll get 
# entirely different output in the *following* results
Invoke-Sqlcmd -Query "SELECT GETDATE() AS TimeOfQuery;"
#Invoke-Sqlcmd -Query "SELECT GETDATE() AS TimeOfQuery, @@servername ServerName, DB_NAME();"
Write-Host '-- 2'
Write-Host Powershell bug: Same command can have different result!
Invoke-Sqlcmd -Query "SELECT GETDATE() AS TimeOfQuery;"

Write-Host
Write-Host '== Variable Testing'
Write-Host

Write-Host '-- 1'
Invoke-Sqlcmd -Query "PRINT N'abc'" -Verbose

$var='abcd'
# nok
#Invoke-Sqlcmd -Query "PRINT N'$(var)'" -Verbose
# ok
Invoke-Sqlcmd -Query "PRINT N'$var'" -Verbose
Invoke-Sqlcmd -Query "PRINT N'`$(var)'" -Verbose -Variable "var=abcde"

# for Invoke-Sqlcmd to use the working directory instead of the 
# working location to resolve relative paths
# http://stackoverflow.com/questions/12604902/how-to-convert-this-sqlcmd-script-to-an-invoke-sqlcmd-script
[Environment]::CurrentDirectory = Get-Location

Write-Host '-- 2'
Invoke-Sqlcmd -Query "SELECT GETDATE() AS TimeOfQuery, @@servername ServerName, DB_NAME();"
Write-Host Powershell bug: Use variable or not impacts the output!
$MyArray = "MYVAR1='String1'", "MYVAR2='String2'"
Invoke-Sqlcmd -Query "SELECT `$(MYVAR1) AS Var1, `$(MYVAR2) AS Var2;" -Variable $MyArray
Invoke-Sqlcmd -Query "SELECT GETDATE() AS TimeOfQuery, @@servername ServerName, DB_NAME();" -Variable "var=abcd"

# http://www.mssqltips.com/sqlservertip/1684/powershell-support-in-sql-server-2008-with-the-invoke-sqlcmd-cmdlet/
Invoke-Sqlcmd -Query "SELECT `$(MYVAR1) AS Var1, `$(MYVAR2) AS Var2;" -Variable $MyArray | export-csv -path result.csv
get-content result.csv

Write-Host
Write-Host '== Sql Script Testing'
Write-Host

$sqlCmd="PRINT N'`$(var1), `$(ComputerName)'; SELECT '`$(var1)', '`$(ComputerName)';"
Write-Host $sqlCmd
echo $sqlCmd | Out-File -filePath "ExampleQuery.sql"

Write-Host '-- 1'
Invoke-Sqlcmd -InputFile 'ExampleQuery.sql' -Variable "var1=abcd", "ComputerName = $Env:COMPUTERNAME" | export-csv -path result.csv
get-content result.csv

Write-Host '-- 2'
Invoke-Sqlcmd -InputFile 'ExampleQuery.sql' -Variable "var1=abcd", "ComputerName = $Env:COMPUTERNAME" -Verbose | export-csv -path result.csv
get-content result.csv
