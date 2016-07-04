#!/usr/bin/python3
# http://www.ibiblio.org/g2swap/byteofpython/read/making-modules.html
# Filename: mymodule_demo2.py

# Here is a version utilising the from..import syntax

from mymodule import sayhi, version
# Alternative:
# from mymodule import *

sayhi()
print('Version', version)

"""

Output
				
$ python mymodule_demo.py
Hi, this is mymodule speaking.
Version 0.1

The output of mymodule_demo2.py is same as the output of mymodule_demo1.py

"""
