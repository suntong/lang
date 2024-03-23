import traceback

try:
    # Code that might raise an exception
    x = 10 / 0
except ZeroDivisionError:
    traceback.print_exc()  # Print the full traceback to standard error
