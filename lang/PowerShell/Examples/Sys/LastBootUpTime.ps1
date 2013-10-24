##----------------------------------------------------------------------------
## Porgram: LastBootUpTime
## Purpose: Get Windows Last Reboot Timestamp (via WMI call)
## Authors: Antonio Sun (c) 2013, All rights reserved
##---------------------------------------------------------------------------

param(
	[Parameter(Mandatory=$true)] [ValidateNotNullOrEmpty()]
	[string] $ServerName 
)

$wmi = Get-WmiObject -Class Win32_OperatingSystem -Computer $ServerName
$wmi.ConvertToDateTime($wmi.LastBootUpTime)