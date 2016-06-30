#!/usr/bin/python

from string import Template
#open the file
filein = open( 'StringTemplate.tmpl' )
#read it
src = Template( filein.read() )
#document data
title = "This is the title"
subtitle = "And this is the subtitle"
list = ['first', 'second', 'third']
d={ 'title':title, 'subtitle':subtitle, 'list':'\n'.join(list) }
#do the substitution
result = src.substitute(d)
print result

"""

Python technique or simple templating system for plain text output
http://stackoverflow.com/questions/6385686/python-technique-or-simple-templating-system-for-plain-text-output

Simple solutions to simple problems

Use the standard library string template:
http://docs.python.org/library/string.html#template-strings

Output:

$ StringTemplate.py
This is the title
...
And this is the subtitle
...
first
second
third


"""
