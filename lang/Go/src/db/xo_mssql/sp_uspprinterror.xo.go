package xo_mssql

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Uspprinterror calls the stored procedure 'dbo.uspPrintError()' on db.
func Uspprinterror(ctx context.Context, db DB) error {
	// call dbo.uspPrintError
	const sqlstr = `dbo.uspPrintError`
	// runlogf(sqlstr)
	if _, err := db.ExecContext(ctx, sqlstr); err != nil {
		return logerror(err)
	}
	return nil
}
