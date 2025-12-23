# JWT Authentication Frontend (React + TypeScript)

Modern React frontend for JWT authentication system with user management interface.

## Features

- Login page with authentication
- Protected routes
- User management dashboard
- Create, view, and delete users (admin only)
- JWT token display
- Responsive design
- Context-based state management

## Tech Stack

- React 18
- TypeScript
- React Router DOM
- Axios
- Context API

## Getting Started

### Installation

```bash
npm install
```

### Environment Variables

Create a `.env` file:

```
REACT_APP_API_URL=http://localhost:8080
```

### Running Locally

```bash
npm start
```

### Building for Production

```bash
npm run build
```

## Deployment to Vercel

1. Connect GitHub repository to Vercel
2. Set `REACT_APP_API_URL` environment variable
3. Deploy automatically on push

## License

MIT
