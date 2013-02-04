# PowerShell ABC's - P is for Parameters by Joe Pruitt
# https://devcentral.f5.com/blogs/us/powershell-abcs-p-is-for-parameters

# The param statement

# PowerShell provides a much more convenient way to declare formal parameters
# for a script with the param statement.  The param statement must be the first
# executable line in the script with only comment or empty lines preceding it.
# The format of the param statement is 

# param([type]$p1 = , [type]$p2 = , ...)

# Where the type literal "[type]" and initialization values "= " are optional
# components to allow for type specification and specifying initial values 
# respectively.

# How to Handle Command Line Arguments in PowerShell
# http://stackoverflow.com/questions/2157554/how-to-handle-command-line-arguments-in-powershell

# You can also assign default values to your params and read them from console 
# if not available:

 param (
    [string]$foo = "foo", [string]$bar = "bar"
    # For switch parameters you specify the parameter but leave out the argument.  
    # The value is assigns based on whether the parameter is present or not. 
    ,[switch]$Recurse
    ,[string]$server = "http://defaultserver"
    #,[string]$Password = $( Read-Host "Input password, please" )
 )

Write-Host "All args: $args";

Write-Host "Switch: $Recurse";
Write-Host "Arg `$foo: $foo";
Write-Host "Arg `$bar: $bar";

# The scriptblock is a block of script code that exists as an object reference 
# but does not require a name. Scriptblocks are also known as anonymous 
# functions or lambda expressions in other languages. 

# the "&" is how to tell PowerShell to invoke the literal as a block of code.
& {param([string]$foo = "foo2", [string]$bar = "bar2") 
    Write-Host "Arg: $foo"; Write-Host "Arg: $bar"; }
