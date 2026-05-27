# REST API in Gin

A RESTful event management API built with Go and the Gin web framework. Learn Go by building a production-ready API with authentication, database operations, and proper architecture.

## 🎯 Features

- **User Management**
  - User registration with bcrypt password hashing
  - User listing and deletion
  - Secure password storage (excluded from API responses)

- **Event Management**
  - Full CRUD operations (Create, Read, Update, Delete)
  - Event ownership with authorization
  - Event filtering by attendees

- **Attendee Registration**
  - Register users for events
  - Prevent duplicate registrations
  - List attendees by event
  - Get all events for a user

- **Authentication & Authorization**
  - JWT-based authentication
  - Protected endpoints with middleware
  - Event ownership validation

- **Database**
  - SQLite with file-based persistence
  - golang-migrate schema migrations
  - Foreign key constraints with CASCADE delete

- **Developer Experience**
  - Live reload development with Air
  - RESTful API design with versioning (`/api/v1`)
  - Comprehensive API documentation
  - Well-structured codebase

## 📋 Project Structure

```
.
├── cmd/
│   ├── api/                  # HTTP API server
│   │   ├── main.go           # Application entry point
│   │   ├── server.go         # HTTP server configuration
│   │   ├── routes.go         # Route definitions
│   │   ├── middleware.go     # JWT authentication middleware
│   │   ├── context.go        # Context utilities
│   │   ├── auth.go           # User registration & login handlers
│   │   ├── events.go         # Event CRUD & attendee handlers
│   │   └── users.go          # User listing/deletion handlers
│   └── migrate/              # Database migration tool
│       ├── main.go           # Migration CLI runner
│       └── migrations/       # SQL migration files
├── internal/
│   ├── database/             # Database models & operations
│   │   ├── models.go         # Models container & constructor
│   │   ├── users.go          # User model & CRUD operations
│   │   ├── events.go         # Event model & CRUD operations
│   │   └── attendees.go      # Attendee model & CRUD operations
│   └── env/                  # Environment variable handling
│       └── env.go            # GetEnvString / GetEnvInt helpers
├── docs/                     # Documentation
│   ├── api.md                # API endpoint reference
│   └── architecture.md       # System architecture overview
├── .env                      # Environment variables
├── .air.toml                 # Air live-reload configuration
├── go.mod                    # Go module definition
├── go.sum                    # Go dependencies lock file
├── data.db                   # SQLite database (generated at runtime)
└── README.md                 # This file
```

## 🚀 Getting Started

### Prerequisites

