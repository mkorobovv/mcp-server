package books_repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/mkorobovv/mcp-server/internal/app/domain/book"
)

type BookDTO struct {
	BookID           uuid.UUID `db:"book_id"`
	AuthorID         uuid.UUID `db:"author_id"`
	Title            string    `db:"title"`
	Description      *string   `db:"description"`
	DateReleased     time.Time `db:"date_released"`
	CreatedAt        time.Time `db:"created_at"`
	Rating           float64   `db:"rating"`
	GenreID          int64     `db:"genre_id"`
	GenreName        string    `db:"genre_name"`
	GenreDescription *string   `db:"genre_description"`
}

func toEntities(dtos []BookDTO) []book.Book {
	books := make([]book.Book, 0, len(dtos))

	for _, dto := range dtos {
		book := book.Book{
			BookID:       dto.BookID,
			AuthorID:     dto.AuthorID,
			Title:        dto.Title,
			Description:  dto.Description,
			DateReleased: dto.DateReleased,
			CreatedAt:    dto.CreatedAt,
			Rating:       dto.Rating,
			Genre: book.Genre{
				ID:          dto.GenreID,
				Name:        dto.GenreName,
				Description: dto.GenreDescription,
			},
		}

		books = append(books, book)
	}

	return books
}
