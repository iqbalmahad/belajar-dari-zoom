package database

import (
	"fmt"
	"testing"
)

func TestOpenConnection(t *testing.T) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	db := DBConnection()

	fmt.Print(db)
}
