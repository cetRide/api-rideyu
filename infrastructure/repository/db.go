package repository

import "database/sql"

type conn struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &conn{
		db: db,
	}
}
