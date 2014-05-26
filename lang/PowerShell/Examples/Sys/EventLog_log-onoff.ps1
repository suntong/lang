# check a remote computer's user logon/logoff sessions and times
# http://stackoverflow.com/questions/11763481/user-logon-logoff-information-using-powershell

# To run:
#   .\EventLog_log-onoff.ps1

$UserProperty = @{n="User";e={(New-Object System.Security.Principal.SecurityIdentifier $_.ReplacementStrings[1]).Translate([System.Security.Principal.NTAccount])}}
$TypeProperty = @{n="Action";e={if($_.EventID -eq 7001) {"Logon"} else {"Logoff"}}}
$TimeProeprty = @{n="Time";e={$_.TimeGenerated}}

Get-EventLog System -Source Microsoft-Windows-Winlogon | select $UserProperty,$TypeProperty,$TimeProeprty

# To check a remote computer, use the Get-EventLog cmdlet's ComputerName parameter:
# Get-EventLog System -Source Microsoft-Windows-Winlogon -ComputerName $computer | ...
