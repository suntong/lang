csc /target:module Mod.cs
csc /target:library Lib.cs
csc /addmodule:Mod.netmodule /reference:Lib.dll Prog.cs
