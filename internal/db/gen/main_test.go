package gen

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/automagic-tools/go-coding-challenge/SeeCV/config"
	_ "github.com/lib/pq"
)

var (
	dbDriver = config.Vauban.Database.DBDriver
	dbSource = os.Getenv("DB_SOURCE")
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
