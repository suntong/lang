:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::
: Purpose: Export Perf-Counters in batch
: Authors: Tong Sun (c) 2014, All rights reserved
:::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::::

@if [%DBG_SCRIPT%]==[] echo off

SET SCRIPT=%0

SET RunId=%1
SHIFT

if "%~1"=="" goto usage

REM FOR %%M IN (1 2 3) DO ECHO %%M here %RunId%

FOR %%M IN (DB06 DB01 APP02 APP01) DO PerfCounterExport -nc -cs PERFDB02 -id %RunId% -m %%M %1 %2
goto end

:usage

echo.
echo Usage:
echo.
echo  %SCRIPT% RunId ResultFilePre
echo.
echo  - RunId: Loadtest RunId
echo  - ResultFilePre: Prefix for the ResultFile. Format: PATH/FilePrefix
echo.

:end

