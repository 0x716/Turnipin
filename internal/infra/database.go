package infra

import (
	"database/sql"

	"github.com/0x716/Turnipin/internal/config"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sql.DB
)

func InitDB() error {
	// Connect Sqlite Database (if file doesn't exist then create sqlite file)
	db, err := sql.Open(config.GlobalConfig.Database.Title, config.GlobalConfig.Database.Name)
	if err != nil {
		return err
	}

	// test connection status
	if err := db.Ping(); err != nil {
		return err
	}

	DB = db

	return nil
}
