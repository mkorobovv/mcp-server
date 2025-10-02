package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/mkorobovv/mcp-server/internal/app/controller"
	"github.com/mkorobovv/mcp-server/internal/app/infrastructure/mcpserver"
	"github.com/mkorobovv/mcp-server/internal/app/infrastructure/postgres"
	books_repository "github.com/mkorobovv/mcp-server/internal/app/repository/books-repository"
	bookscreator "github.com/mkorobovv/mcp-server/internal/app/usecase/book/create"
	booksgetter "github.com/mkorobovv/mcp-server/internal/app/usecase/book/get"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	booksDB, err := postgres.Sqlx(logger, postgres.Config{
		Host:     "localhost",
		Port:     "5432",
		Name:     "books",
		User:     "books_service",
		Password: "admin123",
		TimeZone: "Europe/Moscow",
	})
	if err != nil {
		logger.Error("Failed to connect to database", slog.Any("error", err))

		panic(err)
	}

	booksRepository := books_repository.New(booksDB)

	booksCreator := bookscreator.NewBooksCreator(booksRepository)
	booksGetter := booksgetter.NewBooksGetter(booksRepository)

	ctr := controller.NewController(logger, booksCreator, booksGetter)

	mcpServer := mcpserver.New(logger, ctr)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	defer stop()

	if err := mcpServer.Start(ctx); err != nil {
		logger.Error("server stopped with error", "err", err)
	}
}
