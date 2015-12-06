////////////////////////////////////////////////////////////////////////////
// Porgram: Wp2Hugo
// Purpose: From wordpress meta to Hugo's
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

import (
	"github.com/spakin/awk"
)

/*

Wp2Hugo converts wordpress's markdown meta data format into Hugo's.

E.g., for an input of wordpress's markdown like this,

    # Dbab From Start To Finish

    [category Tech][tags Debian,Ubuntu,Linux,DHCP,DNS,WPAD,dnsmasq,dbab]

The output Hugo's meta data is:

    ---
    title: "Dbab From Start To Finish"
    date: "2015-12-06T09:57:45-05:00"
    categories: ["Tech"]
    tags: ["Debian","Ubuntu","Linux","DHCP","DNS","WPAD","dnsmasq","dbab"]
    ---

Usage:

    Wp2Hugo < wordpress.md > path/to/hugo.md


*/

func main() {
	s := awk.NewScript()

	s.Begin = func(s *awk.Script) {
		s.SetFS("[][ ]")
	}
	s.State = s.NewValue("")

	// NR == 1 & $0 ~ /^# +/ { State = $0 }
	s.AppendStmt(func(s *awk.Script) bool { return s.NR == 1 && s.F(0).Match("^# +") },
		func(s *awk.Script) {
			s.State = s.F(0)
			s.Next()
		})

	// /^\[category/ { convert meta data format }
	s.AppendStmt(func(s *awk.Script) bool { return s.F(0).Match("^\\[category") },
		func(s *awk.Script) {
			re := regexp.MustCompile("^# +(.*)$")
			fmt.Println(re.ReplaceAllString(s.State.(*awk.Value).String(),
				"---\ntitle: \"$1\""))
			fmt.Printf("date: \"%s\"\n", time.Now().Format(time.RFC3339))
			fmt.Printf("categories: [\"%s\"]\n", s.F(3))
			fmt.Printf("%s: [\"%s\"]\n", s.F(5), s.F(6))
			fmt.Println("---\n")
			//s.Exit()
			s.Next()
		})

	// 1; # i.e., print all
	s.AppendStmt(nil, nil)

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}
