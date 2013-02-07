##----------------------------------------------------------------------------
## Porgram: Restore-Database
## Purpose: Restore SQL database properly using SQLPSX
## Authors: Antonio Sun (c) 2013, All rights reserved
##          Chad Miller (c) 2011, http://poshcode.org/2531
##---------------------------------------------------------------------------

param($sqlserver=$(throw 'sqlserver required.'), $filepath=$(throw 'filepath required.'), $dbname='')

import-module sqlserver -force


$server = get-sqlserver $sqlserver

$filepath = Resolve-Path $filepath | select -ExpandProperty Path
if ($dbname -eq '') {
    $dbname = (Get-ChildItem $filePath | select -ExpandProperty basename)  -replace('_.*','')
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