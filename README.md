---

# Reusable Go + React JWT Authentication System

This project provides a complete, reusable JWT authentication system with a Go backend and a React frontend. It's designed for easy deployment to Railway and Vercel, with a focus on clear documentation and straightforward setup.

This system allows you to:

- Authenticate users with JWT tokens.
- Manage users, tenants, and roles.
- Secure your applications with a robust, standalone authentication service.

## Project Structure

```
.
├── backend/         # Go backend application
├── frontend/        # React frontend application
├── DEPLOYMENT.md    # Detailed deployment guide for Railway and Vercel
├── QUICKSTART.md    # 5-minute guide to get up and running
├── TESTING.md       # Comprehensive testing guide
└── README.md        # This file
```

## Key Features

### Backend (Go)

- **JWT Authentication**: Secure token-based authentication.
- **User Management**: CRUD operations for users.
- **Tenant and Role Support**: Multi-tenancy and role-based access control (RBAC).
- **MySQL Database**: Uses Railway's MySQL service.
- **Dockerized**: Ready for containerized deployment.
- **CORS**: Pre-configured for frontend communication.

### Frontend (React)

- **Login and Dashboard**: User-friendly interface for authentication and management.
- **Protected Routes**: Secure pages that require authentication.
- **Token Management**: Automatic handling of JWT tokens.
- **TypeScript**: Type-safe and scalable codebase.
- **Vercel Ready**: Optimized for deployment on Vercel.

## Getting Started

For a quick setup, follow the [**QUICKSTART.md**](QUICKSTART.md) guide. You'll have the entire system deployed in about 5 minutes.

For more detailed instructions, including local development and troubleshooting, refer to the [**DEPLOYMENT.md**](DEPLOYMENT.md) guide.

## Testing

A comprehensive guide for testing the application, both locally and in a deployed environment, is available in [**TESTING.md**](TESTING.md).

## Deployment

Due to the interactive nature of cloud provider authentication, I was unable to complete the final deployment steps. However, I have provided detailed, step-by-step instructions in the `DEPLOYMENT.md` and `QUICKSTART.md` files to make this process as smooth as possible for you.

## GitHub Repository

The complete source code is available on GitHub:

[https://github.com/MerpGoaterman/jwt-auth-system](https://github.com/MerpGoaterman/jwt-auth-system)

---
