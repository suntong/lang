#! /usr/bin/env perl
# -*- perl -*- 

############################################################################
## Porgram: pickr.pl
## Purpose: pick selected recursively
## Authors: Tong Sun (c) 2022, All rights reserved
############################################################################


my @picked = qw( schema_definition table_definition view_definition );
# override from $ENV{"SEL"} if defined
@picked = split " ", $ENV{"SEL"} if $ENV{"SEL"};

# Read whole STDIN into variable bnf defs
my $defs = do { local $/; <> };

my %wanted;
while ($defs =~ m{(\w+)\s*:(.*?);}gsx) {
    #print "$&\n^^ ($1)\n\n";
    #print "$1\n" if (grep $1 eq $_, @picked);
    if ((grep $1 eq $_, @picked) || exists $wanted{$1}) {
	$wanted{$1} = 1;
	# put every new def from RHS into %wanted hash
	my $has_quantifiers = 0;
	for my $nd (split ' ', $2) {
	    if ($nd =~ s/[*?+]$//) {
		$has_quantifiers = 1;
	    }
	    #print STDERR "] $nd\n"; #if $nd =~ m/\s+/;
	    $wanted{$nd} = 0 unless exists $wanted{$nd};
	}
	# output the definition
	if ($has_quantifiers) { print STDERR "$&\n\n" }
	else { print "$&\n\n" }
    }
}

my $missedFN = "/tmp/missed.kw";
open(my $fh, ">", $missedFN) 
    or die "cannot open > $missedFN: $!";

while ( my ($key, $value) = each(%wanted) ) {
    #print STDERR "]] $key\n" if $key =~ m/\s+/;
    #print STDERR "]] $key => $value\n";
    #print {$fh} "] $key\n";
    print {$fh} "$key\n" if !$value;
}

