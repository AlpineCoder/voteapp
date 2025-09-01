package repository

import "fmt"

var (
	ErrNoRows = fmt.Errorf("sql: no rows in result set")
)
