'' To remove the files in C:\Windows\Installer?
'' http://blogging.cwl.cc/2012/06/is-it-safe-remove-files-in.html

'' Identify which patches are registered on the system, and to which
'' products those patches are installed.
''
'' Copyright (C) Microsoft Corporation. All rights reserved.
''
'' THIS CODE AND INFORMATION IS PROVIDED "AS IS" WITHOUT WARRANTY OF ANY
'' KIND, EITHER EXPRESSED OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
'' IMPLIED WARRANTIES OF MERCHANTABILITY AND/OR FITNESS FOR A
'' PARTICULAR PURPOSE.

Option Explicit

Dim msi : Set msi = CreateObject("WindowsInstaller.Installer")

'Output CSV header
WScript.Echo "ProductCode, PatchCode, PatchLocation"

' Enumerate all products
Dim products : Set products = msi.Products
Dim productCode

For Each productCode in products
	' For each product, enumerate its applied patches
	Dim patches : Set patches = msi.Patches(productCode)
	Dim patchCode

	For Each patchCode in patches
		' Get the local patch location
		Dim location : location = msi.PatchInfo(patchCode, "LocalPackage")

		WScript.Echo productCode & ", " & patchCode & ", " & location
	Next
Next
