package fixture

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"io/fs"
	"log"
)

func Up(migrationsDir fs.FS) *sql.DB {
	db, err := sql.Open("sqlite3", "testing.db")
	if err != nil {
		log.Fatal(err)
	}
	migrations := NewMigrationsDb(db)
	return migrations.Up(migrationsDir)
}