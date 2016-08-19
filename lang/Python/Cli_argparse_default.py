#!/usr/bin/python3
# http://stackoverflow.com/questions/15301147/python-argparse-default-value-or-specified-value

import argparse
parser = argparse.ArgumentParser()
parser.add_argument('-e','--example', nargs='?', const=1, type=int)
parser.add_argument('-d','--default', default=1, type=int)
parser.add_argument("--myArg", nargs='?', default=argparse.SUPPRESS)
# http://stackoverflow.com/questions/30487767/check-if-argparse-optional-argument-is-set-or-not
args = parser.parse_args()
print(args)

#if args.myArg is not None:
if hasattr(args, 'myArg'):
    print ("myArg has been set (value is %s)" % args.myArg)


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



"""

Adding the --default

With the `nargs='?',`

    $ Cli_argparse_default.py
    Namespace(default=1, example=None)

    $ Cli_argparse_default.py -d
    Namespace(default=None, example=None)

Without the `nargs='?',`

    $ Cli_argparse_default.py -d
    usage: Cli_argparse_default.py [-h] [-e [EXAMPLE]] [-d DEFAULT]
    Cli_argparse_default.py: error: argument -d/--default: expected one argument

    $ Cli_argparse_default.py
    Namespace(default=1, example=None)

    $ Cli_argparse_default.py -d 5 
    Namespace(default=5, example=None)

"""

"""

With the
  parser.add_argument("--myArg", nargs='?')

$ Cli_argparse_default.py
Namespace(default=1, example=None, myArg=None)

$ Cli_argparse_default.py --myArg
Namespace(default=1, example=None, myArg=None)

I.e., there is no way to check whether --myArg is specified or not!

With the
  parser.add_argument("--myArg", nargs='?', default=argparse.SUPPRESS)

$ Cli_argparse_default.py
Namespace(default=1, example=None)

$ Cli_argparse_default.py --myArg
Namespace(default=1, example=None, myArg=None)
myArg has been set (value is None)

$ Cli_argparse_default.py --myArg=abc
Namespace(default=1, example=None, myArg='abc')
myArg has been set (value is abc)

"""
