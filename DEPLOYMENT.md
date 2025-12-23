# Deployment Guide

Complete guide for deploying the JWT Authentication System to Railway (backend) and Vercel (frontend).

## Prerequisites

- Railway account (https://railway.app)
- Vercel account (https://vercel.com)
- GitHub repository: https://github.com/MerpGoaterman/jwt-auth-system

## Railway MySQL Database

✅ **Already Created!**

Your Railway MySQL database is running with the following credentials:

```
MYSQL_DATABASE=railway
MYSQL_ROOT_PASSWORD=DvWLsMWwNaP0NlHMkkOKHTBTXvErMBTI
MYSQLUSER=root
MYSQLPORT=3306
```

## Deploy Backend to Railway

### Option 1: Railway Dashboard (Recommended)

1. Go to https://railway.app/project/440af528-f60c-40f3-a238-cd9a33a2f922
2. Click "Create" → "GitHub Repo"
3. Select `MerpGoaterman/jwt-auth-system`
4. Railway will detect the Dockerfile in the `backend` folder
5. Configure the following environment variables in Railway:
   ```
   DB_HOST=${MYSQLHOST}
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=${MYSQLPASSWORD}
   DB_NAME=railway
   JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
   PORT=8080
   ```
6. Set the **Root Directory** to `backend` in service settings
7. Railway will automatically deploy the backend

### Option 2: Railway CLI

```bash
cd /home/ubuntu/jwt-auth-system/backend

# Login to Railway
railway login

# Link to your project
railway link 440af528-f60c-40f3-a238-cd9a33a2f922

# Set environment variables
railway variables set DB_HOST=\${MYSQLHOST}
railway variables set DB_PORT=3306
railway variables set DB_USER=root
railway variables set DB_PASSWORD=\${MYSQLPASSWORD}
railway variables set DB_NAME=railway
railway variables set JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
railway variables set PORT=8080

# Deploy
railway up
```

### Get Backend URL

After deployment, Railway will provide a public URL like:
```
https://jwt-auth-system-production.up.railway.app
```

You can find this in:
- Railway Dashboard → Your Service → Settings → Domains
- Or generate a domain: Settings → Networking → Generate Domain

## Initialize Database

Once the backend is deployed, you need to create the database tables and an initial admin user.

### Method 1: Connect via Railway MySQL Client

```bash
# Get connection string from Railway dashboard
# Variables tab → MYSQL_PUBLIC_URL

# Connect to MySQL
mysql -h <host> -P <port> -u root -p

# Enter password: DvWLsMWwNaP0NlHMkkOKHTBTXvErMBTI

# Create tables
USE railway;

CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    tenant_id VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_email (email),
    INDEX idx_tenant (tenant_id)
);

# Create admin user (password: admin123)
INSERT INTO users (id, name, email, password, tenant_id, role) 
VALUES (
    UUID(),
    'Admin User',
    'admin@example.com',
    '$2a$10$rqYvE5zGwZ8xH5xH5xH5xOqYvE5zGwZ8xH5xH5xH5xOqYvE5zGwZ8',
    'default-tenant',
    'admin'
);
```

### Method 2: Use API to Create Admin User

```bash
# After backend is deployed, use curl to create admin user
curl -X POST https://your-backend-url.railway.app/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Admin User",
    "email": "admin@example.com",
    "password": "admin123",
    "tenant_id": "default-tenant",
    "role": "admin"
  }'
```

## Deploy Frontend to Vercel

### Method 1: Vercel Dashboard (Recommended)

1. Go to https://vercel.com/new
2. Import Git Repository
3. Select `MerpGoaterman/jwt-auth-system`
4. Configure project:
   - **Framework Preset**: Create React App
   - **Root Directory**: `frontend`
   - **Build Command**: `npm run build`
   - **Output Directory**: `build`
5. Add environment variable:
   ```
   REACT_APP_API_URL=https://your-backend-url.railway.app
   ```
6. Click "Deploy"

### Method 2: Vercel CLI

```bash
cd /home/ubuntu/jwt-auth-system/frontend

# Install Vercel CLI
npm install -g vercel

# Login
vercel login

# Deploy
vercel --prod

# Set environment variable
vercel env add REACT_APP_API_URL production
# Enter: https://your-backend-url.railway.app
```

## Test the Application

1. **Frontend URL**: Your Vercel deployment URL (e.g., `https://jwt-auth-system.vercel.app`)
2. **Backend URL**: Your Railway deployment URL (e.g., `https://jwt-auth-system-production.up.railway.app`)

### Test Login

1. Open the frontend URL in your browser
2. Login with:
   - Email: `admin@example.com`
   - Password: `admin123`
3. You should see the dashboard with user management

### Test API Endpoints

```bash
# Health check
curl https://your-backend-url.railway.app/health

# Login
curl -X POST https://your-backend-url.railway.app/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "admin123"
  }'

# This will return a JWT token
```

### Test JWT Token

```bash
# Use the token from login response
export TOKEN="your-jwt-token-here"

# Get all users
curl https://your-backend-url.railway.app/api/users \
  -H "Authorization: Bearer $TOKEN"

# Create a new user
curl -X POST https://your-backend-url.railway.app/api/users \
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

## Using JWT Tokens in Other Applications

The JWT tokens returned by this system can be used to authenticate users in other applications.

### Token Structure

The JWT token contains:
```json
{
  "user_id": "uuid",
  "email": "user@example.com",
  "tenant_id": "tenant-id",
  "role": "user|admin",
  "exp": 1234567890
}
```

### Verify Token in Another Application

#### Go Example

```go
import (
    "github.com/golang-jwt/jwt/v5"
)

func verifyToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte("your-super-secret-jwt-key-change-this-in-production"), nil
    })
    
    if claims, ok := token.Claims.(*Claims); ok && token.Valid {
        return claims, nil
    }
    return nil, err
}
```

#### Node.js Example

```javascript
const jwt = require('jsonwebtoken');

