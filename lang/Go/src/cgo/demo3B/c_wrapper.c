// c_wrapper.c
#include <stdlib.h>

// Go-exported functions (generated in the .a)
extern int InitLogger(const char *path);
extern void LogMessage(const char *msg);
extern void CloseLogger(void);

// Public API with stable ABI and default visibility
__attribute__((visibility("default")))
int x_init_logger(const char *path) {
  if (!path || path[0] == '\0')
    return 0;
  return InitLogger(path);
}

__attribute__((visibility("default")))
void x_log(const char *msg) {
  if (!msg)
    return;
  LogMessage(msg);
}

__attribute__((visibility("default")))
void x_close_logger(void) {
  CloseLogger();
}
