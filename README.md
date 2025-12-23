# JWT Authentication System

A production-ready JWT authentication system with user, tenant, and role management. Built with Go backend, React frontend, Railway MySQL database, and deployed on Railway and Vercel.

## Features

- **JWT-based Authentication**: Secure token-based authentication with 24-hour expiration
- **Multi-tenant Support**: Tenant isolation for user management
- **Role-based Access Control**: Admin, user, and custom roles
- **User Management**: Create, read, update, and delete users
- **Cross-application Authentication**: JWT tokens can be used across multiple applications
- **Production-ready**: Deployed on Railway (backend) and Vercel (frontend)

## Architecture

### Backend (Go)
- RESTful API with JWT authentication
- MySQL database on Railway
- Password hashing with SHA-256
- CORS-enabled for frontend communication
- Environment-based configuration

### Frontend (React + TypeScript)
- Modern React with TypeScript
- Context-based authentication state management
- Protected routes
- User management interface
- Deployed on Vercel

### Database (Railway MySQL)
- Users table with tenant and role support
- Secure password storage
- Automatic timestamps

## Tech Stack

| Component | Technology |
|-----------|-----------|
| Backend | Go 1.21+ |
| Frontend | React 18 + TypeScript |
| Database | MySQL (Railway) |
| JWT Library | github.com/golang-jwt/jwt/v5 |
| Password Hashing | crypto/sha256 |
| Backend Deployment | Railway |
| Frontend Deployment | Vercel |

## API Endpoints

### Public Endpoints
- `POST /api/auth/login` - Login and receive JWT token

### Protected Endpoints (require JWT token)
- `GET /api/users` - List all users
- `POST /api/users` - Create new user (admin only)
- `GET /api/users/:id` - Get user by ID
- `PUT /api/users/:id` - Update user
- `DELETE /api/users/:id` - Delete user (admin only)
- `GET /api/auth/me` - Get current user info

## JWT Token Structure

```json
{
  "user_id": "uuid",
  "email": "user@example.com",
  "tenant_id": "tenant-123",
  "role": "admin",
  "exp": 1234567890
}
```

## Quick Start

### Backend Setup

```bash
cd backend
go mod init github.com/MerpGoaterman/jwt-auth-system/backend
go mod tidy
go run main.go
```

### Frontend Setup

```bash
cd frontend
npm install
npm start
```

## Environment Variables

### Backend
```
DB_HOST=<railway-mysql-host>
DB_PORT=3306
DB_USER=root
DB_PASSWORD=<railway-mysql-password>
DB_NAME=railway
JWT_SECRET=<your-secret-key>
PORT=8080
```

### Frontend
```
REACT_APP_API_URL=<backend-url>
```

## Deployment

### Backend (Railway)
1. Connect GitHub repository to Railway
2. Set environment variables
3. Deploy automatically on push

### Frontend (Vercel)
1. Connect GitHub repository to Vercel
2. Set environment variables
3. Deploy automatically on push

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

## Usage Example

### Login
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"password123"}'
```

Response:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "user-1",
    "email": "admin@example.com",
    "tenant_id": "tenant-1",
    "role": "admin"
  }
}
```

### Create User (with JWT token)
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-jwt-token>" \
  -d '{
    "name":"John Doe",
    "email":"john@example.com",
    "password":"password123",
    "tenant_id":"tenant-1",
    "role":"user"
  }'
```

## Security Features

- Password hashing with SHA-256
- JWT token expiration (24 hours)
- Protected routes with middleware
- CORS configuration
- Environment-based secrets
- Role-based access control

## License

MIT

## Author

MerpGoaterman
