# Configure Desktop Module
# http://gallery.technet.microsoft.com/scriptcenter/cc98b735-c943-4158-8b66-9c8aa78eafbd?SRC=Home
# http://blogs.technet.com/b/heyscriptingguy/archive/2011/06/27/don-t-write-scripts-write-powershell-modules.aspx

# This module is comprised of three different scripts. It will configure Windows Explorer settings, Internet Explorer download and Internet Explorer home pages. It also has the ability to detect current configuration of these items.


Function Get-ExplorerSettings() 
{ 
 $RegExplorer =  Get-ItemProperty -Path HKCU:\Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced 
"Display hidden files and folders is $($RegExplorer.SuperHidden)" 
"Hide File extensions is set to $($RegExplorer.HideFileExt)" 
"Show system files and folders is set to $($RegExplorer.ShowSuperHidden)" 
"Hide desktop icons $($RegExplorer.HideIcons)" 
"Use Web view for folders $($RegExplorer.WebView)" 
"Display correct file name capitalization $($RegExplorer.DontPrettyPath)" 
"Prevent automatically locate file shares and printers $($RegExplorer.NoNetCrawling)" 
} #end Get-ExplorerSettings 
 
 
Function Set-ExplorerSettings() 
{ 
 $RegValues = @{ 
                "SuperHidden" = 1 ; 
                "HideFileExt" = 0 ; 
                "ShowSuperHidden" = 0 ; 
                "HideIcons" = 0 ; 
                "WebView" = 0 ; 
                "DontPrettyPath" = 1 ; 
                "NoNetCrawling" = 1 
                                    } 
 $path = "HKCU:\Software\Microsoft\Windows\CurrentVersion\Explorer\Advanced" 
 ForEach ($key in $RegValues.Keys) 
  { 
    Set-ItemProperty -path $path -name $key -value $RegValues[$key] 
   "Setting $path $($key) to $($RegValues[$key])" 
  }  
 
} #end Set-ExplorerSettings 
 
Function Get-ieStartPage() 
{ 
 Param ($computer = $env:computername) 
 $hkcu = 2147483649 
 $key = "Software\Microsoft\Internet Explorer\Main" 
 $property = "Start Page" 
 $property2 = "Secondary Start Pages" 
 $wmi = [wmiclass]"\\$computer\root\default:stdRegProv" 
 ($wmi.GetStringValue($hkcu,$key,$property)).sValue 
 ($wmi.GetMultiStringValue($hkcu,$key, $property2)).sValue 
} #end Get-ieStartPage 
 
Function Set-ieStartPage() 
{ 
 Param($computer = $env:computername) 
  $hkcu = 2147483649 
  $key = "Software\Microsoft\Internet Explorer\Main" 
  $property = "Start Page" 
  $property2 = "Secondary Start Pages" 
  $value = "http://www.microsoft.com/technet/scriptcenter/default.mspx" 
  $aryValues = "http://social.technet.microsoft.com/Forums/en/ITCG/threads/", 
  "http://www.microsoft.com/technet/scriptcenter/resources/qanda/all.mspx" 
  $wmi = [wmiclass]"\\$computer\root\default:stdRegProv" 
  $rtn = $wmi.SetStringValue($hkcu,$key,$property,$value) 
  $rtn2 = $wmi.SetMultiStringValue($hkcu,$key,$property2,$aryValues) 
  "Setting $property returned $($rtn.returnvalue)" 
  "Setting $property2 returned $($rtn2.returnvalue)" 
} #end Set-ieStartPage 
 

Function Set-IEDownload 
{ 
param($connections=10) 
$ErrorActionPreference = "SilentlyContinue" 
$error.Clear() 
 
$PATH_KEY = 'hklm:\SOFTWARE\Microsoft\Internet Explorer\MAIN' +  
  '\FeatureControl\FEATURE_MAXCONNECTIONSPERSERVER' 
$PATH_KEY_1_0 = 'hklm:\SOFTWARE\Microsoft\Internet Explorer\MAIN' + 
  '\FeatureControl\FEATURE_MAXCONNECTIONSPER1_0SERVER' 
$current_key = (Get-ItemProperty $PATH_KEY).'iexplorer.exe' 
$current_key_1_0 = (Get-ItemProperty $PATH_KEY_1_0).'iexplorer.exe' 
 
if ($current_key -eq $null -or $current_key_1_0 -eq $null)  
 { 
  Write-Host -foreground green "Value not currently set, creating new Key with value = $connections." 
  New-ItemProperty $PATH_KEY -Name 'iexplorer.exe' -Value $connections | Out-Null 
  New-ItemProperty $PATH_KEY_1_0 -Name 'iexplorer.exe' -Value $connections | Out-Null 
  if($error.Count -ne 0)  
    { "This script requires admin rights." ; $error[0] ; exit } 
  else { "Registry keys successfully created" } 
 } #end if registry key not present 
else  
 { 
  Write-Host -foreground green "Current value HTTP 1.0 $current_key_1_0" 
  Write-Host -foreground green "Current value for greater than HTTP 1.0 $current_key" 
 
  if($connections -eq $current_key)  
    {"connection value does not need updating" ; exit } 
  else 
    { 
     Set-ItemProperty $PATH_KEY -Name 'iexplorer.exe' -Value $connections 
     Set-ItemProperty $PATH_KEY_1_0 -Name 'iexplorer.exe' -Value $connections 
     if($error.Count -ne 0)  
       { "This script requires admin rights." ; $error[0] ; exit } 
     else { "Registry value successfully modified" } 
    } #end else connection key update 
 } #end else registry key present 
} #end function Set-IEDownload