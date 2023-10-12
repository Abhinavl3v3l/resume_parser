package db

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"sync"

	"github.com/automagic-tools/go-coding-challenge/SeeCV/config"
	"github.com/automagic-tools/go-coding-challenge/SeeCV/internal/db/gen"
	"github.com/automagic-tools/go-coding-challenge/SeeCV/utils/logger"
	_ "github.com/lib/pq"
)

type Database struct {
	Connection *sql.DB
	Queries    *gen.Queries
}

var (
	instance *Database
	once     sync.Once
)

func GetDBInstance() (*Database, error) {
	if instance == nil {
		err := errors.New("database not initialized")
		logger.Error("Cannot get DBinstance.", err)
		return nil, nil
	}
	return instance, nil
}

func StartDB() error {
	once.Do(func() {
		connStr := os.Getenv("DB_SOURCE")
		if connStr == "" {
			log.Fatalf("DB_SOURCE is not set.")
		}

		db, err := sql.Open(config.Vauban.Database.DBDriver, connStr)
		if err != nil {
			log.Fatalf("Failed to open DB connection: %v", err)
			return
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Failed to ping DB: %v", err)
			return
		}

		instance = &Database{
			Connection: db,
			Queries:    gen.New(db),
		}
	})

	if instance == nil {
		return errors.New("failed to initialize the database")
	}
	return nil
}

func CloseDB() error {
	if instance == nil || instance.Connection == nil {
		return nil
	}
	return instance.Connection.Close()
}
