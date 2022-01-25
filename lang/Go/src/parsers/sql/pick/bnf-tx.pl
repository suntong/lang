#!/usr/bin/perl -0777n

# one-char difinition
s/: (.) ;/: '$1' ;/g;

# two-chars difinition
s/: (.)(.) ;/: '$1' '$2' ;/g;

# every all-CAP words should be split into quoted characters
# else all should have leading “_”
s/[a-z]\w+/_$&/g;
s!\b[A-Z_]{2,}!join " ", split //, $&!ge;
# except for "SQL"
s/sQL/SQL/g;

# every single character should be quoted
s/(\s)([^][ :|{}])(?= )/$1'$2'/g;
s/'''/'\\''/;

# rules to skip
s/\n_(space|nonquote_Character|identifier_Start)\s*:.*?;\n//gs;

print;
