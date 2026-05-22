# REST API in Gin

A RESTful API built with Go and the Gin web framework.

## Features

- User authentication and authorization
- Event management
- Attendance tracking
- SQLite database
- Environment configuration
- Database migrations

## Project Structure

```
.
├── cmd/
│   ├── api/          # HTTP API server
│   └── migrate/      # Database migration tool
├── internal/
│   ├── database/     # Database models and operations
│   └── env/          # Environment variable handling
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

### Environment Variables

- `PORT`: Port to run the server on (default: 8080)
- `JWT_SECRET`: Secret key for JWT token signing
- Other variables as defined in `internal/env/env.go`

## API Endpoints

See [API Documentation](docs/api.md) for detailed endpoint information.

## Database Migrations

Migrations are managed using [golang-migrate](https://github.com/golang-migrate/migrate). 
To create a new migration:
```bash
migrate create -ext sql -dir cmd/migrate/migrations -seq <migration_name>
```

## License

This project is licensed under the MIT License.