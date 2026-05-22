# API Documentation

## Base URL

```
http://localhost:8080
```

## Authentication

This API uses JWT (JSON Web Token) for authentication. Include the token in the Authorization header:

```
Authorization: Bearer <your_jwt_token>
```

## Response Format

All API responses follow this format:

### Success Response
```json
{
  "success": true,
  "data": { ... }
}
```

### Error Response
```json
{
  "success": false,
  "error": {
    "message": "Error description",
    "code": "ERROR_CODE"
  }
}
```

## Endpoints

### Users

#### Register a New User
```
POST /api/users
```

**Request Body:**
```json
{
  "email": "user@example.com",
  "name": "John Doe",
  "password": "securepassword123"
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe"
  }
}
```

#### Login User
```
POST /api/auth/login
```

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword123"
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

#### Get User Profile
```
GET /api/users/profile
```

**Headers:**
```
Authorization: Bearer <your_jwt_token>
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe"
  }
}
```

### Events

#### Create a New Event
```
POST /api/events
```

**Headers:**
```
Authorization: Bearer <your_jwt_token>
```

**Request Body:**
```json
{
  "name": "Tech Conference 2026",
  "description": "Annual technology conference featuring latest innovations",
  "date": "2026-12-15",
  "location": "Convention Center, New York"
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "ownerId": 1,
    "name": "Tech Conference 2026",
    "description": "Annual technology conference featuring latest innovations",
    "date": "2026-12-15",
    "location": "Convention Center, New York"
  }
}
```

#### Get All Events
```
GET /api/events
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "ownerId": 1,
      "name": "Tech Conference 2026",
      "description": "Annual technology conference featuring latest innovations",
      "date": "2026-12-15",
      "location": "Convention Center, New York"
    }
  ]
}
```

#### Get Event by ID
```
GET /api/events/:id
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "ownerId": 1,
    "name": "Tech Conference 2026",
    "description": "Annual technology conference featuring latest innovations",
    "date": "2026-12-15",
    "location": "Convention Center, New York"
  }
}
```

#### Update Event
```
PUT /api/events/:id
```

**Headers:**
```
Authorization: Bearer <your_jwt_token>
```

**Request Body:**
```json
{
  "name": "Updated Event Name",
  "description": "Updated description",
  "date": "2026-12-20",
  "location": "Updated Location"
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "ownerId": 1,
    "name": "Updated Event Name",
    "description": "Updated description",
    "date": "2026-12-20",
    "location": "Updated Location"
  }
}
```

#### Delete Event
```
DELETE /api/events/:id
```

**Headers:**
```
Authorization: Bearer <your_jwt_token>
```

**Response:**
```json
{
  "success": true,
  "data": {
    "message": "Event deleted successfully"
  }
}
```

### Attendance

#### Register for an Event
```
POST /api/events/:id/attend
```

**Headers:**
```
Authorization: Bearer <your_jwt_token>
```

**Response:**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "userId": 1,
    "eventId": 1
  }
}
```

#### Cancel Attendance
```
DELETE /api/events/:id/attend
```

**Headers:**
```
Authorization: Bearer <your_jwt_token>
```

**Response:**
```json
{
  "success": true,
  "data": {
    "message": "Attendance cancelled successfully"
  }
}
```

#### Get Event Attendees
```
GET /api/events/:id/attendees
```

**Response:**
```json
{
  "success": true,
  "data": [
    {
      "id": 1,
      "userId": 1,
      "eventId": 1
    }
  ]
}
```

## Error Codes

| Code | Message | Description |
|------|---------|-------------|
| VALIDATION_ERROR | Validation failed | Request data failed validation |
| AUTHENTICATION_ERROR | Authentication failed | Invalid or missing credentials |
| AUTHORIZATION_ERROR | Authorization failed | Insufficient permissions |
| NOT_FOUND | Resource not found | Requested resource doesn't exist |
| INTERNAL_ERROR | Internal server error | Unexpected server error |
| CONFLICT | Resource conflict | Resource already exists or conflicting state |

## Rate Limiting

API endpoints are subject to rate limiting:
- 100 requests per minute per IP address
- Exceeding limits returns HTTP 429 (Too Many Requests)

## Pagination

Endpoints that return lists support pagination:
- `page`: Page number (default: 1)
- `limit`: Items per page (default: 10)

Example: `/api/events?page=2&limit=5`

## HTTP Status Codes

| Code | Description |
|------|-------------|
| 200 | OK |
| 201 | Created |
| 400 | Bad Request |
| 401 | Unauthorized |
| 403 | Forbidden |
| 404 | Not Found |
| 409 | Conflict |
| 422 | Unprocessable Entity |
| 429 | Too Many Requests |
| 500 | Internal Server Error |