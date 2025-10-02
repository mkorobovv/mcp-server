package books_repository

import "github.com/jmoiron/sqlx"

type BooksRepository struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) *BooksRepository {
	return &BooksRepository{
		DB: db,
	}
}
