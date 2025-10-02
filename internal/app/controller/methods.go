package controller

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mkorobovv/mcp-server/internal/app/domain/book"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ListBooksRequest struct {
	BookID       *string  `json:"book_id,omitempty" jsonschema:"a book id, uuid format"`
	AuthorID     *string  `json:"author_id,omitempty" jsonschema:"an author_id, uuid format"`
	Title        *string  `json:"title,omitempty" jsonschema:"a title"`
	GenreID      *int64   `json:"genre_id,omitempty" jsonschema:"a genre id"`
	RatingHigher *float64 `json:"rating_higher,omitempty" jsonschema:"a rating higher for sorting query by rating greater than, if not needed, send nothing"`
	RatingLower  *float64 `json:"rating_lower,omitempty" jsonschema:"a rating lower for sorting query by rating less than, if not needed, send nothing"`
	Limit        *int64   `json:"limit,omitempty" jsonschema:"a limit of query"`
	Offset       *int64   `json:"offset,omitempty" jsonschema:"an offset of query"`
}

type ListBooksResponse struct {
	Books []BookDTO `json:"books"`
}

func toResponse(books []book.Book) ListBooksResponse {
	dtoOut := make([]BookDTO, 0, len(books))

	for _, book := range books {
		dto := BookDTO{
			BookID:       book.BookID.String(),
			AuthorID:     book.AuthorID.String(),
			Title:        book.Title,
			Description:  book.Description,
			DateReleased: book.DateReleased.Format(time.DateOnly),
			CreatedAt:    book.CreatedAt.Format(time.RFC3339),
			Rating:       book.Rating,
			Genre: GenreDTO{
				ID:          book.Genre.ID,
				Name:        book.Genre.Name,
				Description: book.Genre.Description,
			},
		}

		dtoOut = append(dtoOut, dto)
	}

	return ListBooksResponse{dtoOut}
}

type BookDTO struct {
	BookID       string   `json:"book_id"`
	AuthorID     string   `json:"author_id"`
	Title        string   `json:"title"`
	Description  *string  `json:"description,omitempty"`
	DateReleased string   `json:"date_released"`
	CreatedAt    string   `json:"created_at"`
	Rating       float64  `json:"rating"`
	Genre        GenreDTO `json:"genre"`
}

type GenreDTO struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

func (l *ListBooksRequest) toRequest() (book.ListBooks, error) {
	req := book.ListBooks{
		Title:        l.Title,
		GenreID:      l.GenreID,
		RatingHigher: l.RatingHigher,
		RatingLower:  l.RatingLower,
		Limit:        l.Limit,
		Offset:       l.Offset,
	}

	if l.BookID != nil {
		id, err := uuid.Parse(*l.BookID)
		if err != nil {
			return book.ListBooks{}, err
		}

		req.BookID = uuid.NullUUID{UUID: id, Valid: true}
	}

	if l.AuthorID != nil {
		id, err := uuid.Parse(*l.AuthorID)
		if err != nil {
			return book.ListBooks{}, err
		}

		req.AuthorID = uuid.NullUUID{UUID: id, Valid: true}
	}

	return req, nil
}

func (c *Controller) ListBooks(ctx context.Context, _ *mcp.CallToolRequest, dtoIn ListBooksRequest) (*mcp.CallToolResult, ListBooksResponse, error) {
	c.logger.Debug("method called")

	request, err := dtoIn.toRequest()
	if err != nil {
		c.logger.Error(err.Error())

		return nil, ListBooksResponse{}, err
	}

	books, err := c.booksGetter.ListBooks(ctx, request)
	if err != nil {
		c.logger.Error(err.Error())

		return nil, ListBooksResponse{}, err
	}

	return nil, toResponse(books), nil
}

func (c *Controller) CreateBook(w http.ResponseWriter, r *http.Request) {
	var dtoIn CreateBookRequest

	err := json.NewDecoder(r.Body).Decode(&dtoIn)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			c.logger.Warn("Failed to write response body")
		}

		c.logger.Error(err.Error())

		return
	}

	req, err := dtoIn.toRequest()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			c.logger.Warn("Failed to write response body")
		}

		c.logger.Error(err.Error())

		return
	}

	id, err := c.booksCreator.Create(r.Context(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			c.logger.Warn("Failed to write response body")
		}

		c.logger.Error(err.Error())

		return
	}

	dtoOut := CreateBookResponse{ID: id}

	err = json.NewEncoder(w).Encode(dtoOut)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		_, writeErr := w.Write([]byte(err.Error()))
		if writeErr != nil {
			c.logger.Warn("Failed to write response body")
		}

		c.logger.Error(err.Error())

		return
	}
}

type CreateBookRequest struct {
	AuthorID     string  `json:"author_id"`
	Title        string  `json:"title"`
	Description  *string `json:"description"`
	DateReleased string  `json:"date_released"`
}

type CreateBookResponse struct {
	ID uuid.UUID `json:"id"`
}

func (c *CreateBookRequest) toRequest() (book.CreateBook, error) {
	dateReleased, err := time.Parse(time.DateOnly, c.DateReleased)
	if err != nil {
		return book.CreateBook{}, err
	}

	authorID, err := uuid.Parse(c.AuthorID)
	if err != nil {
		return book.CreateBook{}, err
	}

	return book.CreateBook{
		AuthorID:     authorID,
		Title:        c.Title,
		Description:  c.Description,
		DateReleased: dateReleased,
	}, nil
}
