go build -o libgo.so -buildmode=c-shared printer.go
gcc -o loader loader.c ./libgo.so
./loader