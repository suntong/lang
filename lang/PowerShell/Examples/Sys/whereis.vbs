'==================================================
'Find out the location of the nominated executable,
'both from the path and from the registry.
' Version: 2.0
' 21.3.2009 FNL
'
'invoke it from a Command Prompt like so:
'  cscript whereis.vbs sed
'==================================================
Set oWshShell = WScript.CreateObject("WScript.Shell")
Set oArgs     = WScript.Arguments 

If oArgs.count <> 1 Then
    WScript.Echo "Usage: which.vbs NameOfExecutable"
    WScript.Quit
End If

SearchFolders
SearchRegistry
WScript.Echo("File not found")

Sub SearchFolders
    Set oFSO      = CreateObject("Scripting.FileSystemObject")
    aAux = Split(oArgs(0), ".")
    sName = aAux(0)
    If UBound(aAux) > 0 Then sExt = "." & aAux(1)  'Extension is in parameter
    aExt = Split(LCase(oWshShell.ExpandEnvironmentStrings("%PathExt%")),";")

    aPaths = Split(oWshShell.CurrentDirectory & ";" _
      & oWshShell.ExpandEnvironmentStrings("%path%"), ";")

    For p = 0 To UBound(aPaths)
        if right(aPaths(p), 1) <> "\" then aPaths(p) = aPaths(p) & "\"
        For e = 0 To UBound(aExt)
            If oFSO.FileExists(aPaths(p) & sName & aExt(e)) Then
                WScript.Echo "Path entry:",  aPaths(p) & sName & aExt(e)
                WScript.Quit
            End If
        Next
    Next
End Sub

Sub SearchRegistry
    HKLM=&H80000002
    sKeyPath = "HKLM\Software\Microsoft\Windows\CurrentVersion\App Paths\"
    Set oWshShell = WScript.CreateObject("WScript.Shell")

    If InStr(oArgs(0), ".") = 0 _
    Then sName = LCase(oArgs(0)) & ".exe" _
    Else sName = LCase(oArgs(0))
    On Error Resume Next
    sDefault = oWshShell.RegRead(sKeyPath & sName & "\")
    If Err.number <> 0 then Exit Sub
    On Error Goto 0
    WScript.echo "Registry entry:", sDefault
    WScript.Quit
End Sub
