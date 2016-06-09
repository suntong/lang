# 2016-05-31

## [001-hello.go](001-hello.go)

    $ go run 001-hello.go
    Hello, world! Your age is 100?
    
    $ go run 001-hello.go -name Clipher -a 77
    ERR! undefined flag -n
    
    $ go run 001-hello.go --name Clipher -a 77
    Hello, Clipher! Your age is 77?

## [002-flag.go](002-flag.go)
    
    $ go run 002-flag.go 
    port=0, x=false, y=false
    
    $ go run 002-flag.go -h 
    Options:
    
      -h, --help     display help information
      -p, --port     short and long format flags both are supported
      -x             boolean type
      -y             boolean type, too
    
    $ phd
    
    $ go run 002-flag.go 
    port=0, x=false, y=false
    
    $ go run 002-flag.go -h 
    Options:
    
      -h, --help     display help information
      -p, --port     short and long format flags both are supported
      -x             boolean type
      -y             boolean type, too
    
    $ go run 002-flag.go -p=8080 -x
    port=8080, x=true, y=false
    
    $ go run 002-flag.go -p 8080 -x=p=8080 -x
    port=8080, x=true, y=false
    
    $ go run 002-flag.go -p 8080 -x
    port=8080, x=true, y=false
    
    $ go run 002-flag.go -p 8080 -xy
    port=8080, x=true, y=true

## [003-required-flag.go](003-required-flag.go)

    $ go run 003-required-flag.go 
    ERR! required argument --id missing
    
    $ go run 003-required-flag.go --id=2
    2

## [004-default-flag.go](004-default-flag.go)

    $ go run 004-default-flag.go 
    2, /home/user, 1000, /home/user/dev
    
    $ go run 004-default-flag.go -h
    Options:
    
      -h, --help                       display help information
          --basic[=2]                  basic usage of default
          --env[=$HOME]                env variable as default
          --expr[=$BASE_PORT+1000]     expression as default
          --devdir[=$HOME/dev]         directory of developer
    
    $ BASE_PORT=8000 go run 004-default-flag.go --basic=3
    3, /home/user, 9000, /home/user/dev

## [005-slice.go](005-slice.go)

    $ go run 005-slice.go 
    {"Friends":null}
    
    $ go run 005-slice.go -FAlice -FBob -F Charlie
    {"Friends":["Alice","Bob","Charlie"]}

## [006-map.go](006-map.go)

    $ go run 006-map.go 
    {"Macros":null}
    
    $ go run 006-map.go -Dx=not-a-number
    ERR! `not-a-number` couldn't converted to an int value
    
    $ 
    
    $ go run 006-map.go -Dx=1 -D y=2
    {"Macros":{"x":1,"y":2}}

## [007-force-flag.go](007-force-flag.go)

    $ go run 007-force-flag.go 
    ERR! required argument -r missing
    
    $ go run 007-force-flag.go -v
    v0.0.1

## [008-child-command.go](008-child-command.go)

    $ go run 008-child-command.go 
    Hello, root command, I am 
    
    $ go run 008-child-command.go help
    this is root command
    
    Options:
    
      -h, --help     display help information
          --name     your name
    
    Commands:
      help    display help information
      child   this is a child command
    
    $ go run 008-child-command.go child
    Hello, child command, I am 
    
    $ go run 008-child-command.go help child
    this is a child command
    
    Options:
    
      -h, --help     display help information
          --name     your name
    
    $ go run 008-child-command.go child --name=123
    Hello, child command, I am 123
    
    $ go run 008-child-command.go childx
    ERR! command childx not found
    Did you mean child?
    exit status 1
    
    $ go run 008-child-command.go chd
    ERR! command chd not found
    Did you mean child?
    exit status 1

## [009-auto-helper.go](009-auto-helper.go)

    $ go run 009-auto-helper.go 
    
    $ go run 009-auto-helper.go -h
    Options:
    
      -h, --help     show help
    
