#!/usr/bin/python
# http://www.cyberciti.biz/faq/python-command-line-arguments-argv-example/

import argparse
__author__ = 'nixCraft'
 
parser = argparse.ArgumentParser(description='This is a demo script by nixCraft.')
parser.add_argument('-i','--input', help='Input file name',required=True)
parser.add_argument('-o','--output',help='Output file name', required=True)
args = parser.parse_args()
 
## show values ##
print ("Input file: %s" % args.input )
print ("Output file: %s" % args.output )

"""

Output:


$ Cli_argparse2.py 
usage: Cli_argparse2.py [-h] -i INPUT -o OUTPUT
Cli_argparse2.py: error: argument -i/--input is required

$ Cli_argparse2.py -h
usage: Cli_argparse2.py [-h] -i INPUT -o OUTPUT

This is a demo script by nixCraft.

optional arguments:
  -h, --help            show this help message and exit
  -i INPUT, --input INPUT
                        Input file name
  -o OUTPUT, --output OUTPUT
                        Output file name

$ Cli_argparse2.py -i input.txt
usage: Cli_argparse2.py [-h] -i INPUT -o OUTPUT
Cli_argparse2.py: error: argument -o/--output is required

$ Cli_argparse2.py -z
usage: Cli_argparse2.py [-h] -i INPUT -o OUTPUT
Cli_argparse2.py: error: argument -i/--input is required

$ Cli_argparse2.py -i input.txt -o output.txt
Input file: input.txt
Output file: output.txt

$ Cli_argparse2.py -i input.txt -o output.txt -a
usage: Cli_argparse2.py [-h] -i INPUT -o OUTPUT
Cli_argparse2.py: error: unrecognized arguments: -a

"""
