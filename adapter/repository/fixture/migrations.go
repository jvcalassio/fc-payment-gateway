package fixture

import (
	"context"
	"database/sql"
	"github.com/maragudk/migrate"
	"io/fs"
)

type MigrationsDb struct {
	db *sql.DB
}

func NewMigrationsDb(db *sql.DB) *MigrationsDb {
	return &MigrationsDb{db: db}
}

func (db *MigrationsDb) Up(migrationsDir fs.FS) *sql.DB { 
	if err := migrate.Up(context.Background(), db.db, migrationsDir); err != nil {
		panic(err)
	}
	return db.db
}

func Down(db *sql.DB, migrationsDir fs.FS) {
	if err := migrate.Down(context.Background(), db, migrationsDir); err != nil {
		panic(err)
	}
	db.Close()
}