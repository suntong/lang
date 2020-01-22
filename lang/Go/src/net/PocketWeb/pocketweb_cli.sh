templateFile=$GOPATH/src/github.com/go-easygen/easygen/test/commandlineFlag
[ -s $templateFile.tmpl ] || templateFile=/usr/share/doc/easygen/examples/commandlineFlag
[ -s $templateFile.tmpl ] || {
  echo No template file found
  exit 1
}

easygen $templateFile pocketweb_cli | gofmt > pocketweb_cli.go
