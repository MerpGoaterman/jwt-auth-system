# Quick Start Guide

Get your JWT Authentication System up and running in 5 minutes!

## What You Have

✅ **GitHub Repository**: https://github.com/MerpGoaterman/jwt-auth-system  
✅ **Railway MySQL Database**: Already created and running  
✅ **Go Backend**: Complete with JWT authentication, user management, and CORS  
✅ **React Frontend**: Login page, dashboard, and user management UI  

## Step 1: Deploy Backend to Railway (2 minutes)

1. Open Railway: https://railway.app/project/440af528-f60c-40f3-a238-cd9a33a2f922

2. Click **"Create"** → **"GitHub Repo"**

3. Select **`MerpGoaterman/jwt-auth-system`**

4. After Railway creates the service:
   - Click on the new service
   - Go to **Settings** → **Root Directory**
   - Set to: `backend`
   - Click **Save**

5. Go to **Variables** tab and add these environment variables:
   
   **Note**: Railway references like `${MYSQLHOST}` will automatically pull from your MySQL service.
   
   | Variable | Value |
   |----------|-------|
   | `DB_HOST` | `${MYSQLHOST}` |
   | `DB_PORT` | `3306` |
   | `DB_USER` | `root` |
   | `DB_PASSWORD` | `${MYSQLPASSWORD}` |
   | `DB_NAME` | `railway` |
   | `JWT_SECRET` | `your-super-secret-jwt-key-change-this-in-production` |
   | `PORT` | `8080` |
   
   Click **"Add Variable"** for each one, or use **"Raw Editor"** to paste all at once:
   ```
   DB_HOST=${MYSQLHOST}
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=${MYSQLPASSWORD}
   DB_NAME=railway
   JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
   PORT=8080
   ```

6. Go to **Settings** → **Networking** → Click **"Generate Domain"**

7. Copy your backend URL (e.g., `https://jwt-auth-system-production.up.railway.app`)

## Step 2: Initialize Database (1 minute)

Option A: Use Railway MySQL Client

1. In Railway, click on **MySQL** service
2. Click **"Connect"** → Copy the connection command
3. Run in your terminal and execute the SQL from `backend/init.sql`

Option B: Use API (after backend is deployed)

```bash
# Create admin user via API
curl -X POST https://YOUR-BACKEND-URL.railway.app/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Admin User",
    "email": "admin@example.com",
    "password": "admin123",
    "tenant_id": "default-tenant",
    "role": "admin"
  }'
```

## Step 3: Deploy Frontend to Vercel (2 minutes)

1. Go to https://vercel.com/new

2. Click **"Import Git Repository"**

3. Select **`MerpGoaterman/jwt-auth-system`**

4. Configure:
   - **Framework Preset**: Create React App
   - **Root Directory**: `frontend`
   - **Build Command**: `npm run build`
   - **Output Directory**: `build`

5. Add Environment Variable:
   - **Name**: `REACT_APP_API_URL`
   - **Value**: `https://YOUR-BACKEND-URL.railway.app` (from Step 1)

6. Click **"Deploy"**

7. Wait for deployment to complete (~2 minutes)

## Step 4: Test Your Application

1. Open your Vercel URL (e.g., `https://jwt-auth-system.vercel.app`)

2. Login with:
   - **Email**: `admin@example.com`
   - **Password**: `admin123`

3. You should see the dashboard with:
   - Your user information
   - JWT token (copy this!)
   - List of all users
   - "Add User" button (admin only)

## Step 5: Use JWT Token in Other Applications

### Test the Token

```bash
# Copy your token from the dashboard
export TOKEN="your-jwt-token-here"

# Test API calls
curl https://YOUR-BACKEND-URL.railway.app/api/users \
  -H "Authorization: Bearer $TOKEN"
```

### Token Contains

```json
{
  "user_id": "uuid",
  "email": "admin@example.com",
  "tenant_id": "default-tenant",
  "role": "admin",
  "exp": 1234567890
}
```

### Verify Token in Your App

**Go:**
```go
import "github.com/golang-jwt/jwt/v5"

token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return []byte("your-super-secret-jwt-key-change-this-in-production"), nil
})
```

**Node.js:**
```javascript
const jwt = require('jsonwebtoken');
const decoded = jwt.verify(token, 'your-super-secret-jwt-key-change-this-in-production');
```

**Python:**
```python
import jwt
decoded = jwt.decode(token, 'your-super-secret-jwt-key-change-this-in-production', algorithms=['HS256'])
```

## API Endpoints

### Authentication

**POST** `/api/auth/login`
```json
{
  "email": "admin@example.com",
  "password": "admin123"
}
```

Returns:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": "uuid",
    "name": "Admin User",
    "email": "admin@example.com",
    "tenant_id": "default-tenant",
    "role": "admin"
  }
}
```

### User Management (Requires JWT Token)

**GET** `/api/users` - Get all users  
**POST** `/api/users` - Create user (admin only)  
**GET** `/api/users/:id` - Get user by ID  
**PUT** `/api/users/:id` - Update user (admin only)  
**DELETE** `/api/users/:id` - Delete user (admin only)

All requests must include:
```
Authorization: Bearer YOUR_JWT_TOKEN
```

## What's Next?

1. **Change JWT Secret**: Use a strong random secret in production
2. **Add More Users**: Use the dashboard or API to create users
3. **Integrate with Your Apps**: Use the JWT tokens for authentication
4. **Customize**: Modify the code to fit your needs

## Troubleshooting

**Backend won't start?**
- Check environment variables are set correctly
- Verify MySQL service is running
- Check logs in Railway dashboard

**Frontend can't connect?**
- Verify `REACT_APP_API_URL` is set in Vercel
- Check backend is deployed and has a domain
- Try accessing backend URL directly

**Can't login?**
- Make sure database is initialized
- Check admin user was created
- Verify password is `admin123`

## Support

- Full Documentation: See `DEPLOYMENT.md`
- GitHub: https://github.com/MerpGoaterman/jwt-auth-system
- Railway Docs: https://docs.railway.app
- Vercel Docs: https://vercel.com/docs

---

**Total Time**: ~5 minutes  
**Cost**: Free tier on both Railway and Vercel  
**Result**: Production-ready JWT authentication system! 🎉
