# MCP Server

The **MCP Server** project is a backend system for managing data and handling requests using the Model Context Protocol (MCP).

## 📁 Project Structure

- **`main.go`** — entry point of the application, sets up routes and starts the server.
- **`go.mod` and `go.sum`** — Go dependency management files.
- **`internal/`** — internal application packages:
    - **`app/`** — business logic and request handling.
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

The server will be available at http://localhost:8080
