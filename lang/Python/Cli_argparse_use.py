#!/usr/bin/python3
# http://stackoverflow.com/questions/15301147/python-argparse-default-value-or-specified-value

import sys
import argparse

parser = argparse.ArgumentParser()
parser.add_argument('-e','--example', default="abc")
args = parser.parse_args()
print(args)
print(args.example)

v = args.example
print(v)

for arg in sys.argv:
    print(arg)

"""

$ Cli_argparse_use.py
Namespace(example='abc')
abc
abc
./Cli_argparse_use.py

$ Cli_argparse_use.py bcd
usage: Cli_argparse_use.py [-h] [-e EXAMPLE]
Cli_argparse_use.py: error: unrecognized arguments: bcd

"""
