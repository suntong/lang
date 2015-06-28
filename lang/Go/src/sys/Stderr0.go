package main

func main() {
  // the default built-in print prints to stderr
  // as documented in https://golang.org/pkg/builtin/#print
  // because print and println are for debugging so stderr is more appropriate
  print("To StdOut?\n")
}
