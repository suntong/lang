curl -sL 'https://github.com/ronsavage/SQL/raw/master/sql-2016.ebnf' | tee sql-2016.bnf | perl -0777 -wnE's{ (?: /\* .*? \*/ ) }{\n}gsx; for (split /\n\n+/) { s/::=\s*/\n  : /gs; s/\s*$/\n  ;\n/s; say } ;' > sql-2016.ebnf

# do my picks
SEL='schema_definition' pickr.pl < sql-2016.ebnf > sql-2016.definitions.ebnf 2> sql-2016.definitions+.ebnf
SEL=`cat /tmp/missed.kw` pick1.pl < sql-2016.ebnf > sql-2016.missed.ebnf 2> sql-2016.missed+.ebnf

