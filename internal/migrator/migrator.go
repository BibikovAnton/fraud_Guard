package migrator

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

type Migrator struct {
	db           *sql.DB
	migratinsDir string
}

func NewMigrator(db *sql.DB, migratinsDir string) *Migrator {
	return &Migrator{
		db:           db,
		migratinsDir: migratinsDir,
	}
}

func (m *Migrator) Up() error {
	err := goose.Up(m.db, m.migratinsDir)
	if err != nil {
		return err
	}
	return nil
}
