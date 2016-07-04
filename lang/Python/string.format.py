#!/usr/bin/python

"""
The string .format() method
https://infohost.nmt.edu/tcc/help/pubs/python/web/new-str-format.html

The .format() method (added in Python 2.6) of the str type is an extremely convenient way to format text exactly the way you want it.

Here is the general form:
template.format(p0, p1, ..., k0=v0, k1=v1, ...)

The template is a string containing a mixture of one or more format codes embedded in constant text. The format method uses its arguments to substitute an appropriate value for each format code in the template.

The arguments to the .format() method are of two types. The list starts with zero or more positional arguments pi, followed by zero or more keyword arguments of the form ki=vi, where each ki is a name with an associated value vi.

Just to give you the general flavor of how this works, here's a simple conversational example. In this example, the format code ?{0}? is replaced by the first positional argument (49), and ?{1}? is replaced by the second positional argument, the string "okra".
"""

print("We have {0} hectares planted to {1}.".format(49, "okra"))

# supply the values using keyword arguments. The arguments may be supplied in any order. The keyword names must be valid Python names
print("{monster} has now eaten {city}".
      format(city='Tokyo', monster='Mothra'))

# can mix references to positional and keyword arguments
print("The {structure} sank {0} times in {1} years.".
      format(3, 2, structure='castle'))

# to include actual "{" and "}" characters in the result, double them
print("There are {0} members in set {{a}}.".format(15))
print


# Formatted Output
# http://www.python-course.eu/python3_formatted_output.php

print("First argument: {}, second one: {}".format(47,11))
print("First argument: {0}, second one: {1}".format(47,11))
print("Second argument: {1}, first one: {0}".format(47,11))
print("Second argument: {1:3d}, first one: {0:7.2f}".format(47.42,11))
print("various precions: {0:6.2f} or {0:6.3f}".format(1.4148))

# keyword parameters
print("Art: {a:5d},  Price: {p:8.2f}".format(a=453, p=59.058))
print

# justify data with the format method: "<" (left justify) or ">" (right justify)
print("{0:<20s} {1:6.2f}".format('Spam & Eggs:', 6.99))
print("{0:>20s} {1:6.2f}".format('Spam & Eggs:', 6.99))
print("{0:>20s} {1:6.2f}".format('Spam & Ham:', 7.99))
print("{0:<20s} {1:6.2f}".format('Spam & Ham:', 7.99))
print("{0:<20} {1:6.2f}".format('Spam & Ham:', 7.99))
print("{0:>20} {1:6.2f}".format('Spam & Ham:', 7.99))
# centered 
print("{0:^20} {1:6.2f}".format('Spam & Ham:', 7.99))
print

# padding 
x = 378
print("The value is {:06d}".format(x))
x = -378
print("The value is {:06d}".format(x))
print

# thousands separator
x = 5897653423
print("The value is {:,}".format(x))
print("The value is {0:6,d}".format(x))
x = 5897653423.89676
print("The value is {0:12,.3f}".format(x))
print

# Using dictionaries in "format"
print("The capital of {0:s} is {1:s}".format("Ontario","Toronto"))
print("The capital of {province} is {capital}".format(province="Ontario",capital="Toronto"))
k = {"province":"Ontario","capital":"Toronto"}
print("The capital of {province} is {capital}".format(**k))
print

capital_country = {"United States" : "Washington", 
                   "US" : "Washington", 
                   "Canada" : "Ottawa",
                   "Germany": "Berlin",
                   "France" : "Paris",
                   "England" : "London",
                   "UK" : "London",
                   "Switzerland" : "Bern",
                   "Austria" : "Vienna",
                   "Netherlands" : "Amsterdam"}

print("Countries and their capitals:")
for c in capital_country:
    print("{country}: {capital}".format(country=c, capital=capital_country[c]))
print

# using the dictionary directly
for c in capital_country:
    format_string = c + ": {" + c + "}" 
    print(format_string.format(**capital_country))
print

#print()


"""

Output:

We have 49 hectares planted to okra.
Mothra has now eaten Tokyo
The castle sank 3 times in 2 years.
There are 15 members in set {a}.

First argument: 47, second one: 11
First argument: 47, second one: 11
Second argument: 11, first one: 47
Second argument:  11, first one:   47.42
various precions:   1.41 or  1.415
Art:   453,  Price:    59.06

Spam & Eggs:           6.99
        Spam & Eggs:   6.99
         Spam & Ham:   7.99
Spam & Ham:            7.99
Spam & Ham:            7.99
         Spam & Ham:   7.99

The value is 000378
The value is -00378

The value is 5,897,653,423
The value is 5,897,653,423
The value is 5,897,653,423.897

The capital of Ontario is Toronto
The capital of Ontario is Toronto
The capital of Ontario is Toronto

Countries and their capitals:
United States: Washington
Canada: Ottawa
Austria: Vienna
Netherlands: Amsterdam
Germany: Berlin
UK: London
Switzerland: Bern
England: London
US: Washington
France: Paris

United States: Washington
Canada: Ottawa
Austria: Vienna
Netherlands: Amsterdam
Germany: Berlin
UK: London
Switzerland: Bern
England: London
US: Washington
France: Paris

"""
