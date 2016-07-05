#!/usr/bin/python3

#from random import randint

import numpy as np
import pandas as pd

# http://stackoverflow.com/questions/17921010/how-to-query-multiindex-index-columns-values-in-pandas

A = np.array([1.1, 1.1, 3.3, 3.3, 5.5, 6.6])
B = np.array([111, 222, 222, 333, 333, 777])
C = np.random.randint(10, 99, 6)
df = pd.DataFrame(list(zip(A, B, C)), columns=['A', 'B', 'C'])
df.set_index(['A', 'B'], inplace=True)
print(df)

x = df.reset_index()
print(x)

# http://stackoverflow.com/questions/38193003/concat-multiindex-pandas-dataframe-columns

arrays = [np.array(['bar', 'bar', 'baz', 'baz', 'foo', 'foo', 'qux', 'qux']),
          np.array(['one', 'two', 'one', 'two', 'one', 'two', 'one', 'two'])]

s = pd.Series(np.random.randn(8), index=arrays)
print(s)

s1 = s
s2 = s

s2.index = s2.index.to_series().str.join(' ')
print(s2)


# http://stackoverflow.com/questions/34292076/pandas-bar-plot-how-to-annotate-grouped-horizontal-bar-charts

df = pd.DataFrame({'A': np.random.choice(['foo', 'bar'], 100),
                   'B': np.random.choice(['one', 'two', 'three'], 100),
                   'C': np.random.choice(['I1', 'I2', 'I3', 'I4'], 100),
                   'D': np.random.randint(-10,11,100),
                   'E': np.random.randn(100)})

p = pd.pivot_table(df, index=['A','B'], columns='C', values='D')
e = pd.pivot_table(df, index=['A','B'], columns='C', values='E')

df = df.set_index(['A','B','C'])

print(df.head())
print(p)
print(e)
