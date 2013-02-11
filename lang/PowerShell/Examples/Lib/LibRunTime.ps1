##----------------------------------------------------------------------------
## Porgram: LibRunTime
## Purpose: General runtime supporting functions
## Authors: Antonio Sun (c) 2013, All rights reserved
##---------------------------------------------------------------------------

function start-logging ($logName='', $logParams=@{NoCLobber=$true;Append=$true;Path="$env:temp\log.txt"}) {
	#get current login info
	#$CS = Gwmi Win32_ComputerSystem -Comp "."
	#$LogonHost=$CS.Name
	#$LogonUser=$CS.UserName

    if (-not ($logName -eq "")) {
        $logParams.Path = "$env:temp\$logName"
    }

	start-transcript @logParams | Out-Null 
	#write-host "`n`nLogging started at" (get-date -format s) "on $LogonHost by $LogonUser`n"
    return "$logParams.Path"
}

function stop-logging ($logFile){
    stop-transcript
    "`r`n" | Out-File -FilePath $logFile -Append
}


## by Vidrine
## http://poshcode.org/3502

	<#
		.SYNOPSIS
			Writes logging information to screen and log file simultaneously.

		.DESCRIPTION
			Writes logging information to screen and log file simultaneously. Supports multiple log levels.

		.PARAMETER Message
			The message to be logged.

		.PARAMETER Level
			The type of message to be logged.
			
		.PARAMETER NoConsoleOut
			Specifies to not display the message to the console.
			
		.PARAMETER ConsoleForeground
			Specifies what color the text should be be displayed on the console. Ignored when switch 'NoConsoleOut' is specified.
		
		.PARAMETER Indent
			The number of spaces to indent the line in the log file.

		.PARAMETER Path
			The log file path.
		
		.PARAMETER Clobber
			Existing log file is deleted when this is specified.
		
		.PARAMETER EventLogName
			The name of the system event log, e.g. 'Application'.
		
		.PARAMETER EventSource
			The name to appear as the source attribute for the system event log entry. This is ignored unless 'EventLogName' is specified.
		
		.PARAMETER EventID
			The ID to appear as the event ID attribute for the system event log entry. This is ignored unless 'EventLogName' is specified.

		.PARAMETER LogEncoding
			The text encoding for the log file. Default is ASCII.
		
		.EXAMPLE
			PS C:\> Write-Log -Message "It's all good!" -Path C:\MyLog.log -Clobber -EventLogName 'Application'

		.EXAMPLE
			PS C:\> Write-Log -Message "Oops, not so good!" -Level Error -EventID 3 -Indent 2 -EventLogName 'Application' -EventSource "My Script"

		.INPUTS
			System.String

		.OUTPUTS
			No output.
			
		.NOTES
			Revision History:
				2011-03-10 : Andy Arismendi - Created.
				2011-07-23 : Will Steele - Updated.
				2011-07-23 : Andy Arismendi 
					- Added missing comma in param block. 
					- Added support for creating missing directories in log file path.
				2012-03-10 : Pat Richard
					- Added validation sets to $ConsoleForeground and $EventLogName
					- Changed formatting of $msg so that only $message is indented instead of entire line (looks cleaner)
					- suppressed output when creating path/file
	#>

function Write-Log {

	#region Parameters
	
		[cmdletbinding()]
		Param(
			[Parameter(ValueFromPipeline=$true,Mandatory=$true)] [ValidateNotNullOrEmpty()]
			[string] $Message,

			[Parameter()] [ValidateSet(“Error”, “Warn”, “Info”)]
			[string] $Level = “Info”,
			
			[Parameter()]
			[Switch] $NoConsoleOut,
			
			[Parameter()]
			[ValidateSet("Black", "DarkMagenta", "DarkRed", "DarkBlue", "DarkGreen", "DarkCyan", "DarkYellow", "Red", "Blue", "Green", "Cyan", "Magenta", "Yellow", "DarkGray", "Gray", "White")]
			[String] $ConsoleForeground = 'White',
			
			[Parameter()] [ValidateRange(1,30)]
			[Int16] $Indent = 0,

			[Parameter()]
			[IO.FileInfo] $Path = ”$env:temp\PowerShellLog.txt”,
			
			[Parameter()]
			[Switch] $Clobber,
			
			[Parameter()]
			[ValidateSet("Application","System","Security")]
			[String] $EventLogName,
			
			[Parameter()]
			[String] $EventSource,
			
			[Parameter()]
			[Int32] $EventID = 1,

			[Parameter()]
			[String] $LogEncoding = "ASCII"
		)
		
	#endregion

	Begin {}

	Process {
		try {			
			$msg = "{0}`t{1}`t{2}{3}" -f (Get-Date -Format "yyyy-MM-dd HH:mm:ss"), $Level.ToUpper(), (" " * $Indent), $Message			
			if ($NoConsoleOut -eq $false) {
				switch ($Level) {
					'Error' { Write-Error $Message }
					'Warn' { Write-Warning $Message }
					'Info' { Write-Host ('{0}{1}' -f (" " * $Indent), $Message) -ForegroundColor $ConsoleForeground}
				}
			}

			if (-not $Path.Exists) {
				New-Item -Path $Path.FullName -ItemType File -Force | Out-Null
			}
			
			if ($Clobber) {
				$msg | Out-File -FilePath $Path -Encoding $LogEncoding -Force
			} else {
				$msg | Out-File -FilePath $Path -Encoding $LogEncoding -Append
			}
			
			if ($EventLogName) {
			
				if (-not $EventSource) {
					$EventSource = ([IO.FileInfo] $MyInvocation.ScriptName).Name
				}
			
				if(-not [Diagnostics.EventLog]::SourceExists($EventSource)) { 
					[Diagnostics.EventLog]::CreateEventSource($EventSource, $EventLogName) 
		        } 

				$log = New-Object System.Diagnostics.EventLog  
			    $log.set_log($EventLogName)  
			    $log.set_source($EventSource) 
				
				switch ($Level) {
					“Error” { $log.WriteEntry($Message, 'Error', $EventID) }
					“Warn”  { $log.WriteEntry($Message, 'Warning', $EventID) }
					“Info”  { $log.WriteEntry($Message, 'Information', $EventID) }
				}
			}

		} catch {
			throw “Failed to create log entry in: ‘$Path’. The error was: ‘$_’.”
		}
	}

	End {}

}