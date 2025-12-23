# Railway Backend Deployment Guide

## Prerequisites

- Railway account with MySQL database already created
- GitHub repository connected to Railway
- MySQL connection details from Railway

## Deployment Steps

### Step 1: Add Backend Service to Railway Project

1. Go to your Railway project: https://railway.app/project/440af528-f60c-40f3-a238-cd9a33a2f922
2. Click the **"Create"** button
3. Select **"GitHub Repo"**
4. Choose **"MerpGoaterman/jwt-auth-system"**
5. Railway will detect the backend automatically

### Step 2: Configure Root Directory

1. In the service settings, set **Root Directory** to `backend`
2. Railway will use the `nixpacks.toml` and `railway.toml` files for build configuration

### Step 3: Set Environment Variables

Add the following environment variables in Railway service settings:

```
DB_HOST=<from Railway MySQL Variables>
DB_PORT=<from Railway MySQL Variables>
DB_USER=<from Railway MySQL Variables>
DB_PASSWORD=<from Railway MySQL Variables>
DB_NAME=<from Railway MySQL Variables>
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
PORT=8080
```

**Get MySQL credentials from Railway:**
- Click on the MySQL service
- Go to **Variables** tab
- Copy the values for MYSQLHOST, MYSQLPORT, MYSQLUSER, MYSQLPASSWORD, MYSQLDATABASE

### Step 4: Initialize Database

Connect to your Railway MySQL database and run the initialization script:

```sql
-- From backend/init.sql
CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(36) PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    tenant VARCHAR(100) NOT NULL,
    role VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_email (email),
    INDEX idx_tenant (tenant)
);
```

**Using Railway CLI:**
```bash
# Connect to MySQL
railway connect mysql

# Run the init script
source backend/init.sql
```

### Step 5: Deploy

1. Railway will automatically deploy when you push to master
2. Monitor the deployment in the **Deployments** tab
3. Check logs for any errors

### Step 6: Get Backend URL

1. Go to the backend service settings
2. Click on **"Generate Domain"** to get a public URL
3. Your backend will be available at: `https://your-service.up.railway.app`

### Step 7: Update Frontend Environment Variable

Update the frontend to point to your Railway backend URL:

1. Go to Vercel project settings
2. Add environment variable: `REACT_APP_API_URL=https://your-backend.up.railway.app`
3. Redeploy frontend

## Troubleshooting

### Build Fails

- Check that Go version is correct in `nixpacks.toml`
- Verify all dependencies are in `go.mod`
- Check build logs in Railway

### Database Connection Fails

- Verify environment variables are set correctly
- Check MySQL service is running
- Ensure database is initialized with tables

### CORS Errors

- Backend is configured to allow all origins in development
- For production, update CORS settings in `main.go`

## Alternative: Manual Railway CLI Deployment

If the web interface doesn't work, use Railway CLI:

```bash
# Login to Railway
railway login

# Link to your project
railway link 440af528-f60c-40f3-a238-cd9a33a2f922

# Set root directory
railway service --root backend

# Deploy
railway up
```

## Testing the Deployment

Once deployed, test the backend:

```bash
# Health check
curl https://your-backend.up.railway.app/health

# Register a user
curl -X POST https://your-backend.up.railway.app/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "SecurePass123!",
    "tenant": "default",
    "role": "admin"
  }'

# Login
curl -X POST https://your-backend.up.railway.app/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "SecurePass123!"
  }'
```

## Railway MySQL Credentials

Your Railway MySQL database details:
- Project ID: 440af528-f60c-40f3-a238-cd9a33a2f922
- Service: MySQL
- Connection details available in Railway dashboard under Variables tab
