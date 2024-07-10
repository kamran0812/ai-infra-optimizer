package storage

import (
	"database/sql"
	"time"

	"github.com/kamran0812/ai-infra-optimizer/internal/cloud"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(filename string) (*Database, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createTables(db); err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func createTables(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS resource_usage (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			provider TEXT,
			resource_id TEXT,
			type TEXT,
			cpu REAL,
			memory REAL,
			timestamp DATETIME
		)
	`)
	return err
}

func (d *Database) SaveResourceUsage(usages []cloud.ResourceUsage) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`
		INSERT INTO resource_usage (provider, resource_id, type, cpu, memory, timestamp)
		VALUES (?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, usage := range usages {
		_, err := stmt.Exec(usage.Provider, usage.ResourceID, usage.Type, usage.CPU, usage.Memory, time.Now())
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (d *Database) GetHistoricalUsage() ([]cloud.ResourceUsage, error) {
	rows, err := d.db.Query(`
		SELECT provider, resource_id, type, cpu, memory, timestamp
		FROM resource_usage
		ORDER BY timestamp DESC
		LIMIT 1000
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usages []cloud.ResourceUsage
	for rows.Next() {
		var usage cloud.ResourceUsage
		var timestamp time.Time
		if err := rows.Scan(&usage.Provider, &usage.ResourceID, &usage.Type, &usage.CPU, &usage.Memory, &timestamp); err != nil {
			return nil, err
		}
		usages = append(usages, usage)
	}

	return usages, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}
