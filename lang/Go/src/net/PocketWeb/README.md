
Usage:

```
$ PocketWeb -h

Usage:
 PocketWeb [flags]

Flags:

  -d    directory holding web files (POCKETWEB_D)
  -p    port used by the pocketweb server (POCKETWEB_P)
  -dbg  debugging level (POCKETWEB_DBG)
  -h    print help then exit (POCKETWEB_H)
  -ver  print version then exit (POCKETWEB_VER)

Details:

  -d string
        directory holding web files (default ".")
  -dbg int
        debugging level
  -h    print help then exit
  -p string
        port used by the pocketweb server (default "8800")
  -ver
        print version then exit

E.g.:

  PocketWeb &
  PocketWeb -p 8088 -d /some/where/else
  POCKETWEB_D=/some/where/else POCKETWEB_P=8088 PocketWeb
  PocketWeb -ver

```
