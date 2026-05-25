# System Architecture

## Overview

This REST API is built using Go and the Gin web framework. It follows a simplified layered architecture separating concerns between different components.

## Architecture Diagram

```
┌─────────────────┐    ┌──────────────┐    ┌────────────────┐
│   HTTP Layer    │───▶│  Handlers    │───▶│ Repository     │
│ (Gin Framework) │    │              │    │   (DB Models)  │
└─────────────────┘    └──────────────┘    └────────────────┘
                                    │
                                    ▼
                           ┌────────────────┐
                           │   Database     │
                           │ (SQLite)       │
                           └────────────────┘
```

## Component Details

### 1. HTTP Layer (`cmd/api/`)
- Entry point of the application
- Handles HTTP requests and responses
- Uses Gin framework for routing and middleware
- Located in `cmd/api/`
  - `main.go`: Application initialization and server setup
  - `routes.go`: Route definitions
  - `server.go`: HTTP server configuration
  - `auth.go`: Authentication handlers
  - `events.go`: Event handlers

### 2. Repository Layer (`internal/database/`)
- Contains data models and database operations
- Implements CRUD operations for entities
- Located in `internal/database/`
  - `models.go`: Main models container
  - `users.go`: User entity and operations
  - `events.go`: Event entity and operations
  - `attendes.go`: Attendee entity and operations

### 3. Environment Configuration (`internal/env/`)
- Handles loading and accessing environment variables
- Provides helper functions for type-safe environment variable access
- Located in `internal/env/`
  - `env.go`: Environment variable utility functions

### 4. Database Migrations (`cmd/migrate/`)
- Manages database schema versioning
- Uses golang-migrate library
- Located in `cmd/migrate/`
  - `main.go`: Migration tool entry point
  - `migrations/`: SQL migration files

## Data Flow

1. Client sends HTTP request to Gin router
2. Router directs request to appropriate handler (in `cmd/api/`)
3. Handler extracts data from request (body, params, query)
4. Handler calls appropriate repository method (in `internal/database/`)
5. Repository executes SQL query against SQLite database
6. Results are returned back up the chain to the client

## Key Design Decisions

### SQLite Database
- Chosen for simplicity and zero-configuration requirements
- Suitable for development and small-scale deployments
- File-based database stored as `data.db`

### Gin Framework
- Selected for performance and minimalistic design
- Provides essential features without unnecessary bloat
- Good middleware support for extensibility

### Structured Error Handling
- Errors are propagated up the call stack
- Handled at the HTTP layer for appropriate responses
- Logging implemented for debugging and monitoring

### Environment Configuration
- Centralized environment variable access
- Type-safe parsing with fallback defaults
- Supports .env file via godotenv for local development

## Extensibility Points

### Adding New Features
1. Add new model structs in `internal/database/`
2. Implement corresponding database operations (CRUD methods)
3. Update `Models` struct in `models.go` to include new model
4. Add handlers in `cmd/api/` (e.g., `newfeature.go`)
5. Register new routes in `routes.go`

### Switching Databases
1. Database operations are encapsulated in model structs
2. To switch databases, only the SQL queries in model methods need updating
3. Connection initialization in `main.go` would need modification

### Adding Middleware
1. Gin middleware can be added in `routes.go`
2. Cross-cutting concerns like logging, auth, validation can be implemented as middleware

## Deployment Considerations

### Environment Variables
- `PORT`: Server port (default: 8080)
- `JWT_SECRET`: Secret for JWT token generation
- Other variables as needed for external services

### Production Readiness
- Consider adding connection pooling for database
- Implement proper logging rotation
- Add rate limiting and security headers
- Consider using a more robust database (PostgreSQL/MySQL) for production