# REST API in Gin

A RESTful API built with Go and the Gin web framework.

## Features

- User authentication and authorization (JWT-based)
- Event management (CRUD operations)
- Attendance tracking
- SQLite database
- Environment configuration
- Database migrations using golang-migrate
- RESTful API design with versioning
- Rate limiting
- Pagination support
- Standardized JSON response format

## Project Structure

```
.
├── cmd/
│   ├── api/          # HTTP API server
│   └── migrate/      # Database migration tool
├── internal/
│   ├── database/     # Database models and operations
│   └── env/          # Environment variable handling
├── docs/             # API documentation
├── tmp/              # Temporary files (build logs, etc.)
└── data.db           # SQLite database (generated)
```

## Getting Started

### Prerequisites

- Go 1.26 or higher
- SQLite3

### Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Set up environment variables (copy `.env.example` to `.env` and adjust):
   ```bash
   cp .env.example .env
   ```
4. Run database migrations:
   ```bash
   go run ./cmd/migrate
   ```
5. Start the server:
   ```bash
   go run ./cmd/api
   ```
   The API will be available at `http://localhost:8080/api/v1`.

### Environment Variables

- `PORT`: Port to run the server on (default: 8080)
- `JWT_SECRET`: Secret key for JWT token signing
- Other variables as defined in `internal/env/env.go`

## API Endpoints

See [API Documentation](docs/api.md) for detailed endpoint information including:
- Authentication requirements
- Request/response formats
- Error codes
- Rate limiting
- Pagination

## Database Migrations

Migrations are managed using [golang-migrate](https://github.com/golang-migrate/migrate). 
To create a new migration:
```bash
migrate create -ext sql -dir cmd/migrate/migrations -seq <migration_name>
```

Available migration commands:
- Apply migrations: `go run ./cmd/migrate`
- Create new migration: `migrate create -ext sql -dir cmd/migrate/migrations -seq <name>`

## License

This project is licensed under the MIT License.