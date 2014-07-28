#################################################################
## Compare SQL Server Table schemas, Func Def
## http://powershell.org/wp/2013/04/28/comparing-sql-server-table-schemas-with-powershell/
## Enrique Puig Nouselles
## Epuig1984@gmail.com
#################################################################

function Compare-SQLServerTables
(
    [string]$srv1,[string]$bd1,[string]$sch1,[string]$TableName1,
    [string]$srv2,[string]$bd2,[string]$sch2,[string]$TableName2
)
{
        [reflection.assembly]::LoadWithPartialName("Microsoft.SqlServer.Smo") | Out-Null

        $S1=New-Object "Microsoft.SqlServer.Management.Smo.Server" $srv1
        if($S1 -ne $null)
        { 
            if($S1.databases[$bd1] -ne $null)
            {
                $tab1=$S1.databases[$bd1].Tables[$TableName1]
                
                $res=$tab1 | Where-Object {$_.Schema -eq $sch1}

                if($res.Count -eq 0)
                {
                    throw "Error: The schema $sch1 doesn't contain any table called $TableName1."
                }
            }
            else
            {
                throw "Error: The database '$bd1' doesn't exist."
            }
        }
        else
        {
          throw "Error: We couldn't connect to the server '$srv1'. Please check your credentials and the servername"
        }


        $S2=New-Object "Microsoft.SqlServer.Management.Smo.Server" $srv2
        if($S2 -ne $null)
        { 
            if($S2.databases[$bd2] -ne $null)
            {
                $tab2=$S2.databases[$bd2].Tables[$TableName2]

                $res=$tab2 | Where-Object {$_.Schema -eq $sch2}

                if($res.Count -eq 0)
                {
                    throw "Error: The schema $sch2 doesn't contain any table called $TableName2."
                }
            }
            else
            {
                throw "Error: The database '$bd2' doesn't exist."
            }
        }
        else
        {
           throw "Error: We couldn't connect to the server '$srv1'. Please check your credentials and the servername"
        }


        ##check columns
        $ncols1=$tab1.Columns.Count
        $ncols2=$tab2.Columns.Count
        $eqCols=$true
        $eqChecks=$true
        $eqIndexes=$true
        $resultCompare=$true

        if($ncols1 -ne $ncols2)
        {
          return $false;
        }

        [Array]$colList=@()

        ##check data types, nullable columns,computed columns, identity columns, persisted columns, cols with default
        ## rimary keys and foreign keys
        $tab1.Columns | ForEach-Object{
        
            $c1=$_
            $aux=$tab2.Columns | Where-Object {
                $_.Name -eq $c1.Name -and $c1.DataType -eq $_.DataType -and $c1.Nullable -eq $_.Nullable -and $c1.Identity -eq $_.Identity -and $c1.IdentitySeed -eq $_.IdentitySeed -and $c1.Computed -eq $_.Computed -and $c1.ComputedText -eq $_.ComputedText -and $c1.DefaultConstraint.Text -eq $_.DefaultConstraint.Text -and $c1.InPrimaryKey -eq $_.InPrimaryKey -and $c1.IsPersisted -eq $_.IsPersisted -and $c1.IsForeignKey -eq $_.IsForeignKey 
            }
            if($aux -eq $null)
            {
                $eqCols=$false
                return;
            }
        }
        #check the other way to make sure that are completely equal tables
        if($eqCols)
        {
            $tab2.Columns | ForEach-Object{
                $c1=$_
                $aux=$tab1.Columns | Where-Object {
                    $_.Name -eq $c1.Name -and $c1.DataType -eq $_.DataType -and $c1.Nullable -eq $_.Nullable -and $c1.Identity -eq $_.Identity -and $c1.IdentitySeed -eq $_.IdentitySeed -and $c1.Computed -eq $_.Computed -and $c1.ComputedText -eq $_.ComputedText -and $c1.DefaultConstraint.Text -eq $_.DefaultConstraint.Text -and $c1.InPrimaryKey -eq $_.InPrimaryKey -and $c1.IsPersisted -eq $_.IsPersisted
                }
                if($aux -eq $null)
                {
                    $eqCols=$false
                    return;
                }
            }
        }

        ##check constraints
        ##we cannot create 2 constraints with the same name at the same database
        $tab1.Checks | ForEach-Object{
            $tab2.Columns | ForEach-Object{

                $chk1=$_
                $checks=$tab2.Checks | Where-Object { $chk1.Text -eq $_.Text -and $chk1.IsEnabled -eq $_.IsEnabled}
                if($checks -eq $null -or $checks.Count -eq 0)
                {                
                    $eqChecks=$false
                    return;
                }
            }
        }

        ##check it out in the other way
        if($eqChecks)
        {
            $tab2.Checks | ForEach-Object{
            Write-Host "hola que ase"
                $chk1=$_
                $checks=$tab1.Checks | Where-Object { $chk1.Text -eq $_.Text -and $chk1.IsEnabled -eq $_.IsEnabled}
                if($checks -eq $null -or $checks.Count -eq 0)
                {
                    $eqChecks=$false
                    return;
                }
            }
        }

        ##Indexes section
        [Array]$indexes1=@()
        [Array]$cols=@()

        ##check indexes
        $tab1.Indexes | ForEach-Object{
            $ix1=$_
            #check index type and properties
            $ix=$tab2.Indexes | Where-Object{
                $ix1.IsClustered -eq $_.IsClustered -and $ix1.HasFilter -eq $_.HasFilter -and $ix1.IgnoreDuplicateKeys -eq $_.IgnoreDuplicateKeys -and $ix1.IndexedColumns.Count -eq $_.IndexedColumns.Count -and  $ix1.IsIndexOnComputed -eq  $_.IsIndexOnComputed -and $ix1.IsPartitioned -eq $_.IsPartitioned -and $ix1.IsSpatialIndex -eq $_.IsSpatialIndex -and $ix1.IsUnique -eq $_.IsUnique -and $ix1.IsXmlIndex -eq $_.IsXmlIndex 
            }
            if( $ix -eq $null -or $ix.Count -eq 0)
            {
                $eqIndexes=$false
                return;
            }
            else
            {
                ##check index column names
                $ix1.IndexedColumns | ForEach-Object{
            
                    $col1=$_
            
                    #Get all indexed columns
                    $cols= $ix.IndexedColumns | Where-Object{
                
                        $col1.Name -eq $_.Name
                    }

                    if($cols -eq $null -or $cols.Count -eq 0)
                    {
                        $eqIndexes=$false
                        return;
                    }
                }
            }

        }

        if($eqIndexes)
        {
            $tab2.Indexes | ForEach-Object{
    
                $ix1=$_

                #check index type and properties
                $ix=$tab1.Indexes | Where-Object{
                    $ix1.IsClustered -eq $_.IsClustered -and $ix1.HasFilter -eq $_.HasFilter -and $ix1.IgnoreDuplicateKeys -eq $_.IgnoreDuplicateKeys -and $ix1.IndexedColumns.Count -eq $_.IndexedColumns.Count -and  $ix1.IsIndexOnComputed -eq  $_.IsIndexOnComputed -and $ix1.IsPartitioned -eq $_.IsPartitioned -and $ix1.IsSpatialIndex -eq $_.IsSpatialIndex -and $ix1.IsUnique -eq $_.IsUnique -and $ix1.IsXmlIndex -eq $_.IsXmlIndex 
                }

                if( $ix -eq $null -or $ix.Count -eq 0)
                {
       
                    $eqIndexes=$false
                    return;
                }
                else
                {
                    ##check index column names
                    $ix1.IndexedColumns | ForEach-Object{
            
                        $col1=$_
            
                        $cols= $ix.IndexedColumns | Where-Object{
                
                            $col1.Name -eq $_.Name
                        }
                        if($cols -eq $null -or $cols.Count -eq 0)
                        {
                            $eqIndexes=$false
                            return;
                        }
                    }
                }

            }
        }
        if($eqCols -eq $false -or $eqChecks -eq $false -or $eqIndexes -eq $false)
        {
            $resultCompare=$false
        }
        return $resultCompare
} 