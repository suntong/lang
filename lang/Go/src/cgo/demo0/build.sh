go build -o libgo.so -buildmode=c-shared golib.go
gcc -o loader loader.c ./libgo.so
./loader