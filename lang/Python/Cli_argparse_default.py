#!/usr/bin/python3
# http://stackoverflow.com/questions/15301147/python-argparse-default-value-or-specified-value

import argparse
parser = argparse.ArgumentParser()
parser.add_argument('-e','--example', nargs='?', const=1, type=int)
args = parser.parse_args()
print(args)

"""

Output:

$  Cli_argparse_default.py
Namespace(example=None)

$ Cli_argparse_default.py --example
Namespace(example=1)

$ Cli_argparse_default.py -e 3
Namespace(example=3)

"""

"""

- nargs='?' means 0-or-1 arguments
- const=1 sets the default when there are 0 arguments
- type=int converts the argument to int

If you want test.py to set example to 1 even if no --example is specified, then include default=1. That is, with

parser.add_argument('--example', nargs='?', const=1, type=int, default=1)

then

% test.py 
Namespace(example=1)

"""
