
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

### Update 1

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

### Update 2

After turning `--suffix` to use choices.

``` sh
$ wireframed publish
the required flag `-d, --dir' was not specified

Usage:
  wireframed [OPTIONS] publish [publish-OPTIONS] ID Num Rest...

Publish the built network application to central repo

Application Options:
  -H, --host=                  host address (default: localhost) [$REDO_HOST]
  -p, --port=                  listening port (default: 80) [$REDO_PORT]
  -f, --force                  force start [$REDO_FORCE]
  -v, --verbose                Verbose mode (Multiple -v options increase the verbosity)

Help Options:
  -h, --help                   Show this help message

[publish command options]
      -d, --dir=               publish dir
      -s, --suffix=[.go|.c|.h] source file suffix for publish
      -o, --out=               output filename

# see that the --suffix= now use choices



$ wireframed publish -d ../dist v1 123 abc def
Publish the network application
Copyright (C) 2022, Myself <me@mine.org>

Doing Publish, with {Host:localhost Port:80 Force:false Verbflg:0x4a8180 Verbose:0}, []
../dist []  {v1 123 [abc def]}

# the --suffix= has empty choices by default



$ wireframed publish -d ../dist -s .c -s .h v1 123 abc def
Publish the network application
Copyright (C) 2022, Myself <me@mine.org>

Doing Publish, with {Host:localhost Port:80 Force:false Verbflg:0x4a8180 Verbose:0}, []
../dist [.c .h]  {v1 123 [abc def]}

# the --suffix= has two choices now



$ wireframed publish -d ../dist -s .x v1 123 abc def
Invalid value `.x' for option `-s, --suffix'. Allowed values are: .go, .c or .h

Usage:
  wireframed . . .

# when the --suffix= has been provided with wrong choice

```

### Update 3

After switching to `go-easygen` which has `clis` support functions,

``` sh
$ wireframed install -v
Install the network application
Copyright (C) 2022, Myself <me@mine.org>

[redo::install] Doing Install, with {Host:localhost Port:80 Force:false Verbflg:0x4a92e0 Verbose:1}, []

./ .go,.c,.s
[redo::install] Warning: Install, Exec, Sample warning: Instance not found
```

- The `clis.Verbose(1,` will only output if the Verbose level is >=1. 
- Removing the `-v` option from the command line the "Doing Install" line will disappear.
- With `clis.WarnOn` & `clis.AbortOn` reporting warning or critical errors (in color) will be a breeze.

### Update on 2022-01-22

After update to new template, which uses `-V, --version` to show program version,

``` sh
$ wireframed
Please specify one command of: build, install or publish

Usage:
  wireframed [OPTIONS] <build | install | publish>

Application Options:
  -H, --host=    Host address (default: localhost) [$REDO_HOST]
  -p, --port=    Listening port (default: 80) [$REDO_PORT]
  -f, --force    Force start [$REDO_FORCE]
  -v, --verbose  Verbose mode (Multiple -v options increase the verbosity)
  -V, --version  Show program version and exit

Help Options:
  -h, --help     Show this help message

Available commands:
  build    Build the network application
  install  Install the network application
  publish  Publish the network application

$ wireframed -V
redo - global option redo
Copyright (C) 2022, Myself <me@mine.org>

Redo global option via automatic code-gen

Built on 2022-01-22
Version 0.1.0
```

