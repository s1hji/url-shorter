package postgres

import (
	"database/sql"
	"errors"
	"url-shorter/storage"

	"github.com/lib/pq"
)

type PostrgresStorage struct {
	db *sql.DB
}

const createTable = `CREATE TABLE IF NOT EXISTS links (
    id SERIAL PRIMARY KEY,
    origin TEXT UNIQUE NOT NULL,
    short VARCHAR(10) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);`

func NewPostgresStorage(connect string) (*PostrgresStorage, error) {
	db, err := sql.Open("postgres", connect)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if _, err := db.Exec(createTable); err != nil {
		return nil, err
	}

	return &PostrgresStorage{
		db: db,
	}, nil
}

func isUniqueViolation(err error) bool {
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		return pqErr.Code == "23505"
	}
	return errors.Is(err, storage.ErrorSaved)
}

func (p *PostrgresStorage) Save(origin string, short string) error {
	_, err := p.db.Exec(`INSERT INTO links (origin, short) VALUES ($1, $2)`, origin, short)
	if err != nil {
		if isUniqueViolation(err) {
			return storage.ErrorSaved
		}
		return err
	}

	return nil
}

func (p *PostrgresStorage) GetOriginLink(short string) (string, error) {
	var origin string
	err := p.db.QueryRow(`SELECT origin FROM links WHERE short = $1`, short).Scan(&origin)

	if errors.Is(err, sql.ErrNoRows) {
		return "", storage.ErrorNotFound
	}

	if err != nil {
		return "", err
	}

	return origin, nil
}

func (p *PostrgresStorage) GetShortLink(origin string) (string, error) {
	var short string
	err := p.db.QueryRow(`SELECT short FROM links WHERE origin = $1`, origin).Scan(&short)

	if errors.Is(err, sql.ErrNoRows) {
		return "", storage.ErrorNotFound
	}

	if err != nil {
		return "", err
	}

	return short, nil
}

func (p *PostrgresStorage) Close() error {
	return p.db.Close()
}
