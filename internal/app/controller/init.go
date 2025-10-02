package controller

import (
	"log/slog"

	bookscreator "github.com/mkorobovv/mcp-server/internal/app/usecase/book/create"
	booksgetter "github.com/mkorobovv/mcp-server/internal/app/usecase/book/get"
)

type Controller struct {
	logger       *slog.Logger
	booksCreator *bookscreator.BooksCreator
	booksGetter  *booksgetter.BooksGetter
}

func NewController(logger *slog.Logger, booksCreator *bookscreator.BooksCreator, booksGetter *booksgetter.BooksGetter) *Controller {
	return &Controller{
		logger:       logger,
		booksCreator: booksCreator,
		booksGetter:  booksGetter,
	}
}
