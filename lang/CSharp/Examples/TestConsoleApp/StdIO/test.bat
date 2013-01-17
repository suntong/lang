START cmd /C "CALL test2.bat > tt0 2>&1"
call test2.bat > tt1 2>&1
START cmd /C CALL "test2.cmd"
call "cmd /c start test3.cmd"
echo Foo
pause
test2.bat > tt1 2>&1