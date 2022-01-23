#! /usr/bin/env perl
# -*- perl -*- 

############################################################################
## Porgram: bnf-pickr.pl
## Purpose: pick selected from bnf grammar recursively
## Authors: Tong Sun (c) 2022, All rights reserved
############################################################################

=pod

The most important difference between ebnf and bnf grammar is that

- ebnf does the defines top down whereas
- bnf does the defines bottom up

which means that 

- ebnf can basically be scanned in one pass whereas
- bnf need to be scanned over and over until all missed are gone

=cut

my @picked = qw( Schema_Definition Table_Definition View_Definition );
# override from $ENV{"SEL"} if defined
@picked = split " ", $ENV{"SEL"} if $ENV{"SEL"};

# Read whole STDIN into variable bnf defs
my $defs = do { local $/; <> };

my %wanted;
map { $wanted{$_} = 1 } @picked;

while ($defs =~ m{(\w+)\s*:(.*?);}gsx) {
    #print "$&\n^^ ($1)\n\n";
    #print "--> $1\n" if (grep $1 eq $_, @picked);
    if (exists $wanted{$1}) {
	# put every new def from RHS into %wanted hash
	my $has_quantifiers = 0;
	for my $nd (split ' ', $2) {
	    if ($nd =~ s/[*?+]$//) {
		$has_quantifiers = 1;
	    }
	    #print STDERR "] $nd\n"; #if $nd =~ m/\s+/;
	    $wanted{$nd} = 1 unless exists $wanted{$nd};
	}
	# output the definition
	if ($has_quantifiers) { print STDERR "$&\n\n" }
	else { print "$&\n\n" }
	# mark this rule as done
	$wanted{$1} = 0;
    }
}

my $missedFN = "/tmp/missed.kw";
open(my $fh, ">", $missedFN) 
    or die "cannot open > $missedFN: $!";

while ( my ($key, $value) = each(%wanted) ) {
    #print STDERR "]] $key\n" if $key =~ m/\s+/;
    #print STDERR "]] $key => $value\n";
    #print {$fh} "] $key\n";
    print {$fh} "$key\n" if $value;
}
