curl -sL 'https://github.com/ronsavage/SQL/raw/master/sql-2016.ebnf' | tee sql-2016.bnf | perl -0777 -wnE's{ (?: /\* .*? \*/ ) }{\n}gsx; for (split /\n\n+/) { s/::=\s*/\n  : /gs; s/\s*$/\n  ;\n/s; say } ;' > sql-2016.ebnf
curl -sL 'https://github.com/ronsavage/SQL/raw/master/sql-2016.ebnf' > sql-2016.bnf

# do my picks for sql-2016
SEL='schema_definition' ebnf-pickr.pl < sql-2016.ebnf > sql-2016.definitions.ebnf 2> sql-2016.definitions+.ebnf
SEL=`cat /tmp/missed.kw` ebnf-pick1.pl < sql-2016.ebnf > sql-2016.missed.ebnf 2> sql-2016.missed+.ebnf

# for sql-92
sed -n '/h2 SQL Module/,$p' sql-92.bnf | perl -00ne 's|(\n?)--|$1// |g; s{<(.*?)>}{join "_", map { ucfirst($_) } split / /, $1}eg; s/::=/:/; s/\n$/ ;\n/; print' > sql-92.p.ebnf
bnf-pickr.pl < sql-92.p.ebnf > sql-92.p.definitions.ebnf

sed -n '1,/h2 Constraints/p' sql-92.bnf | perl -00ne 's|(\n?)--|$1// |g; s{<(.*?)>}{lcfirst(join "_", map { ucfirst($_) } split / /, $1)}eg; s/::=/:/; s/\n$/ ;\n/; print' > sql-92.l.ebnf
SEL='equals_Operator unsigned_Numeric_Literal greater_Than_Or_Equals_Operator table_Element_List colonHost_Identifier non-join_Query_Expression domain_Name asterisk concatenation_Operator comma searched_When_Clause... left_Paren non-second_Datetime_Field less_Than_Or_Equals_Operator character_Set_Specification schema_Name interval_Qualifier general_Literal column_Name datetime_Value_Function minus_Sign data_Type sign less_Than_Operator authorization_Identifier not_Equals_Operator colon qualified_Local_Table_Name period plus_Sign schema_Element... simple_When_Clause... form-of-use_Conversion right_Paren identifier greater_Than_Operator solidus qualified_Name question_Mark' bnf-pickr.pl < sql-92.l.ebnf > sql-92.l.definitions.ebnf 2> sql-92.l.definitions+.ebnf

# for gocc
cat sql-92.l.definitions*.ebnf | bnf-tx.pl > sql-92.gocc.l.ebnf
cat sql-92.p.definitions+.ebnf | bnf-tx.p.split.pl 2> sql-92.p.definitions+.p.ebnf | bnf-tx.p.pl > sql-92.p.definitions+.l.ebnf
