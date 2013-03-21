#!/usr/bin/env perl
#
# Note: t/test.t searches for the next line.
# Annotation: Demonstrates graphing a database schema.

use strict;
use warnings;

use DBI;

use GraphViz2;
use GraphViz2::DBI;

use Log::Handler;

# ---------------

exit 0 if (! $ENV{DBI_DSN});

my($logger) = Log::Handler -> new;

$logger -> add
	(
	 screen =>
	 {
		 maxlevel       => 'debug',
		 message_layout => '%m',
		 minlevel       => 'error',
	 }
	);

my($graph) = GraphViz2 -> new
	(
	 edge   => {color => 'grey'},
	 global => {directed => 1},
	 graph  => {rankdir => 'TB'},
	 logger => $logger,
	 node   => {color => 'blue', shape => 'oval'},
	);
my($attr)              = {};
$$attr{sqlite_unicode} = 1 if ($ENV{DBI_DSN} =~ /SQLite/i);
my($dbh)               = DBI -> connect($ENV{DBI_DSN}, $ENV{DBI_USER}, $ENV{DBI_PASS}, $attr);

$dbh -> do('PRAGMA foreign_keys = ON')           if ($ENV{DBI_DSN} =~ /SQLite/i);
$dbh -> do("set search_path = $ENV{DBI_SCHEMA}") if ($ENV{DBI_SCHEMA});

my($g) = GraphViz2::DBI -> new(dbh => $dbh, graph => $graph);

$g -> create(name => '');

my($format)      = shift || 'svg';
my($output_file) = shift || File::Spec -> catfile('html', "dbi.schema.$format");

$graph -> run(format => $format, output_file => $output_file);
