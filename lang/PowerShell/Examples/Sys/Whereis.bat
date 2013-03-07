@echo off
REM Batch file to find first executable in the path

echo. 
REM echo Searching for %1 in %path%
set a=%~$PATH:1
If "%a%"=="" (Echo %1 not found) else (echo %1 found at %a%)