# System Architecture

## Overview

This REST API is built using Go and the Gin web framework. It follows a layered architecture separating HTTP handling, business logic, and data persistence concerns.

## Architecture Diagram

```
┌─────────────────┐    ┌──────────────┐    ┌─────────────────┐
│   HTTP Layer    │───▶│  Handlers    │───▶│   Repository    │
│ (Gin Framework) │    │              │    │   (DB Models)   │
└─────────────────┘    └──────────────┘    └─────────────────┘
                                     │
                                     ▼
                              ┌─────────────────┐
                              │    Database     │
                              │    (SQLite)     │
                              └─────────────────┘
```

## Component Details

### 1. HTTP Layer (`cmd/api/`)

The entry point of the application. Handles HTTP requests and responses using Gin.

| File | Purpose |
|------|---------|
| `main.go` | Application initialization: loads env, opens DB, starts server |
| `routes.go` | All route definitions under `/api/v1` group |
| `server.go` | HTTP server configuration (timeouts, port binding) |
| `auth.go` | User registration handler with bcrypt password hashing |
| `events.go` | Event CRUD handlers and attendee management handlers |
| `users.go` | User listing and deletion handlers |

### 2. Repository Layer (`internal/database/`)

Contains data models and all database operations. Each entity has its own model file.

| File | Purpose |
|------|---------|
| `models.go` | `Models` struct container and constructor (`NewModels`) |
| `users.go` | `User` struct + `UserModel` (Insert, Get, GetAll, Delete) |
| `events.go` | `Event` struct + `EventModel` (Insert, Get, GetAll, Update, Delete, DeleteAll) |
| `attendes.go` | `Attendee` struct + `AttendeeModel` (Insert, GetByEventAndAttendee, GetAttendeesByEvent, Delete, GetEventsByAttendee) |

**Struct definitions:**

```go
type User struct {
    Id       int    `json:"id"`
    Email    string `json:"email"`
    Name     string `json:"name"`
    Password string `json:"-"`
}

type Event struct {
    Id          int    `json:"id"`
    OwnerId     int    `json:"ownerId"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Date        string `json:"date"`
    Location    string `json:"location"`
}

type Attendee struct {
    Id      int `json:"id"`
    UserId  int `json:"userId"`
    EventId int `json:"eventId"`
}
```

### 3. Environment Configuration (`internal/env/`)

| File | Purpose |
|------|---------|
| `env.go` | `GetEnvString(key, fallback)` and `GetEnvInt(key, fallback)` helpers |

Environment variables are loaded from `.env` at startup via `godotenv/autoload`.

### 4. Database Migrations (`cmd/migrate/`)

| File/Dir | Purpose |
|----------|---------|
| `main.go` | CLI migration runner (`up` to apply, `down` to rollback) |
| `migrations/` | SQL migration files (up/down pairs) |

**Schema:**

| Table | Columns | Notes |
|-------|---------|-------|
| `users` | id, email, name, password | `email` is UNIQUE |
| `events` | id, owner_id, name, description, date, location | `owner_id` FK → `users(id)` ON DELETE CASCADE |
| `attendees` | id, user_id, event_id | Both FKs ON DELETE CASCADE |

## Data Flow

1. Client sends HTTP request to Gin router
2. Router directs request to the appropriate handler in `cmd/api/`
3. Handler extracts data from request (body, params)
4. Handler calls the appropriate repository method in `internal/database/`
5. Repository executes SQL query against SQLite (with 3-second context timeout)
6. Results are returned back up the chain as JSON

## Key Design Decisions

### SQLite Database

- Chosen for simplicity and zero-configuration requirements
- File-based database stored as `data.db`
- Suitable for development and small-scale deployments
- Uses `$1`, `$2` style parameter placeholders (PostgreSQL-style, supported by the `go-sqlite3` driver)

### Gin Framework

- Selected for performance and minimalistic design
- Provides routing, JSON binding, and validation out of the box
- `ShouldBindJSON` with struct tags handles input validation

### Password Hashing

- Passwords are hashed with `bcrypt` (default cost) before storage
- The `Password` field has `json:"-"` tag to exclude it from all JSON responses

### Context Timeouts

- All database operations use a 3-second context timeout to prevent hanging queries

## Extensibility Points

### Adding New Features

1. Add a new model struct and methods in `internal/database/`
2. Update the `Models` struct in `models.go` to include the new model
3. Add handler functions in `cmd/api/` (e.g., `newfeature.go`)
4. Register new routes in `routes.go`
5. Create migration files in `cmd/migrate/migrations/`

### Switching Databases

- Database operations are encapsulated in model structs
- To switch databases, update the SQL queries in model methods and the connection initialization in `main.go`

### Adding Middleware

- Gin middleware can be added in `routes.go` via `g.Use(...)` or per-route
- Cross-cutting concerns like logging, auth, and validation can be implemented as middleware

## Known Issues

- **`attendes.go:83`** — The `Delete` method passes only `eventId` to `ExecContext`, but the SQL query expects two parameters (`user_id = $1 AND event_id = $2`). This will cause a runtime error when deleting attendees.
- **`events.go:27`** — The error message reads `"Failed To create error"` (typo).
- **`events.go:39`** and **`events.go:60`** — Error messages use `"retreive"` instead of `"retrieve"` (typo).
- **No authentication middleware** — The `JWT_SECRET` is loaded but not used. All endpoints are currently unauthenticated.
- **No tests** — The project has no test files.
