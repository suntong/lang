#!/usr/bin/python3
# http://www.ibiblio.org/g2swap/byteofpython/read/making-modules.html
# Filename: mymodule_demo2.py

# Here is a version utilising the from..import syntax

from mymodule import *

sayhi()
print('Version', version)

ex1()

"""

Output
				
foo
bar
Hi, this is mymodule speaking.
Version 0.1
foo
bar

"""
