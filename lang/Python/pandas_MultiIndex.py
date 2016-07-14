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

print(s1.index.to_series().str.join(' ') + ' ' + s1.astype(str))

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

print(p.index.to_series().str.join(' ') )

# X print(p.index.to_series().str.join(' ') + ' ' + p.astype(str))
# cannot join with no level specified and no overlapping names

x=p.reset_index()
print(x)

#print(x.['A'] + ' ' +x['B'])
# SyntaxError: invalid syntax
#print(x.[A] + ' ' +x[B])
# SyntaxError: invalid syntax

print(x.A + ' - ' + x.B)

x.B = x.A + ' - ' + x.B
print(x)

# Conditionally set pandas dataframe column values
# http://stackoverflow.com/questions/38228082/conditionally-set-pandas-dataframe-column-values

df = pd.DataFrame({'data1': np.random.randn(100),'data2': np.random.randn(100)})
print(df.head())

Col = 'data1'
print(df[Col].head())
df.data1 = df.data1 +.1
print(df[Col].head())
#df.Col = df.Col + 1
df[Col] = df[Col] + .1
print(df[Col].head())
print("\nBetween .25 & .35")
print(df[(df[Col] >=.25) & (df[Col] <= .35)])
df[(df[Col] >=.25) & (df[Col] <= .35)] = df[(df[Col] >=.25) & (df[Col] <= .35)]+.1 
print("Between .35 & .45")
print(df[(df[Col] >=.35) & (df[Col] <= .45)])
#gp = df.groupby(level=('data1', 'data2'))
# TypeError: unorderable types: tuple() > int()

df = pd.DataFrame({'data1': np.random.choice(['foo', 'bar'], 20),
                   'data2': np.random.choice(['I1', 'I2', 'I3', 'I4'], 20),
                   'data3': np.random.randn(20)})
print(df)

Col = 'data1'
print(df[Col].head())
df[Col] = df[Col] + "1"
print(df[Col].head())

#means.loc[w.female != 'female', 'female'] = 0

# https://gist.github.com/suntong/8740c0b6a67cfc79856b
# Generate the data
ix3 = pd.MultiIndex.from_arrays([['a', 'a', 'a', 'a', 'b', 'b', 'b', 'b'], ['foo', 'foo', 'bar', 'bar', 'foo', 'foo', 'bar', 'bar']], names=['letter', 'word'])

df3 = pd.DataFrame({'data1': [3, 2, 4, 3, 2, 4, 3, 2], 'data2': [6, 5, 7, 5, 4, 5, 6, 5]}, index=ix3)

# Group by index labels and take the means and standard deviations for each group
gp3 = df3.groupby(level=('letter', 'word'))

means = gp3.mean()
errors = gp3.std()

print(means)
print(errors)

print(means[Col])
means[means[Col]<3] = 0
print(means[Col])

means = gp3.mean()
means[Col] = means[Col] * 10
means[Col] = means[Col].astype(int)
ThresholdD = "30"
means[means[Col]<float(ThresholdD)] = 0
print(means[Col])
