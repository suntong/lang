# LibrarySqlBackup, Standalone SQL Server backup and restore functions adapted from sqlpsx
# http://poshcode.org/1188

# Usage:
# . .\LibrarySqlBackup.ps1
# $serverName = 'Z002\SqlExpress'; $dbName = 'pubs'
# $server = new-object ("Microsoft.SqlServer.Management.Smo.Server") $serverName
# invoke-sqlbackup $serverName $dbName $($server.BackupDirectory + "\$dbName.bak")

# ---------------------------------------------------------------------------
### <Script>
### <Author>
### Chad Miller 
### </Author>
### <Description>
### Excerpt from SQL Server Powershell Extensions (sqlpsx)
### http://sqlpsx.codeplex.com
### Defines backup and restore functions
### </Description>
### <Usage>
### . ./LibrarySqlBackup.ps1
### $server = new-object ("Microsoft.SqlServer.Management.Smo.Server") 'Z002\SQL2K8'
### invoke-sqlbackup 'Z002\SqlExpress' 'pubs' $($server.BackupDirectory + "\pubs.bak")
### invoke-sqlrestore 'Z002\SqlExpress' 'pubs' $($server.BackupDirectory + "\pubs.bak") -force 
### </Usage>
### </Script>
# ---------------------------------------------------------------------------
$smoAssembly = [reflection.assembly]::LoadWithPartialName("Microsoft.SqlServer.Smo")
if (!($smoVersion))
{ Set-Variable -name SmoVersion  -value $smoAssembly.GetName().Version.Major -Scope Global -Option Constant -Description "SQLPSX variable" }
[void][reflection.assembly]::LoadWithPartialName('Microsoft.SqlServer.SMOExtended')

#######################
function Invoke-SqlBackup
{
    param($sqlserver=$(throw 'sqlserver required.'),$dbname=$(throw 'dbname required.'),$filepath=$(throw 'filepath required.')
          ,$action='Database', $description='',$name='',[switch]$force,[switch]$incremental,[switch]$copyOnly)
    
    #action can be Database or Log

    $server = new-object ("Microsoft.SqlServer.Management.Smo.Server") $sqlserver

    Write-Verbose "Invoke-SqlBackup $($server.Name) $dbname"

    $backup = new-object ("Microsoft.SqlServer.Management.Smo.Backup")
    $backupDevice = new-object ("Microsoft.SqlServer.Management.Smo.BackupDeviceItem") $filepath, 'File'

    $backup.Action = $action
    $backup.BackupSetDescription = $description
    $backup.BackupSetName = $name
    if (!$server.Databases.Contains("$dbname")) {throw 'Database $dbname does not exist on $($server.Name).'}
    $backup.Database = $dbname
    $backup.Devices.Add($backupDevice) 
    $backup.Initialize = $($force.IsPresent)
    $backup.Incremental = $($incremental.IsPresent)
    if ($copyOnly)
    { if ($server.Information.Version.Major -ge 9 -and $smoVersion -ge 10) 
      { $backup.CopyOnly = $true }
      else
      { throw 'CopyOnly is supported in SQL Server 2005(9.0) or higher with SMO version 10.0 or higher.' }
    }
    
    trap {
        $ex = $_.Exception
        Write-Output $ex.message
        $ex = $ex.InnerException
        while ($ex.InnerException)
        {
            Write-Output $ex.InnerException.message
            $ex = $ex.InnerException
        };
        continue
    }
    $backup.SqlBackup($server) 
    
    if ($?)
    { Write-Host "$action backup of $dbname to $filepath complete." }
    else
    { Write-Host "$action backup of $dbname to $filepath failed." }

} #Invoke-SqlBackup

#######################
function Invoke-SqlRestore
{
    param($sqlserver=$(throw 'sqlserver required.'),$dbname=$(throw 'dbname required.'),$filepath=$(throw 'filepath required.'),
          $action='Database',$stopat,$relocatefiles,[switch]$force,[switch]$norecovery,[switch]$keepreplication)

    #action can be Database or Log

    $server = new-object ("Microsoft.SqlServer.Management.Smo.Server") $sqlserver
 
    Write-Verbose "Invoke-SqlRestore $($server.Name) $dbname"

    $restore = new-object ("Microsoft.SqlServer.Management.Smo.Restore")
    $restoreDevice = new-object ("Microsoft.SqlServer.Management.Smo.BackupDeviceItem") $filepath, 'File'

    $restore.Action = $action
    $restore.Database = $dbname
    $restore.Devices.Add($restoreDevice) 
    $restore.ReplaceDatabase = $($force.IsPresent)
    $restore.NoRecovery = $($norecovery.IsPresent)
    $restore.KeepReplication = $($keepreplication.IsPresent)
   
    if ($stopat)
    { $restore.ToPointInTime = $stopAt }

    if ($relocatefiles)
    {
       if ($relocateFiles.GetType().Name -ne 'Hashtable')
       { throw 'Invoke-SqlRestore:Param relocateFile must be a hashtable' }

       $relocateFileAR = New-Object Collections.ArrayList
        
       foreach ($i in $relocatefiles.GetEnumerator())
        {
            $logicalName = $($i.Key); $physicalName = $($i.Value);
            $relocateFile = new-object ("Microsoft.SqlServer.Management.Smo.RelocateFile") $logicalName, $physicalName
            [void]$relocateFileAR.Add($relocateFile)
        }

        $restore.RelocateFiles = $relocateFileAR
     
    }

    trap {
        $ex = $_.Exception
        Write-Output $ex.message
        $ex = $ex.InnerException
        while ($ex.InnerException)
        {
            Write-Output $ex.InnerException.message
            $ex = $ex.InnerException
        };
        continue
    }
    $restore.SqlRestore($server) 
    
    if ($?)
    { Write-Host "$action restore of $dbname from $filepath complete." }
    else
    { Write-Host "$action restore of $dbname from $filepath failed." }

} #Invoke-SqlRestore

#######################