# PowerShell ABC's - P is for Parameters by Joe Pruitt
# https://devcentral.f5.com/blogs/us/powershell-abcs-p-is-for-parameters

# Functions are defined by the function statement
# The format of the parameter list is identical to that of the param statement. 

function foo([string]$foo = "foo", [string]$bar = "bar")
{
    Write-Host "Arg: $foo";
    Write-Host "Arg: $bar";
}

# The param statement is supported in functions as well so if you do not wish 
# to specify it in the function declaration, you can do so in the first line

function bar()
{
    param([string]$foo = "foo", [string]$bar = $(throw "-bar is required."));
    Write-Host "Arg `$foo: $foo";
    Write-Host "Arg `$bar: $bar";
    Write-Host "$args";
}

foo;

bar -foo "abc" -bar def;
# it's permitted to leave the name out and let the interpreter figure out 
# what parameter is it from it's position on the command line. 
bar "abc2" def3;

# will trigger a throw
bar -foo "abc";
