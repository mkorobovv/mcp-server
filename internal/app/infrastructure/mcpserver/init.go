package mcpserver

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/mkorobovv/mcp-server/internal/app/controller"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"golang.org/x/sync/errgroup"
)

type MCPServer struct {
	logger *slog.Logger
	server *http.Server
}

func New(logger *slog.Logger, ctr *controller.Controller) *MCPServer {
	mcpServer := mcp.NewServer(&mcp.Implementation{
		Name:    "books",
		Version: "v1.0.0",
	}, nil)

	mcp.AddTool(mcpServer, &mcp.Tool{
		Name:        "list_books",
		Description: "Shows a list of books by filters",
	}, ctr.ListBooks)

	handler := mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
		switch req.URL.Path {
		case "/mcp":
			return mcpServer
		default:
			return nil
		}
	}, nil)

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	return &MCPServer{
		logger: logger,
		server: httpServer,
	}
}

func (s *MCPServer) Start(ctx context.Context) error {
	s.logger.Info("starting server")

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := s.server.Shutdown(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	g.Go(func() error {
		err := s.server.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				s.logger.Info("server shutdown gracefully")
				return nil
			} else {
				return err
			}
		}

		return nil
	})

	err := g.Wait()
	if err != nil {
		return err
	}

	return nil
}
