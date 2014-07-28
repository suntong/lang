##----------------------------------------------------------------------------
## Porgram: Telerik2Table
## Purpose: Import Telerik log Csv file into SQL Server table using SQLPSX
## Authors: Antonio Sun (c) 2013, All rights reserved
##---------------------------------------------------------------------------

# To execute:
#    . .\Telerik2Table.ps1
#
#    # Import-Csv and pipe into SQL Server table
#    Import-Csv -Delimiter "`t" .\Telerik.csv | Telerik2Table ServerName TableName

# Input csv file struture:
#   Time`tStep
# Output SQL Server table struture
#   CREATE TABLE log (
#     id int IDENTITY(1,1) NOT NULL
#     , testID varchar(12) NOT NULL
#     , agent varchar(12) NOT NULL
#     , comment varchar(128) NOT NULL
#     , tstamp datetime NOT NULL
#   ) 


import-module sqlserver -force

# To catch misspelled variables
Set-PSDebug -strict

# Define a pipeline function
function Telerik2Table {
    param($sqlserver=$(throw 'sqlserver name required.'),$dbname=$(throw 'dbname required.'))

    begin { 
        # Clear the table
        Set-SqlData $sqlserver $dbname "DELETE FROM log"
        $cmd = ''; 
    }
    process {
        $comment = $_.Step -Replace "'", "''"; 
        $cmd += "INSERT INTO log (testID, agent, comment, tstamp) VALUES ('Telerik','1-1', '$comment', '$($_.time)')`n"
    }
    end { $cmd; Set-SqlData $sqlserver $dbname $cmd }
}
