#!/usr/bin/env python3

import traceback

def f():
    g()

def g():
    raise Exception('asdf')

try:
    g()
except Exception as e:
    exc_obj = e

tb_str = ''.join(traceback.format_exception(exc_obj))
print(tb_str)
