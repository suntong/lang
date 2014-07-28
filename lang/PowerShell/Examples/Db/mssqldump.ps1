# A script to export schema to files for us to easily compare with WinMerge. Check out the options and modify as needed.
# Call from powershell command line like so:
# .\GenerateScript.ps1 "Servername" "Adventureworks" "c:\scripts\powershell\"

#Start file
#Set-ExecutionPolicy RemoteSigned
#Set-ExecutionPolicy -ExecutionPolicy:Unrestricted -Scope:LocalMachine

function GenerateDBScript([string]$serverName, [string]$dbname, [string]$scriptpath)
{
[System.Reflection.Assembly]::LoadWithPartialName("Microsoft.SqlServer.SMO") | out-null
[System.Reflection.Assembly]::LoadWithPartialName("System.Data") | out-null
#$error.clear()
#$erroractionpreference = "Continue"
$srv = new-object "Microsoft.SqlServer.Management.SMO.Server" $serverName
$srv.SetDefaultInitFields([Microsoft.SqlServer.Management.SMO.View], "IsSystemObject")
$db = New-Object "Microsoft.SqlServer.Management.SMO.Database"
$db = $srv.Databases[$dbname]
$scr = New-Object "Microsoft.SqlServer.Management.Smo.Scripter"
$deptype = New-Object "Microsoft.SqlServer.Management.Smo.DependencyType"
$scr.Server = $srv
$options = New-Object "Microsoft.SqlServer.Management.SMO.ScriptingOptions"
$options.AllowSystemObjects = $false
$options.IncludeDatabaseContext = $true
$options.IncludeIfNotExists = $false
$options.ClusteredIndexes = $true
$options.Default = $true
$options.DriAll = $true
$options.Indexes = $true
$options.NonClusteredIndexes = $true
$options.IncludeHeaders = $false
$options.ToFileOnly = $true
$options.AppendToFile = $false
$options.ScriptDrops = $false
 
#Set options for SMO.Scripter
$scr.Options = $options
 
#Tables
Foreach ($tb in $db.Tables)
{
   If ($tb.IsSystemObject -eq $FALSE)
   {
      $smoObjects = New-Object Microsoft.SqlServer.Management.Smo.UrnCollection
      $smoObjects.Add($tb.Urn)
      $options.FileName = $scriptpath +"\Tables\"+ $tb.Name + ".sql"
      $scr.Script($smoObjects)
   }
}
 
#Views
$views = $db.Views | where {$_.IsSystemObject -eq $false}
Foreach ($view in $views)
{
    if ($views -ne $null)
    {
     $options.FileName = $scriptpath +"\Views\"+ $view.Name + ".sql"
     $scr.Script($view)
    }
}
#StoredProcedures
$StoredProcedures = $db.StoredProcedures | where {$_.IsSystemObject -eq $false}
Foreach ($StoredProcedure in $StoredProcedures)
{
    if ($StoredProcedures -ne $null)
    {
     $options.FileName = $scriptpath +"\Procedures\"+ $StoredProcedure.Name + ".sql"
     $scr.Script($StoredProcedure)
    }
}  
#Functions
$UserDefinedFunctions = $db.UserDefinedFunctions | where {$_.IsSystemObject -eq $false}
Foreach ($function in $UserDefinedFunctions)
{
    if ($UserDefinedFunctions -ne $null)
    {
     $options.FileName = $scriptpath +"\Functions\"+ $function.Name + ".sql"
     $scr.Script($function)
    }
}  
#DBTriggers
$DBTriggers = $db.Triggers
foreach ($trigger in $db.triggers)
    {
    if ($DBTriggers -ne $null)
    {
        $options.FileName = $scriptpath +"\TriggersDB\"+ $function.Name + ".sql"
        $scr.Script($DBTriggers)
    }
}
Foreach ($tb in $db.Tables)
{          
    if($tb.triggers -ne $null)
    {
        foreach ($trigger in $tb.triggers)
        {
        $options.FileName = $scriptpath +"\Triggers\"+ $trigger.Name + ".sql"
        $scr.Script($trigger)
        }
    }
}
     
}

#Start script
GenerateDBScript $args[0] $args[1] $args[2]