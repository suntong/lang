#include<stdio.h>
#include<stdlib.h>

// test

int main(int argc, char *argv[]) {
  int rows=9;
  int cols=9;
  //
  if (argc == 2) {
    rows=atoi(argv[1]);
    cols=atoi(argv[1]);
  }
  //
  int r;
  int c;
  rows=rows+1;
  cols=cols+1;
  for (r=1;r<rows;r=r+1) {
    for (c=1;c<cols;c=c+1) {
      if ((r==1) || (r==rows-1) || (c==1) || (c==cols-1) || (r==c) || (rows-r==c)) {
        printf("* ");
      } else {
        printf("  ");
      }
    }
    printf("\n");
  }
  exit(0);
}
