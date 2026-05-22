# API Documentation

## Base URL

```
http://localhost:8080/api/v1
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

### Events

#### Create a New Event
```
POST /api/v1/events
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
GET /api/v1/events
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
GET /api/v1/events/:id
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
PUT /api/v1/events/:id
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
DELETE /api/v1/events/:id
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

Example: `/api/v1/events?page=2&limit=5`

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
| 500 | Internal Server Error