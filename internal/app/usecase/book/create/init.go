package create

import (
	"context"

	"github.com/mkorobovv/mcp-server/internal/app/domain/book"
)

type BooksCreator struct {
	booksRepository booksRepository
}

type booksRepository interface {
	CreateBook(ctx context.Context, book book.Book) error
}

func NewBooksCreator(booksRepository booksRepository) *BooksCreator {
	return &BooksCreator{
		booksRepository: booksRepository,
	}
}
