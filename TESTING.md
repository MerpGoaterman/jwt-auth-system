# Testing Guide

Complete guide for testing the JWT Authentication System locally and in production.

## Local Testing

### Prerequisites

- Go 1.21+
- Node.js 16+
- MySQL 8.0+

### Test Backend Locally

```bash
cd backend

# Set environment variables
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=root
export DB_PASSWORD=your-mysql-password
export DB_NAME=jwt_auth
export JWT_SECRET=test-secret-key
export PORT=8080

# Create database
mysql -u root -p << EOF
CREATE DATABASE IF NOT EXISTS jwt_auth;
USE jwt_auth;
source init.sql;
EOF

# Install dependencies
/usr/local/go/bin/go mod download

# Run backend
/usr/local/go/bin/go run main.go
```

Backend should start on http://localhost:8080

### Test Frontend Locally

```bash
cd frontend

# Set environment variable
echo "REACT_APP_API_URL=http://localhost:8080" > .env

# Install dependencies
npm install

# Start development server
npm start
```

Frontend should open on http://localhost:3000

## API Testing

### 1. Health Check

```bash
curl http://localhost:8080/health
```

Expected response:
```json
{
  "status": "healthy"
}
```

### 2. Login

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "admin123"
  }'
```

Expected response:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "uuid-here",
    "name": "Admin User",
    "email": "admin@example.com",
    "tenant_id": "default-tenant",
    "role": "admin",
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-01-01T00:00:00Z"
  }
}
```

### 3. Get All Users (Authenticated)

```bash
# Save token from login response
export TOKEN="your-jwt-token-here"

curl http://localhost:8080/api/users \
  -H "Authorization: Bearer $TOKEN"
```

Expected response:
```json
[
  {
    "id": "uuid-1",
    "name": "Admin User",
    "email": "admin@example.com",
    "tenant_id": "default-tenant",
    "role": "admin",
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-01-01T00:00:00Z"
  },
  {
    "id": "uuid-2",
    "name": "Regular User",
    "email": "user@example.com",
    "tenant_id": "default-tenant",
    "role": "user",
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-01-01T00:00:00Z"
  }
]
```

### 4. Create User (Admin Only)

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "password": "password123",
    "tenant_id": "tenant-001",
    "role": "user"
  }'
```

Expected response:
```json
{
  "id": "new-uuid",
  "name": "John Doe",
  "email": "john@example.com",
  "tenant_id": "tenant-001",
  "role": "user",
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-01T00:00:00Z"
}
```

### 5. Get User by ID

```bash
curl http://localhost:8080/api/users/USER_ID \
  -H "Authorization: Bearer $TOKEN"
```

### 6. Update User (Admin Only)

```bash
curl -X PUT http://localhost:8080/api/users/USER_ID \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe Updated",
    "email": "john.updated@example.com",
    "tenant_id": "tenant-002",
    "role": "admin"
  }'
```

### 7. Delete User (Admin Only)

```bash
curl -X DELETE http://localhost:8080/api/users/USER_ID \
  -H "Authorization: Bearer $TOKEN"
```

Expected response:
```json
{
  "message": "User deleted successfully"
}
```

## Error Cases

### 1. Login with Wrong Password

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "wrongpassword"
  }'
```

Expected: 401 Unauthorized

### 2. Access Protected Route Without Token

```bash
curl http://localhost:8080/api/users
```

Expected: 401 Unauthorized

### 3. Access Protected Route With Invalid Token

```bash
curl http://localhost:8080/api/users \
  -H "Authorization: Bearer invalid-token"
```

Expected: 401 Unauthorized

### 4. Non-Admin Tries to Create User

```bash
# Login as regular user first
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "user123"
  }'

# Use the token to try creating a user
curl -X POST http://localhost:8080/api/users \
  -H "Authorization: Bearer USER_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test User",
    "email": "test@example.com",
    "password": "password123",
    "tenant_id": "tenant-001",
    "role": "user"
  }'
```

Expected: 403 Forbidden

### 5. Create User with Duplicate Email

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Duplicate",
    "email": "admin@example.com",
    "password": "password123",
    "tenant_id": "tenant-001",
    "role": "user"
  }'
```

Expected: 400 Bad Request

## Frontend Testing

### 1. Login Flow

1. Open http://localhost:3000
2. Should redirect to `/login`
3. Enter credentials:
   - Email: `admin@example.com`
   - Password: `admin123`
4. Click "Login"
5. Should redirect to `/dashboard`

### 2. Dashboard View

After login, verify:
- User information card shows correct details
- JWT token is displayed (truncated)
- Users table shows all users
- "Add User" button is visible (admin only)

### 3. Create User Flow (Admin)

1. Click "Add User" button
2. Fill in form:
   - Name: Test User
   - Email: test@example.com
   - Password: password123
   - Tenant ID: tenant-001
   - Role: user
3. Click "Create User"
4. Modal should close
5. New user should appear in users table

### 4. Delete User Flow (Admin)

1. Find a user in the table
2. Click "Delete" button
3. Confirm deletion
4. User should be removed from table

### 5. Logout Flow

1. Click "Logout" button
2. Should redirect to `/login`
3. Token should be cleared from localStorage

### 6. Protected Route

1. Logout if logged in
2. Try to access `/dashboard` directly
3. Should redirect to `/login`

## Production Testing

### Test Deployed Backend

```bash
# Replace with your Railway URL
export BACKEND_URL="https://jwt-auth-system-production.up.railway.app"

# Health check
curl $BACKEND_URL/health

