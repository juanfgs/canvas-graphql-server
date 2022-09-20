package database

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"log"
	"os"
	"path"
)

var Db *sql.DB

func InitDB() {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	if os.Getenv("GO_ENV") == "test" {
		dbUsername = os.Getenv("DB_TEST_USERNAME")
		dbPassword = os.Getenv("DB_TEST_PASSWORD")
		dbHost = os.Getenv("DB_TEST_HOST")
		dbName = os.Getenv("DB_TEST_NAME")
		dbPort = os.Getenv("DB_TEST_PORT")
	}
	db, err := sql.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable port=%s", dbUsername, dbPassword, dbHost, dbName, dbPort))
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
	gopath := os.Getenv("GOPATH")
	if err := Db.Ping(); err != nil {

		log.Fatal(err)
	}

	driver, _ := postgres.WithInstance(Db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+path.Join(gopath, "/src/github.com/juanfgs/canvas/internal/pkg/db/migrations/postgresql"),
		"postgresql",
		driver,
	)
	log.Println(err)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

}
