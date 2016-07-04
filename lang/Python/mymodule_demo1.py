#!/usr/bin/python3
# http://www.ibiblio.org/g2swap/byteofpython/read/making-modules.html
# Filename: mymodule_demo.py

# Remember that
# - the module should be placed in the same directory as the program that we import it in,
# - or the module should be in one of the directories listed in sys.path .

import mymodule

# use the same dotted notation to access members of the module. 
mymodule.sayhi()
print('Version', mymodule.version)


"""

Output
				
$ python mymodule_demo.py
Hi, this is mymodule speaking.
Version 0.1

"""
