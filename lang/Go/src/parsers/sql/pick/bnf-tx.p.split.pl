#!/usr/bin/perl -0777n


my %regDefIds;
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

# split each *e*bnf rule into two parts, lexical & syntax parts
# to STDOUT and STDERR respectively
s/(?:\n|^)(\w+)\s*:\s*(.*?;)/
    print STDERR "$1 : ${\( lcfirst($1) )} ;\n\n";
    cookit($1, $2)
/gsex;

# rename all regDefIds on LHS as such as well
my $re = join '|', keys %regDefIds;
#print STDERR "] $re";
s/(?:\n|^)($re)/_$1/g;
# echo "foo, bar, AVG, Hex_String_Literal, Time_String, MODULE" | perl -ne 'my %map =("Hex_String_Literal"=>1, "Time_String"=>1,); my $re = join "|", keys %map; s/($re)/_$1/g; print'

# == Transforms from the original bnf-tx.p.pl


# fix "Non-xxx" as "Non_xxx"
s/(Non|Form|of)-/$1_/g;

# every all-CAP words should be split into quoted characters
s!\b[A-Z_]{2,}!join " ", split //, $&!ge;

# every single character should be quoted
s/(\s)([^][ ;:|{}])(?=( |\n))/$1'$2'/g;


# Repeats
# 1. { }... is redundant as { } itself means repeating. Thus need to remove
# 2. word... need to be changed as word {word}
s/(\{.*?\})\s*\.\.\./$1/gs;
s/(\S+)\s*\.\.\./$1 {$1}/gs;


print;
