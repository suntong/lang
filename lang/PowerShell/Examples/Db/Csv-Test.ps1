##----------------------------------------------------------------------------
## Porgram: Csv-Test
## Purpose: Import-Csv into SQL Server tables using SQLPSX
## Authors: Antonio Sun (c) 2013, All rights reserved
##---------------------------------------------------------------------------

# To execute:
#    .\Csv-Test.ps1

# To catch misspelled variables
Set-PSDebug -strict

Write-Host "`n== Plain Import-Csv"
Import-Csv .\Csv-Test.csv

Write-Host "`n== Import-Csv and filtered with Where-Object`n"
Import-Csv .\Csv-Test.csv | Where-Object {$_.department -eq "Finance"}
Write-Host "--"
Import-Csv .\Csv-Test.csv | Where-Object {$_.department -ne "Finance"}
Write-Host "--"
Import-Csv .\Csv-Test.csv | Where-Object {$_.department -eq "Finance" -and $_.title -eq "Accountant"}
Write-Host "--"
Import-Csv .\Csv-Test.csv | Where-Object {$_.department -eq "Research" -or $_.title -eq "Accountant"}
# The above are from http://technet.microsoft.com/en-us/library/ee176874.aspx

Write-Host "`n== Import-Csv with Delimiter`n"
Import-Csv .\Csv-Test.csv -Delimiter ","
# To Import Tab Delimited File,
#Import-Csv -Delimiter "`t" csvfile

Write-Host "`n== Get the result of Import-Csv into array`n"
$guys = @(Import-Csv .\Csv-Test.csv | Where-Object {$_.department -eq "Research" -or $_.title -eq "Accountant"})
# process guys
foreach($guy in $guys) {
	Write-Host "The guy: $guy, $($guy.Department)"
    Write-Host "    Department: $($guy.Department) Title: $($guy.Title)"
}

Write-Host "`n== Directly Import-Csv into foreach`n"
Import-Csv .\Csv-Test.csv | Where-Object {$_.department -eq "Research" -or $_.title -eq "Accountant"} | foreach {
	Write-Host "    Department: $($_.Department) Title: $($_.Title)"
}


##-----------------------------------------------

Write-Host "`n== Import-Csv and write into SQL Server table`n"

import-module sqlserver -force

# Define table structure
Set-SqlData "(local)" "tempdb" "IF OBJECT_ID ('Tmp', 'U') IS NOT NULL DROP TABLE Tmp"
Set-SqlData "(local)" "tempdb" "CREATE TABLE Tmp (Name varchar(36), Department varchar(12),  Title varchar(36) )"
# Import-Csv and write into SQL Server table
$cmd = ''
Import-Csv .\Csv-Test.csv | foreach {
	$cmd += "INSERT INTO Tmp VALUES ('$($_.Name)', '$($_.Department)', '$($_.Title)');`n";
}
$cmd
Set-SqlData "(local)" "tempdb" $cmd


##-----------------------------------------------

# http://www.computerperformance.co.uk/powershell/powershell_import_csv.htm

# PowerShell Export-CSV Example 
$FilePath = ".\Service.csv"
Get-Service | Export-CSV $FilePath

# PowerShell Import-CSV to View Properties
Import-CSV $FilePath  | Get-Member

# PowerShell CSV Cmdlet Research
Get-Command -Noun CSV

