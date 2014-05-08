##----------------------------------------------------------------------------
## Porgram: SqlBackupRestore
## Purpose: SQL server Backup & Restore using SQLPSX
## Authors: Antonio Sun (c) 2013, All rights reserved
##---------------------------------------------------------------------------

import-module sqlserver -force

#region RunTimeSupport

	<#
		.SYNOPSIS
			Temp name generator.

		.DESCRIPTION
			Generate a temp name based on the given name skeleton and current date & time.

		.PARAMETER Name
			The name skeleton used for generating the temp name.

		.PARAMETER TimeFormat
			The time format string used for generating the temp name (default: 'MMdd_HHmm').
		
		.PARAMETER AddTemp
			Add system temp path to the front of the return.
		

		.EXAMPLE
			PS C:\> Get-TempName 'abc.txt'
            abc_0211_1559.txt

		.EXAMPLE
			PS C:\> Get-TempName -AddTemp 'abc.txt'
            C:\Users\...\Temp\abc_0211_1600.txt

		.EXAMPLE
			PS C:\> $fn = Get-TempName -TimeFormat "yyyy-MM-dd_hhmmss" 'test.log'
            PS C:\> $fn 
            test_2013-02-12_090523.log

		.INPUTS
			System.String

		.OUTPUTS
			System.String.
			
	#>

function Get-TempName {
    param(
        [Parameter(ValueFromPipeline=$true,Mandatory=$true)] [ValidateNotNullOrEmpty()]
		[string] $Name,

		[Parameter()]
		[String] $TimeFormat='MMdd_HHmm',

		[Parameter()]
		[Switch] $AddTemp
    )

    $dateStr = get-date -format $TimeFormat
    $fileName = $Name -replace '(.*)\.(.*)', "`$1_$dateStr.`$2"

    if ($AddTemp) { $fileName = "$env:temp\$fileName" }
    return $fileName
}

#endregion

##--------------------------------------------------------
## SYNOPSIS: Database Level Backup & Restore
## DESCRIPTION: Backup & Restore a single DB in MS SQL Server
##--------------------------------------------------------

#region DatabaseLevel

<#
	.SYNOPSIS
		SQL Server Database Backup.

	.DESCRIPTION
		Backup the given MS SQL Server database (as given file).

	.PARAMETER sqlserver 
		[Mandatory] The name of the MS SQL Server to backup from.

	.PARAMETER dbname
		[Mandatory] Database within the given server to backup.
		
	.PARAMETER bakname
		backup name. If empty, default to $dbname.bak.

	.PARAMETER Directory
		Directory name underneath the MS SQL Server default backup directory used for backup.
        CAUTION: The directory must pre-exist. default: empty, i.e., not sub directory used.
		
	.PARAMETER Location
		The Location where the backup is stored. Will overide the "-Directory" setting. 
        No default setting. Must be a pre-existing directory on the DB server.
		
	.PARAMETER Tag
		Tag the DB backup with timestamp.

	.EXAMPLE
		Do-SqlBackup MySvr001 MyDB
		Do-SqlBackup MySvr001 MyDB -Tag

	.EXAMPLE
		Do-SqlBackup MySvr001 MyDB MyDB_20130212.bak
		Do-SqlBackup -sqlserver MySvr001 -dbname MyDB -bakname MyDB_20130212.bak

	.EXAMPLE
		Do-SqlBackup MySvr001 MyDB -Directory '20130212'

	.EXAMPLE
		Do-SqlBackup MySvr001 MyDB -Location 'F:\OtherDir'
			
#>

function Do-SqlBackup {
    param($sqlserver=$(throw 'sqlserver required.'), 
        $dbname=$(throw 'dbname required.'), 
        $bakname='',
        $Directory='',
        $Location='',
		
		[Parameter()]
		[Switch] $Tag
)

    if ($bakname -eq "") {
        $bakname = "$dbname.bak";
    }

    if ($Tag) { 
		$bakname = Get-TempName $bakname;
	}

    $server = Get-SqlServer $sqlserver

    if (-not ($Directory -eq "")) {
        $Directory = $server.BackupDirectory+ "\"+ $Directory+ "\";
    } else {
		$Directory = $server.BackupDirectory+ "\";
	}
    
    if (-not ($Location -eq "")) {
        $Directory = $Location+ "\";
    }

    $bakname = $($Directory+ $bakname)
    Invoke-SqlBackup $sqlserver $dbname $bakname
    return $bakname
}