## [010-validator.go](010-validator.go)
  
    $ go run 010-validator.go 
    {"Help":false,"Age":0,"Gender":"male"}
    
    $ go run 010-validator.go --age=-1
    ERR! age -1 out of range
    
    $ go run 010-validator.go --age=1000
    ERR! age 1000 out of range
    
    $ go run 010-validator.go -g balabala
    ERR! invalid gender balabala
    
    $ go run 010-validator.go --age 88 --gender female
    {"Help":false,"Age":88,"Gender":"female"}
    
## [011-prompt-and-password.go](011-prompt-and-password.go)
    
    $ go run 011-prompt-and-password.go 
    type github account: aa
    type the password: 
    username=aa, password=bb

## [012-decoder.go](012-decoder.go)

    $ go run 012-decoder.go 
    null
    
    $ go run 012-decoder.go -d a,b,c
    ["a","b","c"]
    
## [014-time-and-duration.go](014-time-and-duration.go)
    
    $ go run 014-time-and-duration.go 
    .../github.com/mkideal/cli/ext/decoders.go:12:2: cannot find package "github.com/jinzhu/now" in any of:
            /usr/lib/go-1.6/src/github.com/jinzhu/now (from $GOROOT)
            .../github.com/jinzhu/now (from $GOPATH)
    
    $ go get -v github.com/jinzhu/now

## [015-file.go](015-file.go)

    $ go run 015-file.go 
    .../github.com/mkideal/cli/ext/decoders.go:12:2: cannot find package "github.com/jinzhu/now" in any of:
            /usr/lib/go-1.6/src/github.com/jinzhu/now (from $GOROOT)
            .../github.com/jinzhu/now (from $GOPATH)

## [016-parser.go](016-parser.go)

    $ go run 016-parser.go 
    {
        "A": "",
        "B": 0,
        "C": false
    }
    
    $ go run 016-parser.go -c '{"A": "hello", "b": 22, "C": true}'
    {
        "A": "hello",
        "B": 22,
        "C": true
    }

## [017-jsonfile.go](017-jsonfile.go)

    $ go run 017-jsonfile.go 
    {
        "A": "",
        "B": 0,
        "C": false
    }
    
    $ go run 017-jsonfile.go echo '{"A": "hello", "b": 22, "C": true}' > test.json
    
    $ echo '{"A": "hello", "b": 22, "C": true}' > test.json
    
    $ go run 017-jsonfile.go -c test.json
    {
        "A": "hello",
        "B": 22,
        "C": true
    }
    
    $ rm test.json

## [018-custom-parser.go](018-custom-parser.go)

    $ go run 018-custom-parser.go 
    {0 }
    
    $ go run 018-custom-parser.go --cfg xxx
    {2 B}

## [019-hooks.go](019-hooks.go)

    $ go run 019-hooks.go 
    OnRootBefore invoked
    exec root command
    OnRootAfter invoked
    
    $ go run 019-hooks.go child1
    child1 OnBefore invoked
    OnRootBefore invoked
    exec child1 command
    OnRootAfter invoked
    child1 OnAfter invoked
    
    $ go run 019-hooks.go child2
    exec child2 command
    
    $ go run 019-hooks.go -e
    OnRootBefore invoked
    exec root command
    root command returns error
    exit status 1
    
    $ go run 019-hooks.go child1 -e
    child1 OnBefore invoked
    OnRootBefore invoked
    exec child1 command
    child1 command returns error
    exit status 1

## [020-daemon.go](020-daemon.go)

	$ go run 020-daemon.go & sleep 3; ps | grep daemon
	[2] 5777
	start ok
	5796 pts/6    00:00:00 020-daemon
	[2]+  Done                    go run 020-daemon.go

	$ ps | grep daemon | wc 
    0       0       0

## [021-editor.go](021-editor.go)

    $ go run 021-editor.go -m "hello, editor"
    msg: hello, editor
    
    $ go run 021-editor.go 
    msg: hello, editor2

## [022-custom-editor.go](022-custom-editor.go)

    $  go run 022-custom-editor.go -m "hello, editor"
    msg: hello, editor
    
    $ EDITOR=jed go run 022-custom-editor.go
    msg: from Custom Editor

## [027-global-option.go](027-global-option.go)

