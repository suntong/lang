#!/usr/bin/perl -0777n

# one-char difinition
s/: (.) ;/: '$1' ;/g;

# two-chars difinition
s/: (.)(.) ;/: '$1' '$2' ;/g;

# fix "non-xxx" as "non_xxx"
s/(non|year|day|user|implementation|Form|of)-/$1_/g;

# every all-CAP words should be split into quoted characters
# else all should have leading “_”
s/[a-z]\w+/_$&/g;
s!\b[A-Z_]{2,}!join " ", split //, $&!ge;
# except for "SQL"
s/sQL/SQL/g;

# every single character should be quoted
s/(\s)([^][ ;:|{}])(?=( |\n))/$1'$2'/g;
s/'''/'\\''/;

# rules to skip
s/\n_(space|nonquote_Character|identifier_Start|not_Equals_Operator)\s*:.*?;\n//gs;


# Repeats
# 1. { }... is redundant as { } itself means repeating. Thus need to remove
# 2. word... need to be changed as word {word}
s/(\{.*?\})\s*\.\.\./$1/gs;
s/(\S+)\s*\.\.\./$1 {$1}/gs;

print;
