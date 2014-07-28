:: Created by: Shawn Brink
:: http://www.sevenforums.com
:: Tutorial:  http://www.sevenforums.com/tutorials/212923-taskbar-backup-restore-pinned-items-windows-7-a.html


@ECHO OFF

:choice1
set /P c=Do you want to backup or restore your pinned taskbar items [B/R]?
if /I "%c%" EQU "B" goto :backup
if /I "%c%" EQU "R" goto :choice2
goto :choice1


:backup
echo.
echo Creating backup now.
echo.
mkdir "%userprofile%\Desktop\Taskbar-Pinned-Items-Backup\TaskBar"
copy /y "%AppData%\Microsoft\Internet Explorer\Quick Launch\User Pinned\TaskBar" "%userprofile%\Desktop\Taskbar-Pinned-Items-Backup\TaskBar"
REG EXPORT HKCU\Software\Microsoft\Windows\CurrentVersion\Explorer\Taskband %userprofile%\Desktop\Taskbar-Pinned-Items-Backup\Taskbar-Pinned-Items-Backup.reg /y
echo.
pause 
exit

:choice2
IF NOT EXIST "%userprofile%\Desktop\Taskbar-Pinned-Items-Backup\TaskBar" goto :response
goto :restore

:response
echo.
echo Please place the "Taskbar-Pinned-Items-Backup" folder on your desktop now.
echo.
goto :choice1


:restore
echo.
echo Restoring backup now.
echo.
DEL /F /S /Q /A "%AppData%\Microsoft\Internet Explorer\Quick Launch\User Pinned\TaskBar\*"
copy /y "%userprofile%\Desktop\Taskbar-Pinned-Items-Backup\TaskBar" "%AppData%\Microsoft\Internet Explorer\Quick Launch\User Pinned\TaskBar" 
REG DELETE HKCU\Software\Microsoft\Windows\CurrentVersion\Explorer\Taskband /F
REG IMPORT %userprofile%\Desktop\Taskbar-Pinned-Items-Backup\Taskbar-Pinned-Items-Backup.reg
echo.
echo.
echo.
echo Waiting to restart explorer to apply.
echo.
pause
taskkill /f /im explorer.exe
start explorer.exe 
exit