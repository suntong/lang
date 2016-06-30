#!/usr/bin/python

import sys
import getopt

def usage():
    print ("Usage:\n  Cli_getopt.py -d -g grammar")

def main(argv):
    grammar = "kant.xml"                
    try:                                
        opts, args = getopt.getopt(argv, "hg:d", ["help", "grammar="])
    except getopt.GetoptError:          
        usage()                         
        sys.exit(2)                     
    for opt, arg in opts:
        if opt in ("-h", "--help"):
            usage()                     
            sys.exit()                  
        elif opt == '-d':
            global _debug               
            _debug = 1                  
        elif opt in ("-g", "--grammar"):
            grammar = arg               

    source = "".join(args)

    print (grammar, source)

if __name__ == '__main__':
    main(sys.argv[1:])

"""

Output:

$ Cli_getopt.py -h 
Usage:
  Cli_getopt.py -d -g grammar

$ Cli_getopt.py
('kant.xml', '')

$ Cli_getopt.py a bb
('kant.xml', 'abb')

$ Cli_getopt.py -g gggg a bb ccc
('gggg', 'abbccc')

"""
