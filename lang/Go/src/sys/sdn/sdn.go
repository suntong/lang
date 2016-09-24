////////////////////////////////////////////////////////////////////////////
// Purpose: sdn - Static Domain Name
// Authors: Tong Sun (c) 2016, All rights reserved
// Credits: https://gobyexample.com/line-filters
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// `Text` returns the current token, here the next line from the input.
		iline := scanner.Text()

		if regexp.MustCompile(`^\s*$|^\s*#`).FindStringIndex(iline) == nil {
			args := []string{
				"+short",
			}
			args = append(args, iline)
			cmd := exec.Command("dig", args...)
			out, err := cmd.Output()
			if err != nil {
				println(err.Error())
				os.Exit(1)
			}
			fmt.Printf("%s\t", regexp.MustCompile(`^[^0-9].*\n|\r*\n.*`).
				ReplaceAllString(string(out), ""))
		}
		fmt.Println(iline)
	}

	// Check for errors during `Scan`. End of file is
	// expected and not reported by `Scan` as an error.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

/*

Input format:

#Android
0.client-channel.google.com
1.client-channel.google.com
2.client-channel.google.com
3.client-channel.google.com
4.client-channel.google.com
5.client-channel.google.com
6.client-channel.google.com
7.client-channel.google.com
8.client-channel.google.com
9.client-channel.google.com
android.clients.google.com

Output;


#Android
209.85.200.189  0.client-channel.google.com
209.85.200.189  1.client-channel.google.com
209.85.200.189  2.client-channel.google.com
209.85.200.189  3.client-channel.google.com
209.85.200.189  4.client-channel.google.com
209.85.200.189  5.client-channel.google.com
209.85.200.189  6.client-channel.google.com
209.85.200.189  7.client-channel.google.com
209.85.200.189  8.client-channel.google.com
209.85.200.189  9.client-channel.google.com
172.217.0.174   android.clients.google.com

*/