<#
	.SYNOPSIS
		Restore SQL Server Database Backup.

	.DESCRIPTION
		Restore the given db-backup file to the given MS SQL Server

	.PARAMETER sqlserver 
		[Mandatory] The name of the MS SQL Server to restore to.

	.PARAMETER filepath
		[Mandatory] The db-backup file to restore.
		
	.PARAMETER dbname
		The actual database name if it contains '_'s or extra '.'s.

	.EXAMPLE
        Do-SqlRestore MySvr002 D:\MyDBBackups\MyDB.bak
		Do-SqlRestore MySvr002 D:\MyDBBackups\MyDB_20130212.bak

	.EXAMPLE
		Do-SqlRestore MySvr002 D:\MyDBBackups\My_DB_20130212.bak My_DB
			
#>

function Do-SqlRestore {
    ## Authors: Chad Miller (c) 2011, http://poshcode.org/2531
    param($sqlserver=$(throw 'sqlserver required.'), $filepath=$(throw 'filepath required.'), $dbname='')


    $server = get-sqlserver $sqlserver

    if ($dbname -eq '') {
        $dbname = $filepath -replace('^.*[\\/]','') -replace('[_.].*','')
        }

    $dataPath = Get-SqlDefaultDir -sqlserver $server -dirtype Data
    $logPath = Get-SqlDefaultDir -sqlserver $server -dirtype Log

    $relocateFiles = @{}
    Invoke-SqlRestore -sqlserver $server  -filepath $filepath -fileListOnly | foreach { `
        if ($_.Type -eq 'L')
        { 
			# $physicalName = "$logPath\{0}" -f [system.io.path]::GetFileName("$($_.PhysicalName)") 
			$physicalName = "$logPath\$dbname.ldf"
		}
        else
        { 
			# $physicalName = "$dataPath\{0}" -f [system.io.path]::GetFileName("$($_.PhysicalName)") 
			$physicalName = "$logPath\$dbname.mdf"
		}
        $relocateFiles.Add("$($_.LogicalName)", "$physicalName")
    }

    $server.KillAllProcesses($dbname)

    Invoke-SqlRestore -sqlserver $server -dbname $dbname -filepath $filepath -relocatefiles $relocateFiles -Verbose -force
}

<#
	.SYNOPSIS
		Zap the Log file

	.DESCRIPTION
		Clean up the Log file by empty it first then pregrow to a given size.

	.PARAMETER ServerName 
		[Mandatory] The name of the MS SQL Server to work with.

	.PARAMETER DbName
		[Mandatory] DB within the given server to work on.
		
	.PARAMETER Size
		The size to pregrow the Log file to, default to 4GB. 

	.EXAMPLE
		Do-LogZap 'MySvr001' MyDB

	.EXAMPLE
		Do-LogZap 'MySvr001' MyDB 8GB
		Do-LogZap -ServerName 'MySvr001' -DbName MyDB -Size 8GB
			
#>

function Do-LogZap {
    param(
        [Parameter(Mandatory=$true)] [ValidateNotNullOrEmpty()]
		[string] $ServerName ,

		[Parameter(Mandatory=$true)] [ValidateNotNullOrEmpty()]
		[String] $DbName,
    
		[Parameter()]
		[String] $Size='4GB'

    )

	#Set-SqlData -sqlserver $ServerName -dbname $DbName -qry 'dbcc loginfo'
	Set-SqlData -sqlserver $ServerName -dbname $DbName -qry 'DBCC SHRINKFILE(2, TRUNCATEONLY )'
    $logFile = Get-SqlDatabase $ServerName $DbName | Get-SqlLogFile
	Set-SqlData -sqlserver $ServerName -dbname $DbName -qry `
		"ALTER DATABASE $DbName MODIFY FILE (NAME = `"$($logFile.Name)`", SIZE = $Size)" 

}

#endregion

##--------------------------------------------------------
## SYNOPSIS: Server Level Backup & Restore
## DESCRIPTION: Backup & Restore the entire MS SQL Server
##--------------------------------------------------------

#region ServerLevel

<#
	.SYNOPSIS
		Server Backup.

	.DESCRIPTION
		Backup the given MS SQL Server of the given DBs.

	.PARAMETER ServerName 
		[Mandatory] The name of the MS SQL Server to backup.

	.PARAMETER DBs
		DBs within the given server to backup (regexp). If empty, all dbs are to be backup.
		
	.PARAMETER Directory
		Directory name underneath the MS SQL Server default backup directory used for backup.
        CAUTION: The directory must pre-exist.

	.PARAMETER Location
		The Location where the backup is stored. Will overide the "-Directory" setting. 
        No default setting. Must be a pre-existing directory on the DB server.
		
	.PARAMETER Check
		Check the DB selection of the given DBs.
		
	.EXAMPLE
		Do-SvrBackup MySvr001 -Check
		Do-SvrBackup 'MySvr001' 'this|that|th[eo]se' -Check

	.EXAMPLE
		Do-SvrBackup 'MySvr001' 'this|that|th[eo]se' '20130212'
		Do-SvrBackup -ServerName 'MySvr001' -DBs 'this|that|th[eo]se' -Directory '20130212'

	.EXAMPLE
		Do-SvrBackup 'MySvr001' 'this|that|th[eo]se' -Location 'F:\OtherDir'
			
#>

function Do-SvrBackup {
    param(
        [Parameter(Mandatory=$true)] [ValidateNotNullOrEmpty()]
		[string] $ServerName ,

		[Parameter()]
		[String] $DBs='',
    
		[Parameter()]
		[String] $Directory='',

		[Parameter()]
		[String] $Location='',

		[Parameter()]
		[Switch] $Check
    )

    if ($Check) { 
            Write-Host "DBs to be backed up are:`n------------------------"
    }

    Get-SqlDatabase $ServerName |
         where-object { $_.Name -match $DBs } | 
         Select-Object Name | foreach {
         if ($Check) { 
            Write-Host "$($_.Name)"
         } elseif (-not ($Location -eq "")) {
             Write-Host -NoNewline "Backing up '$($_.Name)' under '$Location'... "
             Do-SqlBackup -Location $Location $ServerName $($_.Name)
             Write-Host "Done."
         } else {
             Write-Host -NoNewline "Backing up '$($_.Name)' under '$Directory'... "
             Do-SqlBackup -Directory $Directory $ServerName $($_.Name)
             Write-Host "Done."
         }
         }
}

