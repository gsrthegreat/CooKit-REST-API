# CookIt API - JWT Authentication

A Go-based REST API with JWT authentication and bcrypt password hashing.

## Features

- ✅ JWT Authentication from scratch
- ✅ Bcrypt password hashing
- ✅ User registration and login
- ✅ MariaDB database
- ✅ Docker containerized
- ✅ RESTful endpoints

## Prerequisites

- Docker & Docker Compose
- Go 1.21+ (for local development)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/YOUR_USERNAME/CookIt.git
cd CookIt
```

2. Copy the example environment file:
```bash
cp .env.example .env
```

3. Edit `.env` and set your own secure values:
```bash
nano .env  # or use your preferred editor
```

4. Start the services:
```bash
docker-compose up -d --build
```

5. Check if services are running:
```bash
docker-compose ps
```

## API Endpoints

### Register User
```bash
POST /api/v1/register
Content-Type: application/x-www-form-urlencoded

username=john&password=password123
```

### Login
```bash
POST /api/v1/login
Content-Type: application/x-www-form-urlencoded

username=john&password=password123

Response: {"token":"eyJhbGc..."}
```

### Homepage (Protected)
```bash
GET /api/v1/
Authorization: Bearer YOUR_JWT_TOKEN

Response: Welcome john
```

### Logout
```bash
POST /api/v1/logout

Response: {"message":"Logged out successfully. Please remove the token from client."}
```

## Testing

Run the test script:
```bash
chmod +x test_api.sh
./test_api.sh
```

## Manual Testing

1. Register a user:
```bash
curl -X POST http://localhost:8080/api/v1/register \
  -d "username=john&password=password123"
```

2. Login:
```bash
TOKEN=$(curl -s -X POST http://localhost:8080/api/v1/login \
  -d "username=john&password=password123" | grep -o '"token":"[^"]*' | grep -o '[^"]*$')
```

3. Access protected endpoint:
```bash
curl http://localhost:8080/api/v1/ \
  -H "Authorization: Bearer $TOKEN"
```

## Project Structure

```
CookIt/
├── cmd/
│   └── api/
│       ├── main.go          # Main application
│       └── auth.go          # JWT logic
├── handlers/
│   └── authHandler.go       # Authentication handlers
├── internal/
│   └── database/
│       └── db.go            # Database connection
├── init.sql                 # Database schema
├── docker-compose.yml       # Docker services
├── Dockerfile               # API container
├── .env.example             # Example environment variables
├── .gitignore               # Git ignore rules
└── README.md                # This file
```

## Dependencies

```bash
go get github.com/golang-jwt/jwt/v5
go get github.com/go-sql-driver/mysql
go get golang.org/x/crypto/bcrypt
```

## Security Notes

- Passwords are hashed with bcrypt (cost: 14)
- JWT tokens expire after 24 hours
- Never commit `.env` file to version control
- Change default secrets in production
- Use HTTPS in production

## License

MIT