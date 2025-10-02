package get

import (
	"context"

	"github.com/mkorobovv/mcp-server/internal/app/domain/book"
)

type BooksGetter struct {
	booksRepository booksRepository
}

type booksRepository interface {
	ListBooks(ctx context.Context, listBooksRequest book.ListBooks) ([]book.Book, error)
}

func NewBooksGetter(booksRepository booksRepository) *BooksGetter {
	return &BooksGetter{
		booksRepository: booksRepository,
	}
}
