package book

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var ErrReleasedInFuture = errors.New("book cannot released in future")

type CreateBook struct {
	AuthorID     uuid.UUID `json:"author_id"`
	Title        string    `json:"title"`
	Description  *string   `json:"description"`
	DateReleased time.Time `json:"date_released"`
}

type Book struct {
	BookID       uuid.UUID `json:"book_id"`
	AuthorID     uuid.UUID `json:"author_id"`
	Title        string    `json:"title"`
	Description  *string   `json:"description"`
	DateReleased time.Time `json:"date_released"`
	CreatedAt    time.Time `json:"created_at"`
	Rating       float64   `json:"rating"`
	Genre        Genre     `json:"genre"`
}

func New(book CreateBook) (Book, error) {
	if book.DateReleased.After(time.Now()) {
		return Book{}, ErrReleasedInFuture
	}

	return Book{
		BookID:       uuid.New(),
		AuthorID:     book.AuthorID,
		Title:        book.Title,
		DateReleased: book.DateReleased,
		Description:  book.Description,
		Rating:       0,
	}, nil
}

type Genre struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type ListBooks struct {
	BookID       uuid.NullUUID `json:"book_id"`
	AuthorID     uuid.NullUUID `json:"author_id"`
	Title        *string       `json:"title"`
	GenreID      *int64        `json:"genre_id"`
	Limit        *int64        `json:"limit"`
	RatingHigher *float64      `json:"rating_higher"`
	RatingLower  *float64      `json:"rating_lower"`
	Offset       *int64        `json:"offset"`
}
