use DBI;

my $db_instance = $ENV{DBI_SVR};
my $db_name     = "tempdb";
my $db_user     = $ENV{DBI_USER};
my $db_pass     = $ENV{DBI_PASS};

# lists all the drivers and data sources for each driver on your system
# http://oreilly.com/catalog/perldbi/chapter/ch04.html

### Probe DBI for the installed drivers
my @drivers = DBI->available_drivers();

die "No drivers found!\n" unless @drivers; # should never happen

### Iterate through the drivers and list the data sources for each one
foreach my $driver ( @drivers ) {
    print "Driver: $driver\n";
    next if $driver eq 'Proxy';
    my @dataSources = DBI->data_sources( $driver );
    foreach my $dataSource ( @dataSources ) {
        print "\tData Source is $dataSource\n";
    }
    print "\n";
}

# Right out of box of Strawberry Perl:
#
# Driver: ADO
#
# Driver: DBM
#         Data Source is DBI:DBM:f_dir=.
#
# Driver: ExampleP
#         Data Source is dbi:ExampleP:dir=.
#
# Driver: File
#         Data Source is DBI:File:f_dir=.
#
# Driver: Gofer
#
# Driver: ODBC
#         Data Source is dbi:ODBC:dBASE Files
#         Data Source is dbi:ODBC:Excel Files
#         Data Source is dbi:ODBC:MS Access Database
#
# Driver: Pg
#
# Driver: SQLite
#
# Driver: Sponge
#
# Driver: mysql

my $n = 0;
for my $cs (
     "DBI:ODBC:Driver={SQL Server};Server=$db_instance;Database=$db_name;UID=$db_user;PWD=$db_pass"
  , "DBI:ODBC:Driver={SQL Server};Server=$db_instance;Database=$db_name;Trusted_Connection=yes;"
  , "DBI:ODBC:Driver={SQL Server};Server=$db_instance;Database=$db_name;"
  , "DBI:ODBC:Driver={SQL Server};Server=(local);Database=$db_name;"
  ,  'DBI:ODBC:Provider=Microsoft.ACE.OLEDB.12.0;Data Source=D:\Projects\DBs\Northwind.accdb;User Id=;Password='
#    , "DBI:ODBC:DSN=AdvWork;"
#    , "DBI:ADO:DSN=AdvWork;"
  , "DBI:ADO:Provider=SQLOLEDB;Data Source=(local);Initial Catalog=tempdb;Integrated Security=SSPI"
  , "DBI:ADO:Provider=SQLOLEDB;Data Source=$db_instance;Initial Catalog=$db_name;Integrated Security=SSPI"
  , "DBI:ADO:Provider=SQLOLEDB.1;Persist Security Info=False;User ID=$db_user;PWD=$db_pass;Initial Catalog=$db_name;Data Source=$db_instance;"
  , "DBI:ADO:Provider=SQLNCLI10.1;Integrated Security=SSPI;Persist Security Info=False;User ID=\"\";Initial Catalog=$db_name;Data Source=$db_instance;Initial File Name=\"\";Server SPN=\"\";"
  , "DBI:ADO:Provider=SQLNCLI.1;Integrated Security=SSPI;Persist Security Info=False;Initial Catalog=$db_name;Data Source=$db_instance;"
) {
   my $dbh = DBI->connect($cs) or next;
   my $wtf = $cs;
   $wtf =~ s/$db_pass/secret/;
   printf "%2d CS: '%s' OK\n", ++$n, $wtf;
   $dbh->disconnect();
}
