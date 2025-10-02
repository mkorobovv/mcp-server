package books_repository

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/mkorobovv/mcp-server/internal/app/domain/book"
)

func (repo *BooksRepository) CreateBook(ctx context.Context, book book.Book) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	qb := psql.
		Insert("books.books").
		Columns(
			"book_id",
			"author_id",
			"title",
			"date_released",
			"rating",
		).
		Values(
			book.BookID.String(),
			book.AuthorID.String(),
			book.Title,
			book.DateReleased.Format(time.DateOnly),
			book.Rating,
		)

	if book.Description != nil {
		qb = qb.
			Columns("description").
			Values(*book.Description)
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return err
	}

	_, err = repo.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (repo *BooksRepository) ListBooks(ctx context.Context, request book.ListBooks) (books []book.Book, err error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	qb := psql.
		Select(
			"b.book_id",
			"b.author_id",
			"b.title",
			"b.description",
			"b.date_released",
			"b.created_at",
			"b.rating",
			"g.genre_id",
			"g.name as genre_name",
			"g.description as genre_description",
		).
		From("books.books b").
		LeftJoin("books.genres g ON b.genre_id = g.genre_id").
		OrderBy("b.date_released DESC")

	if request.BookID.Valid {
		qb = qb.
			Where(
				squirrel.Eq{
					"b.book_id": request.BookID.UUID.String(),
				},
			)
	}

	if request.AuthorID.Valid {
		qb = qb.
			Where(
				squirrel.Eq{
					"b.author_id": request.AuthorID.UUID.String(),
				},
			)
	}

	if request.Title != nil {
		qb = qb.
			Where(
				squirrel.Eq{
					"b.title": *request.Title,
				},
			)
	}

	if request.GenreID != nil {
		qb = qb.
			Where(
				squirrel.Eq{
					"b.genre_id": *request.GenreID,
				},
			)
	}

	if request.RatingHigher != nil {
		qb = qb.
			Where(
				squirrel.GtOrEq{
					"b.rating": *request.RatingHigher,
				},
			)
	}

	if request.RatingLower != nil {
		qb = qb.
			Where(
				squirrel.LtOrEq{
					"b.rating": *request.RatingLower,
				},
			)
	}

	if request.Limit != nil {
		qb = qb.Limit(uint64(*request.Limit))
	} else {
		qb = qb.Limit(10)
	}

	if request.Offset != nil {
		qb = qb.Offset(uint64(*request.Offset))
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, err
	}

	var dto []BookDTO

	err = repo.DB.SelectContext(ctx, &dto, query, args...)
	if err != nil {
		return nil, err
	}

	books = toEntities(dto)

	return books, nil
}
