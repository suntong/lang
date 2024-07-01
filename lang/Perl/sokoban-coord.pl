#!/usr/bin/perl -l
use strict;
use warnings;
use open ':std', ':encoding(UTF-8)';

my $letter = 'A';
my $fill = @ARGV >= 1;

sub AZaz {
    my ($letter) = @_;
    if ($letter eq "AA") { $letter = "a"; }
    $letter
}

sub sokoban_puzzle_to_coordinate {
    # http://sokobano.de/wiki/index.php?title=Level_format
    my ($str) = @_;

    if ($fill) {
        $str =~ s{([ \$\@\*\.\+\-]+)
            }{
            my $replacement = join '', map {
                AZaz($letter++);
                } 1 .. length $&;
            "$replacement";
        }gex;
    }

    $str =~ s{.}{
        chr(ord($&) + 0xfee0);
        }ge;
    return $str;
}

while (<STDIN>) {
    printf sokoban_puzzle_to_coordinate($_);
}
