# MCP Server

The **MCP Server** project is a backend system for managing data and handling requests using the Model Context Protocol (MCP).

## 📁 Project Structure

- **`main.go`** — entry point of the application, sets up routes and starts the server.
- **`go.mod` and `go.sum`** — Go dependency management files.
- **`internal/`** — internal application packages:
    - **`app/`** — business logic and request handling.
    - **`controller`** — mcp tools 
- **`migrations/`** — database migration scripts.
- **`deploy/`** — deployment configurations (Dockerfile, CI/CD scripts, etc.).
- **`LICENSE`** — project license (MIT).

## ⚙️ Installation and Running

1. Clone the repository:
   ```shell
   git clone https://github.com/mkorobovv/mcp-server.git
   cd mcp-server
   ```

2. Install dependencies:
    ```shell
    go mod tidy
    ```
   
3. Run migrations with goose:
    ```shell
    goose -dir ./migrations postgres "host=127.0.0.1 port=5432 user=books_service password=admin123 dbname=books sslmode=disable TimeZone=Europe/Moscow" up
    ```

4. Run server:
    ```shell
    go run main.go
    ```

The server will be available at http://localhost:8080
