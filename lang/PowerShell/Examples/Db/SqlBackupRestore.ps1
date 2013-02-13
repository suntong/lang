##----------------------------------------------------------------------------
## Porgram: SqlBackupRestore
## Purpose: SQL server Backup & Restore using SQLPSX
## Authors: Antonio Sun (c) 2013, All rights reserved
##---------------------------------------------------------------------------

import-module sqlserver -force

##--------------------------------------------------------
## SYNOPSIS: Database Level Backup & Restore
## DESCRIPTION: Backup & Restore a single DB in MS SQL Server
##--------------------------------------------------------

#region DatabaseLevel

function Do-SqlBackup {
    param($sqlserver=$(throw 'sqlserver required.'), 
        $dbname=$(throw 'dbname required.'), 
        $bakname='',
        $Directory='')

    if ($bakname -eq "") {
        $bakname = "$dbname.bak";
    }

    if (-not ($Directory -eq "")) {
        $Directory += "\";
    }
    
    $server = Get-SqlServer $sqlserver
    $bakname = $($server.BackupDirectory+ "\"+ $Directory+ $bakname)
    Invoke-SqlBackup $sqlserver $dbname $bakname
    return $bakname
}


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
        { $physicalName = "$logPath\{0}" -f [system.io.path]::GetFileName("$($_.PhysicalName)") }
        else
        { $physicalName = "$dataPath\{0}" -f [system.io.path]::GetFileName("$($_.PhysicalName)") }
        $relocateFiles.Add("$($_.LogicalName)", "$physicalName")
    }

    $server.KillAllProcesses($dbname)

    Invoke-SqlRestore -sqlserver $server -dbname $dbname -filepath $filepath -relocatefiles $relocateFiles -Verbose -force
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
			The name of the MS SQL Server to backup.

		.PARAMETER DBs
			DBs within the given server to backup (regexp). If empty, all dbs are to be backup.
		
		.PARAMETER Directory
			Directory name underneath the MS SQL Server default backup directory used for backup.
            CAUTION: The directory must pre-exist.

		.PARAMETER Check
			Check the DB selection of the given DBs.
		
		.EXAMPLE
			Do-SvrBackup MySvr001 -Check
			Do-SvrBackup 'MySvr001' 'this|that|th[eo]se' -Check

		.EXAMPLE
			Do-SvrBackup 'MySvr001' 'this|that|th[eo]se' '20130212'
			
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
		
		
	.EXAMPLE
		Do-SvrBackup D:\MyDBBackups\20130212
			
#>
function Do-SvrRestore {
    param(
        [Parameter(Mandatory=$true)] [ValidateNotNullOrEmpty()]
		[String] $Directory='',

		[Parameter()]
		[String] $DbExt='bak'
    )

	#get current login info
	$CS = Gwmi Win32_ComputerSystem -Comp "."
	$LogonHost=$CS.Name
	$LogonUser=$CS.UserName

    get-childitem $Directory *.$DbExt | foreach {
        write-host "$Directory\$_ restoring started ... "
        Do-SqlRestore $LogonHost $Directory\$_
        Write-Host "$Directory\$_ restoring finished."
    }
}

#endregion
