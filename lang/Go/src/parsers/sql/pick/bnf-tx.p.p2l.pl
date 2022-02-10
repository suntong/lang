#!/usr/bin/perl -0777n

# IDs missing from sql-92.p.definitions+.l.ebnf
my @CvtIds = qw(Space Separator Escape_Character Nonquote_Character Numeric_Primary Character_Primary Datetime_Primary Boolean_Primary Value_Expression_Primary Day_time_Literal Year_month_Literal Parameter_Name Embedded_Variable_Name Correlation_Name Implementation_defined_Universal_Character_Form_of_use_Name Standard_Universal_Character_Form_of_use_Name User_defined_Character_Repertoire_Name Implementation_defined_Character_Repertoire_Name Qualifier Set_Function_Type Trim_Source Trim_Character Trim_Specification Row_Value_Constructor Value_Expression Character_Value_Expression String_Value_Expression Bit_Value_Expression Query_Expression Match_Value In_Predicate_Value Truth_Value Joined_Table Derived_Table Table_Name Collate_Clause Searched_When_Clause Simple_When_Clause Else_Clause Schema_Name_Clause Levels_Clause Schema_Element Outer_Join_Type Join_Specification Table_Subquery Derived_Column_List View_Column_List Case_Operand Pattern Introducercharacter_Set_Specification Schema_Character_Set_Specification Table_Constraint_Definition Column_Constraint_Definition Identifier_Start Start_Position Delimited_Identifier_Part Time_Zone Non_second_Datetime_Field Set_Quantifier String_Length);

# init hash from an array
my %regDefIds = map { $_ => 1 } @CvtIds;
my $re = join '|', keys %regDefIds;

sub cookit {
    my($lhs, $rhs) = @_;
    # all terms on RHS should be regDefIds
    $rhs =~ s/([A-Z][a-z]\w+)/
        my $wants = lcfirst($&);
        $regDefIds{$wants} = 1;
        "_$wants"
        /gex;
    return "\n${\( lcfirst($lhs) )} : $rhs"
}

s/(?:\n|^)($re)\s*:\s*(.*?;)/
    cookit($1, $2)
/gsex;

# rename init %regDefIds to lcfirst first
%regDefIds = map { lcfirst($_) => 1 } keys %regDefIds;

# rename all (init & added) regDefIds on LHS as such as well
$re = join '|', keys %regDefIds;
#print STDERR "] $re";
s/(?:\n|^)($re)/\n_$1/g;

print;