# Login
curl -X POST $BACKEND_URL/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "admin123"
  }'

# Get users
export TOKEN="your-token-here"
curl $BACKEND_URL/api/users \
  -H "Authorization: Bearer $TOKEN"
```

### Test Deployed Frontend

1. Open your Vercel URL
2. Test login with admin credentials
3. Verify dashboard loads correctly
4. Test creating a user
5. Test deleting a user
6. Test logout

## JWT Token Testing

### Decode Token

Use https://jwt.io to decode and verify your token:

1. Copy token from dashboard or API response
2. Paste into jwt.io
3. Verify payload contains:
   - `user_id`
   - `email`
   - `tenant_id`
   - `role`
   - `exp` (expiration)

### Verify Token in Code

**Go:**
```go
package main

import (
    "fmt"
    "github.com/golang-jwt/jwt/v5"
)

func main() {
    tokenString := "your-token-here"
    secret := "your-super-secret-jwt-key-change-this-in-production"
    
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })
    
    if err != nil {
        fmt.Println("Invalid token:", err)
        return
    }
    
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        fmt.Println("User ID:", claims["user_id"])
        fmt.Println("Email:", claims["email"])
        fmt.Println("Tenant:", claims["tenant_id"])
        fmt.Println("Role:", claims["role"])
    }
}
```

**Node.js:**
```javascript
const jwt = require('jsonwebtoken');

const token = 'your-token-here';
const secret = 'your-super-secret-jwt-key-change-this-in-production';

try {
    const decoded = jwt.verify(token, secret);
    console.log('User ID:', decoded.user_id);
    console.log('Email:', decoded.email);
    console.log('Tenant:', decoded.tenant_id);
    console.log('Role:', decoded.role);
} catch(err) {
    console.log('Invalid token:', err.message);
}
```

**Python:**
```python
import jwt

token = 'your-token-here'
secret = 'your-super-secret-jwt-key-change-this-in-production'

try:
    decoded = jwt.decode(token, secret, algorithms=['HS256'])
    print('User ID:', decoded['user_id'])
    print('Email:', decoded['email'])
    print('Tenant:', decoded['tenant_id'])
    print('Role:', decoded['role'])
except jwt.InvalidTokenError as e:
    print('Invalid token:', str(e))
```

## Load Testing

### Using Apache Bench

```bash
# Install apache2-utils
sudo apt-get install apache2-utils

# Test login endpoint
ab -n 1000 -c 10 -p login.json -T application/json \
  http://localhost:8080/api/auth/login

# login.json content:
# {"email":"admin@example.com","password":"admin123"}
```

### Using wrk

```bash
# Install wrk
sudo apt-get install wrk

# Test authenticated endpoint
wrk -t4 -c100 -d30s \
  -H "Authorization: Bearer YOUR_TOKEN" \
  http://localhost:8080/api/users
```

## Security Testing

### 1. SQL Injection

Try SQL injection in login:
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com OR 1=1--",
    "password": "anything"
  }'
```

Should fail (protected by prepared statements)

### 2. XSS

Try XSS in user creation:
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "<script>alert(\"XSS\")</script>",
    "email": "xss@example.com",
    "password": "password123",
    "tenant_id": "tenant-001",
    "role": "user"
  }'
```

Should be sanitized in frontend

### 3. Token Expiration

Wait 24 hours after login, then try to use the token:
```bash
curl http://localhost:8080/api/users \
  -H "Authorization: Bearer $EXPIRED_TOKEN"
```

Should return 401 Unauthorized

## Automated Testing

### Backend Unit Tests

```bash
cd backend
/usr/local/go/bin/go test ./... -v
```

### Frontend Tests

```bash
cd frontend
npm test
```

## Monitoring

### Check Logs

**Railway:**
1. Go to your service in Railway dashboard
2. Click "Logs" tab
3. Monitor for errors

**Vercel:**
1. Go to your deployment in Vercel dashboard
2. Click "Logs" tab
3. Monitor for errors

### Health Check Monitoring

Set up a cron job to check health:
```bash
*/5 * * * * curl -f https://your-backend.railway.app/health || echo "Backend down!"
```

## Troubleshooting

### Backend Issues

**Database connection fails:**
- Check environment variables
- Verify MySQL is running
- Check network connectivity

**JWT errors:**
- Verify JWT_SECRET is set
- Check token format
- Verify token hasn't expired

### Frontend Issues

**Can't connect to backend:**
- Check REACT_APP_API_URL
- Verify CORS is enabled
- Check network tab in browser

**Login fails:**
- Check credentials
- Verify backend is running
- Check browser console for errors

## Test Checklist

- [ ] Backend health check works
- [ ] Login with valid credentials succeeds
- [ ] Login with invalid credentials fails
- [ ] Protected routes require authentication
- [ ] Admin can create users
- [ ] Admin can delete users
- [ ] Non-admin cannot create users
- [ ] Non-admin cannot delete users
- [ ] JWT token contains correct claims
- [ ] JWT token expires after 24 hours
- [ ] Frontend redirects to login when not authenticated
- [ ] Frontend shows dashboard after login
- [ ] Frontend can create users (admin)
- [ ] Frontend can delete users (admin)
- [ ] Logout clears token and redirects

## Performance Benchmarks

Target metrics:
- Login: < 200ms
- Get users: < 100ms
- Create user: < 300ms
- JWT verification: < 10ms

## Next Steps

1. Set up automated testing in CI/CD
2. Add integration tests
3. Set up monitoring and alerting
4. Add performance testing to CI/CD
5. Set up security scanning

---

Happy Testing! 🧪
