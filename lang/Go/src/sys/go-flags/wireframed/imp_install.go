////////////////////////////////////////////////////////////////////////////
// Program: redo
// Purpose: global option redo
// Authors: Myself <me@mine.org> (c) 2022, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/go-easygen/go-flags/clis"
)

// *** Sub-command: install ***
// Exec implements the business logic of command `install`
func (x *InstallCommand) Exec(args []string) error {
	err := fmt.Errorf("Sample warning: Instance not found")
	clis.WarnOn("Install, Exec", err)
	// or,
	// clis.AbortOn("Doing Install", err)
	return nil
}
