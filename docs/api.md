# API Documentation

## Base URL

```
http://localhost:8080
```

## Health Check

```
GET /ping
```

**Response:** `200 OK`

```json
{
  "message": "PONG"
}
```

## API Status

```
GET /api/v1/
```

**Response:** `200 OK`

```json
{
  "message": "The api is running fine"
}
```

---

## Users

### Register a New User

```
POST /api/v1/auth/register
```

**Request Body:**

```json
{
  "email": "user@example.com",
  "password": "securepassword123",
  "name": "John Doe"
}
```

**Validation rules:**
- `email`: required
- `password`: required, minimum 8 characters
- `name`: required, minimum 2 characters

**Response:** `201 Created`

```json
{
  "id": 1,
  "email": "user@example.com",
  "name": "John Doe"
}
```

### List All Users

```
GET /api/v1/users
```

**Response:** `200 OK`

```json
[
  {
    "id": 1,
    "email": "user@example.com",
    "name": "John Doe"
  }
]
```

### Delete a User

```
DELETE /api/v1/delete/:id
```

**Response:** `204 No Content`

---

## Events

### Create an Event

```
POST /api/v1/events
```

**Request Body:**

```json
{
  "ownerId": 1,
  "name": "Tech Conference 2026",
  "description": "Annual technology conference featuring latest innovations",
  "date": "2026-12-15",
  "location": "Convention Center, New York"
}
```

**Validation rules:**
- `ownerId`: required (integer)
- `name`: required, minimum 3 characters
- `description`: required, minimum 10 characters
- `date`: required, format `YYYY-MM-DD`
- `location`: required, minimum 3 characters

**Response:** `201 Created`

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

### List All Events

```
GET /api/v1/events
```

**Response:** `200 OK`

```json
[
  {
    "id": 1,
    "ownerId": 1,
    "name": "Tech Conference 2026",
    "description": "Annual technology conference featuring latest innovations",
    "date": "2026-12-15",
    "location": "Convention Center, New York"
  }
]
```

### Get Event by ID

```
GET /api/v1/events/:id
```

**Response:** `200 OK`

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

### Update an Event

```
PUT /api/v1/events/:id
```

**Request Body:**

```json
{
  "ownerId": 1,
  "name": "Updated Event Name",
  "description": "Updated description text",
  "date": "2026-12-20",
  "location": "Updated Location"
}
```

**Response:** `200 OK`

```json
{
  "id": 1,
  "ownerId": 1,
  "name": "Updated Event Name",
  "description": "Updated description text",
  "date": "2026-12-20",
  "location": "Updated Location"
}
```

### Delete an Event

```
DELETE /api/v1/events/:id
```

**Response:** `204 No Content`

### Delete All Events

```
DELETE /api/v1/events
```

**Response:** `204 No Content`

---

## Attendees

### Add Attendee to Event

```
POST /api/v1/events/:id/attendees/:userId
```

Checks that both the event and user exist, and prevents duplicate registrations.

**Response:** `201 Created`

```json
{
  "id": 1,
  "userId": 2,
  "eventId": 1
}
```

### List Attendees for Event

```
GET /api/v1/events/:id/attendees
```

Returns the list of users registered as attendees for the given event.

**Response:** `200 OK`

```json
[
  {
    "id": 2,
    "name": "Jane Smith",
    "email": "jane@example.com"
  }
]
```

### Remove Attendee from Event

```
DELETE /api/v1/events/:id/attendees/:userId
```

**Response:** `204 No Content`

### Get Events by Attendee

```
GET /api/v1/attendees/:id/events
```

Returns all events that a specific user is attending.

**Response:** `200 OK`

```json
[
  {
    "id": 1,
    "ownerId": 1,
    "name": "Tech Conference 2026",
    "description": "Annual technology conference",
    "date": "2026-12-15",
    "location": "Convention Center, New York"
  }
]
```

---

## Error Responses

Errors are returned as JSON with an `error` field:

```json
{
  "error": "Error description"
}
```

### HTTP Status Codes

| Code | Description | When |
|------|-------------|------|
| 200 | OK | Successful GET/PUT request |
| 201 | Created | Successful POST request |
| 204 | No Content | Successful DELETE request |
| 400 | Bad Request | Invalid request body or parameters |
| 404 | Not Found | Resource doesn't exist |
| 409 | Conflict | Attendee already registered for event |
| 500 | Internal Server Error | Unexpected server error |
