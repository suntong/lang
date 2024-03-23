import traceback


def stack_lvl_3():
    #print(1/0)
    raise Exception('a1', 'b2', 'c3')


def stack_lvl_2():
    stack_lvl_3()


def stack_lvl_1():
    stack_lvl_2()


print('Started at stack lvl 0')

try:
    stack_lvl_1()
except Exception as exc:
    tb = traceback.TracebackException.from_exception(exc)
    print(''.join(tb.stack.format()))

print('End')
