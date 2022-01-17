
### Synopsis

Check out [`easygen` cli code-gen example for go-flags](https://github.com/go-easygen/easygen/issues/46) on how these Go code were ***automatically** generated*.

### Demo

``` sh
$ wireframed
Please specify one command of: build, install or publish

Usage:
  wireframed [OPTIONS] <build | install | publish>

Application Options:
  -H, --host=    host address (default: localhost) [$REDO_HOST]
  -p, --port=    listening port (default: 80) [$REDO_PORT]
  -f, --force    force start [$REDO_FORCE]
  -v, --verbose  Verbose mode (Multiple -v options increase the verbosity)

Help Options:
  -h, --help     Show this help message

Available commands:
  build    Build the network application
  install  Install the network application
  publish  Publish the network application



$ wireframed publish
the required flag `-d, --dir' was not specified

Usage:
  wireframed [OPTIONS] publish [publish-OPTIONS] ID Num Rest...

Publish the built network application to central repo

Application Options:
  -H, --host=       host address (default: localhost) [$REDO_HOST]
  -p, --port=       listening port (default: 80) [$REDO_PORT]
  -f, --force       force start [$REDO_FORCE]
  -v, --verbose     Verbose mode (Multiple -v options increase the verbosity)

Help Options:
  -h, --help        Show this help message

[publish command options]
      -d, --dir=    publish dir
          --suffix= source file suffix (default: .go,.c,.s)
      -o, --out=    output filename



$ wireframed publish -d .
the required arguments `ID` and `Num` were not provided

Usage:
  wireframed [OPTIONS] publish [publish-OPTIONS] ID Num Rest...

Publish the built network application to central repo

Application Options:
  -H, --host=       host address (default: localhost) [$REDO_HOST]
  -p, --port=       listening port (default: 80) [$REDO_PORT]
  -f, --force       force start [$REDO_FORCE]
  -v, --verbose     Verbose mode (Multiple -v options increase the verbosity)

Help Options:
  -h, --help        Show this help message

[publish command options]
      -d, --dir=    publish dir
          --suffix= source file suffix (default: .go,.c,.s)
      -o, --out=    output filename



$ wireframed publish -d ../dist v1 123 abc def
Publish the network application
Copyright (C) 2022, Myself <me@mine.org>

Doing Publish, with []string{}
../dist .go,.c,.s  {v1 123 [abc def]}

```

The last two lines are output from:


``` go
	fmt.Printf("Doing Publish, with %#v\n", args)
	fmt.Println(x.Dir, x.Suffix, x.Out, x.Args)
```


``` sh
$ wireframed install --dir ../new abc def
Install the network application
Copyright (C) 2022, Myself <me@mine.org>

Doing Install, with []string{"abc", "def"}
../new .go,.c,.s
```

### Update

After [ouput global option](https://github.com/suntong/lang/commit/5ab110dfde2c2a481e14f8cb8afa4d4b3cb4bd23#diff-4040661b6ba228abb2484408a7a62935f7660d89545b11fc65567f54d36d7501),

``` sh
$ wireframed publish -d ../dist v1 123 abc def
Publish the network application
Copyright (C) 2022, Myself <me@mine.org>

Doing Publish, with {Host:localhost Port:80 Force:false Verbflg:0x4a8120 Verbose:0}, []
../dist .go,.c,.s  {v1 123 [abc def]}



$ REDO_HOST=myserver REDO_PORT=8080 wireframed publish -v -d ../dist v1 123 abc def
Publish the network application
Copyright (C) 2022, Myself <me@mine.org>

Doing Publish, with {Host:myserver Port:8080 Force:false Verbflg:0x4a8120 Verbose:1}, []
../dist .go,.c,.s  {v1 123 [abc def]}



$ REDO_HOST=myserver REDO_PORT=8080 wireframed publish -vv -H newserver --port 8888 --force -d ../dist v1 123 abc def
Publish the network application
Copyright (C) 2022, Myself <me@mine.org>

Doing Publish, with {Host:newserver Port:8888 Force:true Verbflg:0x4a8120 Verbose:2}, []
../dist .go,.c,.s  {v1 123 [abc def]}
```