<#
	.SYNOPSIS
		Server Restore.

	.DESCRIPTION
		Restore the MS SQL Server backups from the given directory locally.

	.PARAMETER Directory
		Directory containing the SQL Server backups files.

	.PARAMETER DbExt
		DB extension of the SQL Server backups files (default: 'bak').
		
	.PARAMETER Remote
		Do DB restore remotely on the given SQL Server
		
	.EXAMPLE
		Do-SvrBackup D:\MyDBBackups\20130212

	.EXAMPLE
		Do-SvrBackup D:\MyDBBackups\20130212 -Remote MySvr002
			
#>
function Do-SvrRestore {
    param(
        [Parameter(Mandatory=$true)] [ValidateNotNullOrEmpty()]
		[String] $Directory='',

		[Parameter()]
		[String] $DbExt='bak', 

		[Parameter()]
		[String] $Remote=''
    )

	#get current login info
	$CS = Gwmi Win32_ComputerSystem -Comp "."
	$LogonHost=$CS.Name
	$LogonUser=$CS.UserName

    $ServerName = $LogonHost
    $LookupDirectory = $Directory
    if (-not ($Remote -eq "")) {
        $ServerName = $Remote;
        $LookupDirectory="\\$ServerName\"+$Directory.Replace(':','$')
        Write-Verbose $LookupDirectory
    }
    
    get-childitem $LookupDirectory *.$DbExt | foreach {
        Write-Host "$Directory\$_ restoring started ... "
        Do-SqlRestore $ServerName $Directory\$_
        Write-Host "$Directory\$_ restoring finished."
    }
}



