#!/usr/bin/python
# http://www.tutorialspoint.com/python/python_command_line_arguments.htm

import sys, getopt

def main(argv):
   inputfile = ''
   outputfile = ''
   try:
      opts, args = getopt.getopt(argv,"hi:o:",["ifile=","ofile="])
   except getopt.GetoptError:
      print 'test.py -i <inputfile> -o <outputfile>'
      sys.exit(2)
   for opt, arg in opts:
      if opt == '-h':
         print 'test.py -i <inputfile> -o <outputfile>'
         sys.exit()
      elif opt in ("-i", "--ifile"):
         inputfile = arg
      elif opt in ("-o", "--ofile"):
         outputfile = arg
   print 'Input file is:', inputfile
   print 'Output file is:', outputfile

if __name__ == "__main__":
   main(sys.argv[1:])

"""

Output:


$ Cli_getopt2.py -h
test.py -i <inputfile> -o <outputfile>

$ Cli_getopt2.py -i a -o 
test.py -i <inputfile> -o <outputfile>

$ Cli_getopt2.py -i a 
Input file is: a
Output file is: 

$ Cli_getopt2.py -i a -o b
Input file is: a
Output file is: b

"""
