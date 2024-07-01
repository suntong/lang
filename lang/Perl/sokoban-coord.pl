#!/usr/bin/perl -l
use strict;
use warnings;

my $letter = 'A';

sub sokoban_puzzle_to_coordinate {
    # http://sokobano.de/wiki/index.php?title=Level_format
    my ($str) = @_;

    $str =~ s{([ \$\@\*\.\+\-]+)
        }{
        my $replacement = join '', map { 
            if ($letter++ eq "AA") { $letter = "a"; }
            $letter;
            } 1 .. length $&;
        "$replacement";
    }gex;

    return $str;
}

while (<STDIN>) {
    printf sokoban_puzzle_to_coordinate($_);
}
