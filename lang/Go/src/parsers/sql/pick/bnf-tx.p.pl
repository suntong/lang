#!/usr/bin/perl -0777n

# every all-CAP words should be split into quoted characters
# else all should have lcfirst then add leading “_”
s!\b[A-Z_]{2,}!join " ", split //, $&!ge;
s/[A-Z][a-z]\w+/"_${\( lcfirst($&) )}"/ge;

# fix "non-xxx" as "non_xxx"
s/(non|form|of)-/$1_/g;

# every single character should be quoted
s/(\s)([^][ ;:|{}])(?=( |\n))/$1'$2'/g;

# rules to skip
#s/\n_(space|nonquote_Character|identifier_Start|not_Equals_Operator)\s*:.*?;\n//gs;


print;
