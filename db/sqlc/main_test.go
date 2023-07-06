package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/stock_exchange?sslmode=disable"
)

var testDb *sql.DB


var testQueries *Queries

func TestMain(m *testing.M) {
	var err error

	testDb, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())

}
