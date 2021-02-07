package repository

import (
	"database/sql"
)

type conn struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &conn{
		db: db,
	}
}

func ConnectDB(dbUrI string) *sql.DB {
	db, err := sql.Open("postgres", dbUrI)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
