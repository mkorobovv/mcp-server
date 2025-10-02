package create

import (
	"context"

	"github.com/google/uuid"
	"github.com/mkorobovv/mcp-server/internal/app/domain/book"
)

func (c *BooksCreator) Create(ctx context.Context, createBookRequest book.CreateBook) (id uuid.UUID, err error) {
	book, err := book.New(createBookRequest)
	if err != nil {
		return uuid.Nil, err
	}

	err = c.booksRepository.CreateBook(ctx, book)
	if err != nil {
		return uuid.Nil, err
	}

	return book.BookID, nil
}
