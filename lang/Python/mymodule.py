#!/usr/bin/python
# http://www.ibiblio.org/g2swap/byteofpython/read/making-modules.html
# Filename: mymodule.py

def sayhi():
        print('Hi, this is mymodule speaking.')

version = '0.1'

# The above was a sample module. As you can see, there is nothing particularly special about compared to our usual Python program.

global_var = 'foo'

def ex1():
        local_var = 'bar'
        print(global_var)
        print(local_var)

ex1()

# End of mymodule.py
