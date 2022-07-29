package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"

	config "back_go/pkg/config"
)

type DbContainer struct {
	Db *sql.DB
}

func NewDbContainer() DbContainer {
	database_host, err := config.GetDatabaseHost()
	if err != nil {
		log.Panic(err)
	}

	db, err := sql.Open("mysql", "admin_go:deuxmillekangourous@tcp("+database_host+":3306)/testgo?multiStatements=true&parseTime=true")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return DbContainer{
		Db: db,
	}
}

func (c DbContainer) CloseDB() error {
	return c.Db.Close()
}

func (c DbContainer) Migrate() {
	if err := c.Db.Ping(); err != nil {
		log.Fatal(err)
	}

	driver, _ := mysql.WithInstance(c.Db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://internal/pkg/db/migrations/mysql",
		"mysql",
		driver,
	)

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
