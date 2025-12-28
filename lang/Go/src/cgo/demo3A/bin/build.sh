# 1. Compile the consumer
# -L. looks for libs in current dir, -lx links against libx.so
gcc -o app main.c -L../lib -lx -lpthread
# Run
LD_LIBRARY_PATH=../lib app
cat system.log
