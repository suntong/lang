// main.c
#include <stdio.h>
#include <unistd.h>

// Declare the stable API from libx.so
int x_init_logger(const char *path);
void x_log(const char *msg);
void x_close_logger(void);

int main() {
  if (!x_init_logger("app.log")) {
    printf("Failed to initialize logger\n");
    return 1;
  }

  x_log("Starting application from pure C");
  x_log("User 'alice' logged in");
  x_log("Processing payment of $999.99");
  x_log("Shutting down gracefully");

  sleep(1);			// allow fsync
  x_close_logger();

  printf("All messages logged → check app.log\n");
  return 0;

}
