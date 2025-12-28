#!/bin/sh
#build.sh

set -e

echo "Building Go into c-archive..."
go build -trimpath -buildmode=c-archive -o logger.a go_logger.go

echo "Compiling C wrapper and linking into single libx.so..."
gcc -shared -pthread -o libx.so c_wrapper.c logger.a -ldl

echo "libx.so successfully built! ($(du -h libx.so | cut -f1))"

echo Compile the main program against the single lib
gcc -o app main.c -L. -lx -ldl

# Run
LD_LIBRARY_PATH=. ./app
cat app.log