function Defrag-SvrIndexes
{
# Adapted by: Antonio Sun,  Date: 12/02/2013
# Initial by: Steve Wright, Date: 20/01/2012
#             

<#
.SYNOPSIS
Defrag the Indexes of the requestd SQL Server databases

.DESCRIPTION
For the selected databases on the SQL Server will loop through all the user tables indexes to see if they 
need to be Reorganize or Rebuild.
Using windows authentication for connecting to remote servers.
Also to be able to use the funcation the SQL Server SMO and powershell SQL Provider

.NOTES
The function will either Reorganize Index or Rebuild Index
If the index AverageFragmentation ranges in between 5% to 30% then it is better to perform Reorganize Index.
If the index AverageFragmentation is greater than 30% then the best strategy will be to use Rebuild Index.
Recommandations where found on many articles on the internet.

WARNING: It takes several magnitudes longer than building offline in SSMS.

.PARAMETER server 
The name of the SQL Server to connect to gather the file usages of the database.

.PARAMETER DBs
DBs within the given server to defrag the indexes (regexp). Empty means all dbs.

.PARAMETER Full
Output full report, including the skipped tables.

.PARAMETER fragmentationOption
The specify the levels of detail of collected fragmentation information
Fast:Calculates statistics based on parent level pages only. This option is available starting with SQL Server 2000.  
Sampled:Calculates statistics based on samples of data. This option is available starting with SQL Server 2005.  
Detailed:(Default) Calculates statistics based on 100% of the data. This option is available starting with SQL Server 2005.  

.EXAMPLE
Connects to the local server database myDB using windows authentication, with verbose progress info
	
Defrag-SvrIndexes '(local)' -databaseName myDB -Verbose

.EXAMPLE
Connects to the remote server database using windows authentication, output to GridView
	
Defrag-SvrIndexes -server SQLServer01 -databaseName myDB | 
  select-Object -property Database, Table, Index, AverageFragmentation, ActionTaken | 
  Out-GridView

.EXAMPLE
Connects to the local server database using windows authentication with fragmentationOption of Fast
	
Defrag-SvrIndexes '(local)' -databaseName myDB -fragmentationOption [Microsoft.SqlServer.Management.Smo.FragmentationOption]::Fast 

.EXAMPLE
Connects to the remote server database using windows authentication with fragmentationOption of Sampled
	
Defrag-SvrIndexes -server SQLServer01 -databaseName myDB -fragmentationOption [Microsoft.SqlServer.Management.Smo.FragmentationOption]::Sampled 
	
.INPUTS
None. You cannot pipe objects to Defrag-SvrIndexes
 
.Outputs
Array of PSObject with the following properties: 

    Database
    Table
    Index
    AverageFragmentation
    ActionTaken

.COMPONENT
Microsoft® Windows PowerShell Extensions for SQL Server® 2008 R2.

.COMPONENT
Microsoft® SQL Server® 2008 R2 Shared Management Objects.

.LINK
Initial script: http://zogamorph.blogspot.ca/2012/02/sql-server-maintenance-via-powershell.html
Components Download: http://www.microsoft.com/download/en/details.aspx?displaylang=en&id=16978 
#>
	[CmdletBinding()]
	param (
        [Parameter(Mandatory=$true)] [ValidateNotNullOrEmpty()]
		[string]$server,

		[Parameter()]
		[string]$DBs='',

		[Parameter()]
		[Switch] $Full,

		[Parameter()]
		[Microsoft.SqlServer.Management.Smo.FragmentationOption]
		$fragmentationOption = [Microsoft.SqlServer.Management.Smo.FragmentationOption]::Detailed
	)
	
	$properties = @{Database = [string] "";
	                Table = [string] "";
					Index = [string] "";
					AverageFragmentation = [float] 0.0;
					ActionTaken = [string] "";
					}

	$srv = New-Object -TypeName Microsoft.SqlServer.Management.SMO.Server -ArgumentList $server

    $srv.Databases | where-object {(!$_.IsSystemObject) -and 
                          $_.IsAccessible -and $_.name -match $DBs } | 
    foreach {

        $databaseName = $_.name

	    $db = $srv.Databases[$databaseName]
	    $results = @()

	    foreach($dbtable in $db.Tables) 
	    {
		    foreach($dbIndex in $dbtable.Indexes) 
		    {
			    $indexResults = $dbIndex.EnumFragmentation($fragmentationOption)
			    $methodTaken = New-Object PSObject -Property $properties
			    $methodTaken.Database = $db.Name
			    $methodTaken.Table = $dbtable
			    $methodTaken.Index = $dbIndex.Name
			    $methodTaken.AverageFragmentation = $indexResults.Rows[0]["AverageFragmentation"]
			  
			     if($methodTaken.AverageFragmentation -ge 30 )
			     {
                    Write-Verbose "$databaseName.$dbtable.$($dbIndex.Name): Rebuilding"
			 	    $methodTaken.ActionTaken = "Rebuild"
			 	    $dbIndex.Rebuild()
			     }
			     elseif($methodTaken.AverageFragmentation -ge 5) 
			     {
                    Write-Verbose "$databaseName.$dbtable.$($dbIndex.Name): Reorganizing"
			 	    $methodTaken.ActionTaken = "Reorganize"
			 	    $dbIndex.Reorganize()
			     }
			     else
			     {
                    if ($Full) { 
                        Write-Verbose "$databaseName.$dbtable.$($dbIndex.Name): Skipped"
			 	        $methodTaken.ActionTaken = "None"
                    }
			     }
			 
			     $results += $methodTaken
		    }	
	    }
    }
	return $results
}


#endregion
