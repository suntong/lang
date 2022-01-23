#! /usr/bin/env perl
# -*- perl -*- 

############################################################################
## Porgram: pick1.pl
## Purpose: pick single selected bnf def without recurs
## Authors: Tong Sun (c) 2022, All rights reserved
############################################################################


my @picked = qw( schema_definition table_definition view_definition );
# override from $ENV{"SEL"} if defined
@picked = split " ", $ENV{"SEL"} if $ENV{"SEL"};

# Read whole STDIN into variable bnf defs
my $defs = do { local $/; <> };

while ($defs =~ m{(\w+)\s*:(.*?);}gsx) {
    #print "$&\n^^ ($1)\n\n";
    #print "$1\n" if (grep $1 eq $_, @picked);
    if ((grep $1 eq $_, @picked)) {
	# check every new def from RHS
	my $has_quantifiers = 0;
	for my $nd (split ' ', $2) {
	    if ($nd =~ s/[*?+]$//) {
		$has_quantifiers = 1;
	    }
	}
	# output the definition
	if ($has_quantifiers) { print STDERR "$&\n\n" }
	else { print "$&\n\n" }
    }
}
