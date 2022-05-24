package migrations

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	"gorm.io/gorm"
)

const DefaultPath = "migrations/sql"

func Run(dbConn *gorm.DB, args []string) {
	executeCmd(dbConn, args)
}

func executeCmd(dbConn *gorm.DB, args []string) {

	// migrate := getMigrate(dbConn)

	switch args[0] {
	case "create":
		create(DefaultPath, args[1], time.Now().Unix())
	}
}

func getMigrate(dbConn *gorm.DB) *migrate.Migrate {
	sqlDB, err := dbConn.DB()
	if err != nil {
		log.Fatal(err)
	}

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}

	migrate, err := migrate.NewWithDatabaseInstance(
		"file://"+DefaultPath,
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}

	return migrate
}

func create(path, name string, timestamp int64) {
	base := fmt.Sprintf("%v_%v.", timestamp, name)
	_ = os.MkdirAll(path, os.ModePerm)
	createFile(base + "up.sql")
	createFile(base + "down.sql")
}

func createFile(name string) {
	if _, err := os.Create(name); err != nil {
		log.Fatal(err)
	}
}
