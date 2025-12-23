# JWT Authentication Backend (Go)

RESTful API backend for JWT authentication with user, tenant, and role management.

## Features

- JWT-based authentication
- User management (CRUD operations)
- Tenant and role support
- Password hashing with SHA-256
- MySQL database integration
- CORS-enabled
- Protected routes with middleware

## Tech Stack

- Go 1.21+
- MySQL (Railway)
- JWT (github.com/golang-jwt/jwt/v5)
- Gorilla Mux (routing)
- CORS (github.com/rs/cors)

## API Endpoints

### Public Endpoints

#### Login
```
POST /api/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}

Response:
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "uuid",
    "name": "User Name",
    "email": "user@example.com",
    "tenant_id": "tenant-1",
    "role": "admin",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

#### Health Check
```
GET /health

Response:
{
  "status": "healthy"
}
```

### Protected Endpoints (require JWT token in Authorization header)

#### Get All Users
```
GET /api/users
Authorization: Bearer <token>

Response:
[
  {
    "id": "uuid",
    "name": "User Name",
    "email": "user@example.com",
    "tenant_id": "tenant-1",
    "role": "admin",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
]
```

#### Create User (Admin only)
```
POST /api/users
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "New User",
  "email": "newuser@example.com",
  "password": "password123",
  "tenant_id": "tenant-1",
  "role": "user"
}

Response:
{
  "id": "uuid",
  "name": "New User",
  "email": "newuser@example.com",
  "tenant_id": "tenant-1",
  "role": "user"
}
```

#### Get User by ID
```
GET /api/users/:id
Authorization: Bearer <token>

Response:
{
  "id": "uuid",
  "name": "User Name",
  "email": "user@example.com",
  "tenant_id": "tenant-1",
  "role": "admin",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

#### Update User
```
PUT /api/users/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "name": "Updated Name",
  "email": "updated@example.com",
  "tenant_id": "tenant-1",
  "role": "user"
}

Response:
{
  "id": "uuid",
  "name": "Updated Name",
  "email": "updated@example.com",
  "tenant_id": "tenant-1",
  "role": "user"
}
```

#### Delete User (Admin only)
```
DELETE /api/users/:id
Authorization: Bearer <token>

Response: 204 No Content
```

#### Get Current User
```
GET /api/auth/me
Authorization: Bearer <token>

Response:
{
  "id": "uuid",
  "name": "User Name",
  "email": "user@example.com",
  "tenant_id": "tenant-1",
  "role": "admin",
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

## Environment Variables

Create a `.env` file based on `.env.example`:

```
MYSQLHOST=your-railway-mysql-host
MYSQLPORT=3306
MYSQLUSER=root
MYSQLPASSWORD=your-railway-mysql-password
MYSQLDATABASE=railway
JWT_SECRET=your-secret-key
PORT=8080
```

## Running Locally

1. Install Go 1.21 or higher
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Set environment variables (or create .env file)
4. Run the server:
   ```bash
   go run main.go
   ```

## Building

```bash
go build -o jwt-auth-backend
./jwt-auth-backend
```

## Docker

Build:
```bash
docker build -t jwt-auth-backend .
```

Run:
```bash
docker run -p 8080:8080 \
  -e MYSQLHOST=your-host \
  -e MYSQLPORT=3306 \
  -e MYSQLUSER=root \
  -e MYSQLPASSWORD=your-password \
  -e MYSQLDATABASE=railway \
  -e JWT_SECRET=your-secret \
  jwt-auth-backend
```

## Deployment to Railway

1. Connect GitHub repository to Railway
2. Set environment variables in Railway dashboard
3. Railway will automatically detect Dockerfile and deploy

## Database Schema

```sql
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    tenant_id VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

## JWT Token Structure

```json
{
  "user_id": "uuid",
  "email": "user@example.com",
  "tenant_id": "tenant-1",
  "role": "admin",
  "exp": 1234567890
}
```

Token expires after 24 hours.

## Project Structure

```
backend/
├── main.go              # Main application entry point
├── auth/
│   └── auth.go         # JWT generation and validation
├── handlers/
│   └── handlers.go     # HTTP request handlers
├── go.mod              # Go module definition
├── go.sum              # Go module checksums
├── Dockerfile          # Docker configuration
├── .env.example        # Example environment variables
└── README.md           # This file
```

## Security Features

- Password hashing with SHA-256
- JWT token expiration (24 hours)
- Protected routes with middleware
- Role-based access control (admin-only endpoints)
- CORS configuration
- Environment-based secrets

## License

MIT