```sh
$ 027-global-option.sh
+ cat
+ echo ===run root command
+ go run 027-global-option.go -h
===run root command
application

Options:

  -h, --help              display help information
  -c, --config[=1.json]   config
  -H, --host[=$HOST]      host addr
  -p, --port              listening port

Commands:
  sub   subcommand

+ echo =case 1: read from file
+ go run 027-global-option.go
=case 1: read from file
root
{"Help":false,"Host":"127.0.0.1","Port":8080}{"Help":false,"Host":"127.0.0.1","Port":8080}

+ echo =case 2: read from ENV variables
+ HOST=10.0.0.1
+ go run 027-global-option.go
=case 2: read from ENV variables
root
{"Help":false,"Host":"10.0.0.1","Port":8080}{"Help":false,"Host":"10.0.0.1","Port":8080}

+ echo =case 3: read from command line
=case 3: read from command line
+ HOST=10.0.0.1
+ go run 027-global-option.go -H 168.0.0.1
root
{"Help":false,"Host":"168.0.0.1","Port":8080}{"Help":false,"Host":"168.0.0.1","Port":8080}

+ echo ===run sub command
+ go run 027-global-option.go sub -h
===run sub command
subcommand

Options:

  -h, --help              display help information
  -c, --config[=1.json]   config
  -H, --host[=$HOST]      host addr
  -p, --port              listening port
  -w                      world is a sub flag

+ echo =case 1: read from file
+ go run 027-global-option.go sub
=case 1: read from file
{"Help":false,"Host":"127.0.0.1","Port":8080}{"World":""}
sub
{"Help":false,"Host":"127.0.0.1","Port":8080}{"World":""}

+ echo =case 2: read from ENV variables
+ HOST=10.0.0.1
=case 2: read from ENV variables
+ go run 027-global-option.go sub
{"Help":false,"Host":"10.0.0.1","Port":8080}{"World":""}
sub
{"Help":false,"Host":"10.0.0.1","Port":8080}{"World":""}

+ echo =case 3: read from command line
+ HOST=10.0.0.1
+ go run 027-global-option.go sub -H 168.0.0.1
=case 3: read from command line
{"Help":false,"Host":"168.0.0.1","Port":8080}{"World":""}
sub
{"Help":false,"Host":"168.0.0.1","Port":8080}{"World":""}

+ HOST=10.0.0.1
+ go run 027-global-option.go sub -H 168.0.0.2 -w something
{"Help":false,"Host":"168.0.0.2","Port":8080}{"World":"something"}
sub
{"Help":false,"Host":"168.0.0.2","Port":8080}{"World":"something"}

+ rm 1.json

```

# 2016-06-05

After [issues#8](https://github.com/mkideal/cli/issues/8) is fixed.

## [014-time-and-duration.go](014-time-and-duration.go)

```sh
$ go run 014-time-and-duration.go -t '2016-01-02 03:04:05' -d=10ms
time=2016-01-02 03:04:05 +0000 UTC, duration=10ms
```

## [015-file.go](015-file.go)

```sh
$ echo hello | go run 015-file.go -f
hello

echo hello > /tmp/test.txt

$ go run 015-file.go -f /tmp/test.txt
hello

rm /tmp/test.txt
```

# 2016-06-09

After [issues#19](https://github.com/mkideal/cli/issues/19) is fixed.

## [028-reader.go](028-reader.go)

```sh
$ echo hello | go run 028-reader.go -r
read from file(or stdin): hello

reade from reader: string reader

$ echo hello A B C > /tmp/test.txt

$ go run 028-reader.go -r /tmp/test.txt
read from file(or stdin): hello A B C

reade from reader: string reader
```


## [029-writer.go](029-writer.go)

```sh
$ go run 029-writer.go -w
hello,writer
writes 13 byte(s) to file or stdout
writes 18 bytes to bytes.Writer: hello,bytes.Writer

$ go run 029-writer.go -w /tmp/test.txt
writes 13 byte(s) to file or stdout
writes 18 bytes to bytes.Writer: hello,bytes.Writer

$ cat /tmp/test.txt
hello,writer

rm /tmp/test.txt
```

