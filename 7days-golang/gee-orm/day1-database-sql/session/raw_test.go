package session

import (
	"database/sql"
	"os"
	"testing"

	"github.com/mattn/go-sqlite3"
)

var TestDB *sql.DB


func TestMain(m *testing.M) {
	TestDB, _ = sql.Open("sqlite3", "../gee.db")
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}