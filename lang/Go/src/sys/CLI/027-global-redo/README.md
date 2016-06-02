## 2016-06-01

```sh
$ redo 
global option redo

  redo global option via automatic code-gen

Options:

  -h, --help
      display help information

  -c, --config[=redo.json]
      config file

  -H, --host[=$HOST]
      host address

  -p, --port
      listening port

Commands:

  build     Build the network application
  install   Install the network application
  publish   Publish the network application


$ redo build
Build the network application

Usage:
  redo build [Options] Arch(i386|amd64)

Options:

  -h, --help
      display help information

  -c, --config[=redo.json]
      config file

  -H, --host[=$HOST]
      host address

  -p, --port
      listening port

  --dir[=./]
      source code root dir

  --suffix[=.go,.c,.s]
      source file suffix

  -o, --out
      output filename


$ redo build i386
[build]:
  {"Help":false,"Host":"127.0.0.1","Port":8080}
  {"Dir":"./","Suffix":".go,.c,.s","Out":""}
  [i386]


$ HOST=10.0.0.1 redo build i386
[build]:
  {"Help":false,"Host":"10.0.0.1","Port":8080}
  {"Dir":"./","Suffix":".go,.c,.s","Out":""}
  [i386]


$ HOST=10.0.0.1 redo build -H 168.0.0.1 i386
[build]:
  {"Help":false,"Host":"168.0.0.1","Port":8080}
  {"Dir":"./","Suffix":".go,.c,.s","Out":""}
  [i386]


$ redo install
Install the network application

Usage:
  redo install [Options] package [package...]

Options:

  -h, --help
      display help information

  -c, --config[=redo.json]
      config file

  -H, --host[=$HOST]
      host address

  -p, --port
      listening port

  --dir[=./]
      source code root dir

  --suffix[=.go,.c,.s]
      source file suffix

  -o, --out
      output filename


$ HOST=10.0.0.1 redo install -H 168.0.0.2 pkg{1,2,3}
[install]:
  {"Help":false,"Host":"168.0.0.2","Port":8080}
  {"Dir":"./","Suffix":".go,.c,.s","Out":""}
  [pkg1 pkg2 pkg3]


$ HOST=10.0.0.1 redo install -H 168.0.0.2 -p 8088 --dir /tmp pkg{1,2,3}
[install]:
  {"Help":false,"Host":"168.0.0.2","Port":8088}
  {"Dir":"/tmp","Suffix":".go,.c,.s","Out":""}
  [pkg1 pkg2 pkg3]


$ redo publish -H my.server.net
[publish]:
  {"Help":false,"Host":"my.server.net","Port":8080}
  {"Dir":"./","Suffix":".go,.c,.s","Out":"","List":false}
  []
```
