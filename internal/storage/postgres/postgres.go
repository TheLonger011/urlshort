package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	createTableQuery := `
CREATE TABLE IF NOT EXISTS urls (
    id serial PRIMARY KEY,
    alias text NOT NULL UNIQUE,
    url text NOT NULL
)`

	if _, err := db.Exec(createTableQuery); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	createIndexQuery := `
		CREATE INDEX IF NOT EXISTS urls_alias ON urls(alias)
	`

	if _, err := db.Exec(createIndexQuery); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{db: db}, nil

}

func (s *Storage) SaveURL(urlToSave string, alias string) (int, error) {
	const op = "storage.postgres.SaveURL"
	query := `
	INSERT INTO urls(alias, url) VALUES ($1, $2) RETURNING id
`
	var id int
	err := s.db.QueryRow(query, alias, urlToSave).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	const op = "storage.postgres.GetURL"

	createTableQuery, err := s.db.Prepare("SELECT url FROM urls WHERE alias = $1")
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("%s: alias '%s' not found", op, alias)
		}
		return "", fmt.Errorf("%s: %w", op, err)
	}
	defer createTableQuery.Close()
	var resURL string

	err = createTableQuery.QueryRow(alias).Scan(&resURL)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)

	}
	return resURL, nil

}
