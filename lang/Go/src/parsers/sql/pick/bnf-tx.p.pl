#!/usr/bin/perl -0777n

# fix "Non-xxx" as "Non_xxx"
s/(Non|Form|of)-/$1_/g;

# every all-CAP words should be split into quoted characters
s!\b[A-Z_]{2,}!join " ", split //, $&!ge;

# every single character should be quoted
s/(\s)([^][ ;:|{}])(?=( |\n))/$1'$2'/g;
s/'''/'\\''/;

# rules to skip
#s/\n_(space|nonquote_Character|identifier_Start|not_Equals_Operator)\s*:.*?;\n//gs;


# Repeats
# 1. { }... is redundant as { } itself means repeating. Thus need to remove
# 2. word... need to be changed as word {word}
s/(\{.*?\})\s*\.\.\./$1/gs;
s/(\S+)\s*\.\.\./$1 {$1}/gs;

print;
