 pwd=`pwd`
 cd $GOPATH/src/github.com/go-easygen/easygen/test
 easygen commandlineGoFlags.header,commandlineGoFlags.ityped.tmpl,commandlineGoFlags "$pwd/demo"_cli | gofmt > "$pwd/demo"_cliDef.go
