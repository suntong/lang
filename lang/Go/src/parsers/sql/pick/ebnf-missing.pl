#! /usr/bin/env perl
# -*- perl -*- 

############################################################################
## Porgram: ebnf-missing.pl
## Purpose: pick missing ids from ebnf grammar
##	    ids that only show up on RHS but not on LHS
## Authors: Tong Sun (c) 2022, All rights reserved
############################################################################

=pod

panic: Unknown production: _implementation_defined_Universal_Character_Form_of_use_Name

goroutine 1 [running]:
github.com/goccmack/gocc/internal/ast.(*LexPart).Production(0xc00014a480, {0xc0000202c0, 0x3c})
        /path/to/Go/pkg/mod/github.com/goccmack/gocc@v0.0.0-20211213154817-7ea699349eca/internal/ast/lexpart.go:102 +0xdd
github.com/goccmack/gocc/internal/lexer/items.NewItem({0xc0000202c0, 0x5c}, 0xc0000202c0, 0xc000030360)
        /path/to/Go/pkg/mod/github.com/goccmack/gocc@v0.0.0-20211213154817-7ea699349eca/internal/lexer/items/item.go:47 +0x57
github.com/goccmack/gocc/internal/lexer/items.ItemList.Closure({0xc194e7f000, 0x5c, 0x80}, 0x0, 0xc000030360)
        /path/to/Go/pkg/mod/github.com/goccmack/gocc@v0.0.0-20211213154817-7ea699349eca/internal/lexer/items/itemlist.go:60 +0x165
github.com/goccmack/gocc/internal/lexer/items.(*ItemSet).Next(0xc14f4dfb20, {0x94e7e400, 0xc1})
        /path/to/Go/pkg/mod/github.com/goccmack/gocc@v0.0.0-20211213154817-7ea699349eca/internal/lexer/items/itemset.go:136 +0xed
github.com/goccmack/gocc/internal/lexer/items.(*ItemSets).Closure(0xc0001b11a0)
        /path/to/Go/pkg/mod/github.com/goccmack/gocc@v0.0.0-20211213154817-7ea699349eca/internal/lexer/items/itemsets.go:55 +0x170
github.com/goccmack/gocc/internal/lexer/items.GetItemSets(0xc00014a480)
        /path/to/Go/pkg/mod/github.com/goccmack/gocc@v0.0.0-20211213154817-7ea699349eca/internal/lexer/items/itemsets.go:40 +0xdb
main.main()
        /path/to/Go/pkg/mod/github.com/goccmack/gocc@v0.0.0-20211213154817-7ea699349eca/main.go:85 +0x5e5

real    2186m58.446s
user    2412m45.462s
sys     5m40.899s


2187/60/24 = 1.51875 -- gocc ran for one and half days to pick out
_implementation_defined_Universal_Character_Form_of_use_Name.
See https://github.com/goccmack/gocc/issues/132.



whereas using ebnf-missing.pl:

$ time ebnf-missing.pl < sqldef.bnf
] _implementation_defined_Universal_Character_Form_of_use_Name (1)
row_Value_Constructor_List
qualified_Join
_user_defined_Character_Repertoire_Name
. . .
real    0m0.005s
user    0m0.005s
sys     0m0.000s


=cut

# Read whole STDIN into variable bnf defs
my $defs = do { local $/; <> };

# remove comments
$defs =~ s{/\*.*?\*/}{\n}gs;
$defs =~ s{\n\s*\/\/.*?\n}{\n}gs;

# collections of LHS & RHS
my %clhs;
my %crhs;

sub collect {
    my($lhs, $rhs) = @_;
    # collect RHS to %crhs
    $rhs =~ s/\w+/
        $crhs{$&} = 1;
        "$&"
        /gex;
    # and LHS to %clhs
    $clhs{$lhs} = 1;
    return "\n$lhs : $rhs"
}

# split ebnf rule into two parts, LHS & RHS, then collect them individually
$defs =~ s/(?:\n|^)(\w+)\s*:\s*(.*?;)/
    collect($1, $2)
/gsex;

# compute the difference of two arrays/sets
# https://stackoverflow.com/questions/2933347/difference-of-two-arrays-using-perl

my %count;
for my $element (keys %clhs, keys %crhs) { $count{$element}++ }

my ( @union, @intersection, @difference );
for my $element (keys %count) {
    print STDERR "] $element ($count{$element})\n" if $element eq '_implementation_defined_Universal_Character_Form_of_use_Name';
    # push @union, $element;
    # push @{ $count{$element} > 1 ? \@intersection : \@difference }, $element;
    push @difference, $element unless $count{$element} > 1;
}

for my $element (@difference) {
    next if length $element == 1;
    print "$element\n";
}
