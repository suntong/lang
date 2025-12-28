go build -o libgo.so -buildmode=c-shared printer.go
g++ -o loader main.cpp ./libgo.so
./loader
