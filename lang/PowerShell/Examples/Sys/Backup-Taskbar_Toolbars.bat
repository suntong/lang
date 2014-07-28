:: Created by: Shawn Brink
:: http://www.eightforums.com
:: Tutorial:  http://www.eightforums.com/tutorials/9612-taskbar-toolbars-back-up-restore-windows-8-a.html


mkdir "C:\Tmp"

REG EXPORT HKCU\Software\Microsoft\Windows\CurrentVersion\Explorer\Streams\Desktop "C:\Tmp\Taskbar-Toolbars-Backup.reg" /y