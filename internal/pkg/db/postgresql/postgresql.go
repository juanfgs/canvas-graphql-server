package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"log"
	"os"
	"fmt"
)

var Db *sql.DB

func InitDB() {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",dbUsername,dbPassword,dbHost,dbName))
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
 		log.Panic(err)
	}
	Db = db
}

func CloseDB() error {
	return Db.Close()
}

func Migrate() {
	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
	driver, _ := postgres.WithInstance(Db, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/postgresql",
		"postgresql",
		driver,
	)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

}
