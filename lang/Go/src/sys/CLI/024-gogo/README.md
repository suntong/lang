# Multi command application

Try type these commands in terminal

```sh
go build -o gogo
./gogo --help
./gogo --version
./gogo --list
./gogo build
./gogo build --help
./gogo build --dir=.. -o out
./gogo clean
./gogo doc --help
./gogo install
./gogo install --help
...
```

## 2016-05-26

    $ go build 

    $ 024-gogo
    try ` --help for more information'
    
    $ 024-gogo --help
    Golang package manager
    
    gogo is a new golang package manager
    very very good
    
    Options:
    
      -h, --help
          display help information
    
      -v, --version
          display version
    
      -l, --list[=false]
          list all sub commands or not
    
    Commands:
      build     Build golang application
      clean     Clean build data
      doc       Generate documents
      install   Install golang application
      publish   Publish golang application
      test      Test golang application
    
    $ 024-gogo --version
    v1.0.0
    
    $ 024-gogo --list
     build    =>  Build golang application
     clean    =>  Clean build data
     doc      =>  Generate documents
     install  =>  Install golang application
     publish  =>  Publish golang application
     test     =>  Test golang application
    
    $ 024-gogo build
    build: {
        "Help": false,
        "Dir": "./",
        "Suffix": ".go,.c,.s",
        "Out": ""
    }
    
    $ 024-gogo build --help
    Build golang application
    
    Options:
    
      -h, --help
          display help information
    
      --dir[=./]
          source code root dir
    
      --suffix[=.go,.c,.s]
          source file suffix
    
      -o, --out
          output filename
    
    $ 024-gogo build --dir=.. -o out
    build: {
        "Help": false,
        "Dir": "..",
        "Suffix": ".go,.c,.s",
        "Out": "out"
    }
    
    $ 024-gogo clean
    clean: {
        "Help": false,
        "Recursion": true
    }
    
    $ 024-gogo doc --help
    Generate documents
    
    Options:
    
      -h, --help
          display help information
    
      --suffix[=.go,.c,.s]
          source file suffix
    
      -o, --out
          output filename
    
    
    $ 024-gogo install
    install: {
        "Help": false,
        "Dir": "./",
        "Suffix": ".go,.c,.s",
        "Out": ""
    }
    
    $ 024-gogo install --help
    Install golang application
    
    Options:
    
      -h, --help
          display help information
    
      --dir[=./]
          source code root dir
    
      --suffix[=.go,.c,.s]
          source file suffix
    
      -o, --out
          output filename

## 2016-05-27

    $ 024-gogo
    Golang package manager
    
      gogo is a new golang package manager
      very very good
    
    Options:
    
      -h, --help
          display help information
    
      -v, --version
          display version
    
      -l, --list[=false]
          list all sub commands or not
    
    Commands:
      build     Build golang application
      clean     Clean build data
      doc       Generate documents
      install   Install golang application
      publish   Publish golang application
      test      Test golang application
    
    $ 024-gogo build
    Build golang application
    
    Usage:
      gogo build [Options] Arch(i386|amd64)
    
    Options:
    
      -h, --help
          display help information
    
      --dir[=./]
          source code root dir
    
      --suffix[=.go,.c,.s]
          source file suffix
    
      -o, --out
          output filename
    
    $ 024-gogo build --dir=.. -o out i386
    build: {
        "Help": false,
        "Dir": "..",
        "Suffix": ".go,.c,.s",
        "Out": "out"
    }
    Arch: i386

