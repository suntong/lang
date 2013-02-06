# PowerShell ABC's - P is for Parameters by Joe Pruitt
# https://devcentral.f5.com/blogs/us/powershell-abcs-p-is-for-parameters

Write-Host "Num Args:" $args.Length;
foreach ($arg in $args)
{
  Write-Host "Arg: $arg";
}

# command line arguments in Powershell
# http://stackoverflow.com/questions/3385052/command-line-arguments-in-powershell
for($i=0;$i -lt $args.length;$i++)
{
    "Arg $i is <$($args[$i])>"
}

# Wrong -- If you wrap $args[0], for example, in double quotes, it will interpret $args and stop, 
# never getting to the [], and therefore printing off the $arg, the entire array of command line arguments.
write-host "`$args`[0`] = $args[0]"

# You need to wrap the variable into $(..) like this
write-host "`$args`[0`] = $($args[0])"
write-host "`$args`[1`] = $($args[1])"
write-host "`$args`[2`] = $($args[2])"

# This applies for any expression that is not simple scalar variable
$date = get-date
write-host "day: $($date.day)"
