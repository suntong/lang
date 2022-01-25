#!/usr/bin/perl -0777n

sub cookit {
    my($lhs, $rhs) = @_;
    # all terms on RHS should be regDefIds
    $rhs =~ s/([A-Z][a-z]\w+)/"_".lcfirst($&)/ge;
    return "\n${\( lcfirst($lhs) )} : $rhs"
}

# split each *e*bnf rule into two parts, lexical & syntax parts
# to STDOUT and STDERR respectively
s/(?:\n|^)(\w+)\s*:\s*(.*?;)/
    print STDERR "$1 : ${\( lcfirst($1) )} ;\n\n";
    cookit($1, $2)
/gsex;

print;
