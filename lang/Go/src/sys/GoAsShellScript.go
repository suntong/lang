/// 2>/dev/null; exec go run "$0" "$@"

// https://gist.github.com/posener/73ffd326d88483df6b1cb66e8ed1e0bd?permalink_comment_id=5554623#gistcomment-5554623

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello", os.Args[1])
	os.Exit(10)
}
