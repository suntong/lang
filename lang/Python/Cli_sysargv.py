#!/usr/bin/python

import sys

# http://www.diveintopython.net/scripts_and_streams/command_line_arguments.html
print("Method 1")

for arg in sys.argv:
    print arg

# http://stackoverflow.com/a/1009879/2125837
print("Method 2a")
print "\n".join(sys.argv)
print("Method 2b")
print sys.argv[1:]

