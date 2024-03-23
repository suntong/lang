import traceback


def stack_lvl_3():
    print(1/0)


def stack_lvl_2():
    stack_lvl_3()


def stack_lvl_1():
    stack_lvl_2()


print('Started at stack lvl 0')

try:
    stack_lvl_1()
except Exception as exc:
    traceback.print_exc()  # Print the full traceback to standard error

print('End')
