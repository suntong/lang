#!/usr/bin/perl -0777n

# split each *e*bnf rule into two parts, lexical & syntax parts
# to STDOUT and STDERR respectively
s/(?:\n|^)(\w+)\s*:\s*(.*?;)/
    print STDERR "$1 : ${\( lcfirst($1) )} ;\n\n";
    "\n${\( lcfirst($1) )} : $2"
/gsex;

print;
