////////////////////////////////////////////////////////////////////////////
// Purpose: Go awk demo
// Authors: Tong Sun (c) 2015, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"

	"github.com/spakin/awk"
)

func main() {
	s := awk.NewScript()
	s.MaxRecordSize = 24 * 1024 * 1024
	s.MaxFieldSize = 24 * 1024 * 1024

	s.AppendStmt(func(s *awk.Script) bool {
		return s.F(1).Match("<Comment")
	}, func(s *awk.Script) {})
	// 1; # i.e., print all
	s.AppendStmt(nil, nil)

	if err := s.Run(os.Stdin); err != nil {
		panic(err)
	}
}
