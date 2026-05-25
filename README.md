# REST API in Gin

A RESTful event management API built with Go and the Gin web framework.

## Features

- User registration with bcrypt password hashing
- Event management (CRUD operations)
- Attendee registration for events
- SQLite database with golang-migrate schema migrations
- RESTful API design with versioning (`/api/v1`)
- Live reload development with Air

## Project Structure

```
.
├── cmd/
│   ├── api/              # HTTP API server
│   │   ├── main.go       # Application entry point
│   │   ├── routes.go     # Route definitions
│   │   ├── server.go     # HTTP server configuration
│   │   ├── auth.go       # User registration handler
│   │   ├── events.go     # Event CRUD and attendee handlers
│   │   └── users.go      # User listing/deletion handlers
│   └── migrate/          # Database migration tool
│       ├── main.go       # Migration CLI runner
│       └── migrations/   # SQL migration files
├── internal/
│   ├── database/         # Database models and operations
│   │   ├── models.go     # Models container + constructor
│   │   ├── users.go      # User model + CRUD operations
│   │   ├── events.go     # Event model + CRUD operations
│   │   └── attendes.go   # Attendee model + CRUD operations
│   └── env/              # Environment variable handling
│       └── env.go        # GetEnvString / GetEnvInt helpers
├── docs/                 # Documentation
│   ├── api.md            # API endpoint reference
│   └── architecture.md   # System architecture overview
├── .env                  # Environment variables
├── .air.toml             # Air live-reload configuration
├── data.db               # SQLite database (generated)
└── commit.sh             # Interactive git commit/push helper
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
3. Run database migrations:
   ```bash
   go run ./cmd/migrate
   ```
4. Start the server:
   ```bash
   go run ./cmd/api
   ```
   The API will be available at `http://localhost:8080`.

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Port to run the server on |
| `JWT_SECRET` | `secret` | Reserved for future JWT authentication |

### Live Reload (Development)

[Air](https://github.com/cosmtrek/air) is configured for live reload. Start with:

```bash
air
```

The server rebuilds automatically when `.go` files change (see `.air.toml` for configuration).

## API Endpoints

See [API Documentation](docs/api.md) for detailed endpoint information including request/response formats and error codes.

**Quick reference:**

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/ping` | Health check |
| `GET` | `/api/v1/` | API status |
| `POST` | `/api/v1/auth/register` | Register a new user |
| `GET` | `/api/v1/users` | List all users |
| `DELETE` | `/api/v1/delete/:id` | Delete a user |
| `POST` | `/api/v1/events` | Create an event |
| `GET` | `/api/v1/events` | List all events |
| `GET` | `/api/v1/events/:id` | Get an event by ID |
| `PUT` | `/api/v1/events/:id` | Update an event |
| `DELETE` | `/api/v1/events/:id` | Delete an event |
| `DELETE` | `/api/v1/events` | Delete all events |
| `POST` | `/api/v1/events/:id/attendees/:userId` | Add attendee to event |
| `GET` | `/api/v1/events/:id/attendees` | List attendees for event |
| `DELETE` | `/api/v1/events/:id/attendees/:userId` | Remove attendee from event |
| `GET` | `/api/v1/attendees/:id/events` | Get events by attendee |

## Database Migrations

Migrations are managed using [golang-migrate](https://github.com/golang-migrate/migrate).

**Schema:**
- `users` (id, email, name, password)
- `events` (id, owner_id, name, description, date, location) — `owner_id` references `users(id)` ON DELETE CASCADE
- `attendees` (id, user_id, event_id) — both foreign keys ON DELETE CASCADE

**Commands:**

| Command | Description |
|---------|-------------|
| `go run ./cmd/migrate` | Apply all pending migrations |
| `go run ./cmd/migrate down` | Rollback last migration |
| `migrate create -ext sql -dir cmd/migrate/migrations -seq <name>` | Create a new migration pair |

## License

This project is licensed under the MIT License.
