#!/usr/bin/python

import sys

# http://www.diveintopython.net/scripts_and_streams/command_line_arguments.html
print("Method 1")

for arg in sys.argv:
    print arg

# http://stackoverflow.com/a/1009879/2125837
print("\nMethod 2a")
print "\n".join(sys.argv)
print("\nMethod 2b")
print sys.argv[1:]

# http://www.tutorialspoint.com/python/python_command_line_arguments.htm
print '\nNumber of arguments:', len(sys.argv), 'arguments.'
print 'Argument List:', str(sys.argv)