- **Go** 1.26 or higher ([download](https://golang.org/dl/))
- **SQLite3** (usually pre-installed on macOS/Linux)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/Akene-Uzezi/go-rest-api.git
   cd go-rest-api
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run database migrations**
   ```bash
   go run ./cmd/migrate
   ```
   This creates the SQLite database and initializes the schema.

4. **Start the server**
   ```bash
   go run ./cmd/api
   ```
   The API will be available at `http://localhost:8080`.

### Environment Variables

Create a `.env` file in the project root with:

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Port to run the server on |
| `JWT_SECRET` | `secret` | Secret key for JWT token signing (change in production!) |

**Example `.env`:**
```env
PORT=8080
JWT_SECRET=your-secret-key-min-32-chars-recommended
```

### Live Reload (Development)

[Air](https://github.com/cosmtrek/air) is configured for hot-reload during development:

```bash
# Install Air (one-time)
go install github.com/cosmtrek/air@latest

# Start with live reload
air
```

The server rebuilds automatically when `.go` files change. Configuration is in `.air.toml`.

## 📚 API Documentation

Complete API endpoint documentation is available in [docs/api.md](docs/api.md).

### Quick Reference

#### Public Endpoints (No Authentication Required)

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/ping` | Health check |
| `GET` | `/api/v1/` | API status |
| `POST` | `/api/v1/auth/register` | Register a new user |
| `POST` | `/api/v1/auth/login` | Login & receive JWT token |
| `GET` | `/api/v1/users` | List all users |
| `GET` | `/api/v1/events` | List all events |
| `GET` | `/api/v1/events/:id` | Get event by ID |
| `GET` | `/api/v1/events/:id/attendees` | List event attendees |
| `GET` | `/api/v1/attendees/:id/events` | Get events attended by user |

#### Protected Endpoints (JWT Required)

| Method | Path | Description |
|--------|------|-------------|
| `POST` | `/api/v1/events` | Create a new event |
| `PUT` | `/api/v1/events/:id` | Update an event (owner only) |
| `DELETE` | `/api/v1/events/:id` | Delete an event (owner only) |
| `POST` | `/api/v1/events/:id/attendees/:userId` | Register attendee for event |
| `DELETE` | `/api/v1/events/:id/attendees/:userId` | Remove attendee from event |
| `DELETE` | `/api/v1/delete/:id` | Delete a user |
| `DELETE` | `/api/v1/events` | Delete all events |

### Authentication

Protected endpoints require a JWT token in the `Authorization` header:

```bash
curl -H "Authorization: Bearer <your-jwt-token>" \
     http://localhost:8080/api/v1/events
```

Obtain a token by logging in:
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "securepassword123"
  }'
```

## 🗄️ Database Schema

Migrations are managed using [golang-migrate](https://github.com/golang-migrate/migrate).

### Tables

**`users`**
- `id` (INTEGER PRIMARY KEY)
- `email` (TEXT UNIQUE NOT NULL)
- `name` (TEXT NOT NULL)
- `password` (TEXT NOT NULL) — bcrypt hashed

**`events`**
- `id` (INTEGER PRIMARY KEY)
- `owner_id` (INTEGER NOT NULL) — FK → `users(id)` ON DELETE CASCADE
- `name` (TEXT NOT NULL)
- `description` (TEXT NOT NULL)
- `date` (TEXT NOT NULL) — format: YYYY-MM-DD
- `location` (TEXT NOT NULL)

**`attendees`**
- `id` (INTEGER PRIMARY KEY)
- `user_id` (INTEGER NOT NULL) — FK → `users(id)` ON DELETE CASCADE
- `event_id` (INTEGER NOT NULL) — FK → `events(id)` ON DELETE CASCADE
- Unique constraint on `(user_id, event_id)`

### Migration Commands

```bash
# Apply all pending migrations
go run ./cmd/migrate

# Rollback the last migration
go run ./cmd/migrate down

# Create a new migration pair
migrate create -ext sql -dir cmd/migrate/migrations -seq <name>
```

## 🏗️ Architecture

The application follows a **layered architecture** separating concerns:

```
┌──────────────────┐      ┌──────────────┐      ┌──────────────────┐
│   HTTP Layer     │──→   │  Handlers    │──→   │   Repository     │
│ (Gin Framework)  │      │  (cmd/api)   │      │   (db models)    │
└──────────────────┘      └──────────────┘      └──────────────────┘
                                        │
                                        ▼
                               ┌──────────────────┐
                               │   Database       │
                               │   (SQLite)       │
                               └──────────────────┘
```

### Layer Details

**HTTP Layer (`cmd/api/`)**
- Handles HTTP requests/responses with Gin
- Input validation and binding
- Error handling and response formatting
- Authentication middleware

**Handler Layer (`cmd/api/`)**
- `auth.go` — User registration & login
- `events.go` — Event CRUD & attendee management
- `users.go` — User listing & deletion
- `middleware.go` — JWT authentication & context injection
- `context.go` — Utility functions

**Repository Layer (`internal/database/`)**
- Entity models (User, Event, Attendee)
- SQL query execution
- Database context management

**Database Layer**
- SQLite persistence
- Schema migrations

For detailed architecture documentation, see [docs/architecture.md](docs/architecture.md).

## 💡 Usage Examples

### 1. Register a User

```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "alice@example.com",
    "password": "securepass123",
    "name": "Alice"
  }'
```

Response:
```json
{
  "id": 1,
  "email": "alice@example.com",
  "name": "Alice"
}
```

### 2. Login

```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "alice@example.com",
    "password": "securepass123"
  }'
```

Response:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### 3. Create an Event (Protected)

```bash
curl -X POST http://localhost:8080/api/v1/events \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -d '{
    "ownerId": 1,
    "name": "Tech Conference 2026",
    "description": "Annual technology conference featuring latest innovations",
    "date": "2026-12-15",
    "location": "Convention Center, New York"
  }'
```

Response:
```json
{
  "id": 1,
  "ownerId": 1,
  "name": "Tech Conference 2026",
  "description": "Annual technology conference featuring latest innovations",
  "date": "2026-12-15",
  "location": "Convention Center, New York"
}
```

### 4. List All Events

```bash
curl http://localhost:8080/api/v1/events
```

### 5. Register for an Event

```bash
curl -X POST http://localhost:8080/api/v1/events/1/attendees/2 \
  -H "Authorization: Bearer <token>"
```

## 🛠️ Technology Stack

- **Language**: Go 1.26+
- **Web Framework**: [Gin](https://github.com/gin-gonic/gin)
- **Database**: SQLite
- **Authentication**: JWT (golang-jwt)
- **Password Hashing**: bcrypt
- **Migrations**: golang-migrate
- **Environment**: godotenv
- **Development**: Air (live reload)

## 📝 Code Structure & Conventions

- **Layered Architecture**: HTTP → Handlers → Repository → Database
- **Error Handling**: Explicit error returns with meaningful messages
- **Validation**: Struct tags with Gin's `ShouldBindJSON` validator
- **Context**: Request context with 3-second timeout for DB operations
- **Security**: 
  - Passwords hashed with bcrypt
  - JWT tokens in Authorization header
  - Ownership validation for sensitive operations
  - Password excluded from JSON responses (`json:"-"`)

## 📖 Learning Resources

This project is designed to teach Go best practices:

1. **Gin Framework**: Route handling, middleware, validation
2. **Database Access**: SQL, migrations, connection pooling
3. **Authentication**: JWT tokens, password hashing
4. **API Design**: RESTful principles, error handling
5. **Go Idioms**: Error handling, context usage, struct composition
6. **Concurrency**: sync/errgroup for parallel operations

## ⚠️ Known Limitations

This is a learning project. In production, consider:

- [ ] Add comprehensive error logging
- [ ] Add request/response logging middleware
- [ ] Add rate limiting
- [ ] Add CORS configuration
- [ ] Add input sanitization
- [ ] Add unit and integration tests
- [ ] Use environment-specific configurations
- [ ] Add API versioning strategy
- [ ] Add database connection pooling configuration
- [ ] Use a proper secrets management system

## 🤝 Contributing

This is a personal learning project, but improvements are welcome!

## 📄 License

MIT License — feel free to use this for learning and reference.

## 📞 Support

For questions about Go and this API:
- [Go Documentation](https://golang.org/doc/)
- [Gin Documentation](https://gin-gonic.com/)
- [JWT-Go Documentation](https://pkg.go.dev/github.com/golang-jwt/jwt)