function verifyToken(token) {
    try {
        const decoded = jwt.verify(token, 'your-super-secret-jwt-key-change-this-in-production');
        return decoded;
    } catch(err) {
        return null;
    }
}
```

#### Python Example

```python
import jwt

def verify_token(token):
    try:
        decoded = jwt.decode(token, 'your-super-secret-jwt-key-change-this-in-production', algorithms=['HS256'])
        return decoded
    except:
        return None
```

## Environment Variables Summary

### Backend (Railway)

```env
DB_HOST=${MYSQLHOST}
DB_PORT=3306
DB_USER=root
DB_PASSWORD=${MYSQLPASSWORD}
DB_NAME=railway
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
PORT=8080
```

### Frontend (Vercel)

```env
REACT_APP_API_URL=https://your-backend-url.railway.app
```

## Troubleshooting

### Backend not connecting to database

1. Check that environment variables are set correctly in Railway
2. Verify MySQL service is running
3. Check logs in Railway dashboard

### Frontend can't reach backend

1. Verify CORS is enabled in backend (already configured)
2. Check `REACT_APP_API_URL` is set correctly in Vercel
3. Ensure backend is deployed and running

### JWT token invalid

1. Verify `JWT_SECRET` is the same in all environments
2. Check token hasn't expired (24 hour default)
3. Ensure token is sent in `Authorization: Bearer <token>` header

## Security Notes

1. **Change JWT_SECRET**: Use a strong, random secret in production
2. **Use HTTPS**: Both Railway and Vercel provide HTTPS by default
3. **Password Hashing**: Passwords are hashed with bcrypt (cost 10)
4. **Token Expiration**: Tokens expire after 24 hours
5. **CORS**: Configure allowed origins in production

## Next Steps

1. Deploy backend to Railway
2. Initialize database with admin user
3. Deploy frontend to Vercel
4. Test the complete authentication flow
5. Use JWT tokens in your other applications

## Support

- Railway Docs: https://docs.railway.app
- Vercel Docs: https://vercel.com/docs
- GitHub Repository: https://github.com/MerpGoaterman/jwt-auth-system
