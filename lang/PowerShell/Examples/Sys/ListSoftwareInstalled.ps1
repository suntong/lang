# Get a List of Installed Software
# http://www.howtogeek.com/165293/how-to-get-a-list-of-software-installed-on-your-pc-with-a-single-command/
# http://blogs.msdn.com/b/powershell/archive/2009/11/15/i-can-do-that-with-1-line-of-powershell-installed-software.aspx

$wp = Get-WmiObject -Class Win32_Product | Select-Object Name, Vendor, Version  | Sort-Object Vendor, Name
$wp | export-csv AppList.csv
$wp | ft -auto > AppList.txt
# $wp | ogv
