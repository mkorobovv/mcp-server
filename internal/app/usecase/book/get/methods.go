package get

import (
	"context"

	"github.com/mkorobovv/mcp-server/internal/app/domain/book"
)

func (g *BooksGetter) ListBooks(ctx context.Context, listBooksRequest book.ListBooks) ([]book.Book, error) {
	books, err := g.booksRepository.ListBooks(ctx, listBooksRequest)
	if err != nil {
		return nil, err
	}

	return books, nil
}
