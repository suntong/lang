##----------------------------------------------------------------------------
## Porgram: LastBootUpTime
## Purpose: Get Windows Last Reboot Timestamp (via WMI call)
## Authors: Tong Sun (c) 2013, All rights reserved
##---------------------------------------------------------------------------

<#
	.SYNOPSIS
		Get Windows Last Reboot Timestamp

	.DESCRIPTION
		Get Windows Last Reboot Timestamp for the given machine

	.PARAMETER ServerName 
		The name of the Server to query reboot timestamp

	.PARAMETER val
		Return the reboot timestamp value only
		
	.EXAMPLE
		.\LastBootUpTime . 
	.EXAMPLE
		.\LastBootUpTime OtherServer

	.EXAMPLE
		.\LastBootUpTime -val .
	.EXAMPLE
		.\LastBootUpTime . -val

	.EXAMPLE
		Get-Help .\LastBootUpTime -detailed
#>

param(
	[Parameter(Mandatory=$true)] [ValidateNotNullOrEmpty()]
	[string] $ServerName, 
	
	[Parameter()]
	[Switch] $val

)
#$val

$wmi = Get-WmiObject -Class Win32_OperatingSystem -Computer $ServerName
$LastBootUpTime = $wmi.ConvertToDateTime($wmi.LastBootUpTime)
if ($val) { 
	$LastBootUpTime
	exit
}

write-host "`nThe '$ServerName' was last booted at`n  $($LastBootUpTime.dayOfWeek ), $LastBootUpTime"

# calculate the difference between two dates
$today = Get-Date
$tspan=New-TimeSpan $LastBootUpTime $today
$diffDays=($tspan).days
write-host "  $diffDays days ago.`n"

<#

Refs:

Using Windows PowerShell to Work with Dates
http://blogs.technet.com/b/heyscriptingguy/archive/2010/08/02/using-windows-powershell-to-work-with-dates.aspx

Date arithmetic within PowerShell
http://iannotes.wordpress.com/2012/06/02/date-arithmetic-powershell/

Date differences
http://msmvps.com/blogs/richardsiddaway/archive/2010/04/02/date-differences.aspx

#>

