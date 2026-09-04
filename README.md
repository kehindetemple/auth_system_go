# 🔐 Go Authentication System

A backend authentication system built with **Go**, **Fiber**, **MongoDB**, **bcrypt**, and **JWT**.

## 🚀 Features

* User registration
* Password hashing with bcrypt
* User login
* Password verification
* JWT authentication
* Protected routes with authentication middleware
* MongoDB database storage
* Environment variables for secrets

## 🛠️ Tech Stack

* **Go**
* **Fiber v2**
* **MongoDB**
* **bcrypt**
* **JWT**
* **godotenv**

## 📁 Project Structure

```text
auth-system/
├── main.go
├── go.mod
├── go.sum
├── .env
├── .gitignore
│
├── config/
│   └── database.go
│
├── models/
│   └── user.go
│
├── handlers/
│   └── auth.go
│
├── middleware/
│   └── auth.go
│
├── routes/
│   └── routes.go
│
└── utils/
    ├── password.go
    └── token.go
```

## ⚙️ Setup

Clone the project and enter the directory:

```bash
git clone <your-repository-url>
cd auth-system
```

Install dependencies:

```bash
go mod tidy
```

Create a `.env` file:

```env
MONGO_URI=your-mongodb-connection-string
JWT_SECRET=your-long-random-secret
```

**Never commit your `.env` file to GitHub.**

## ▶️ Run the Server

```bash
go run .
```

The API runs on:

```text
http://localhost:8080
```

## 🔑 API Endpoints

### Register

```http
POST /register
```

Request:

```json
{
  "email": "test@example.com",
  "password": "12345678"
}
```

### Login

```http
POST /login
```

Request:

```json
{
  "email": "test@example.com",
  "password": "12345678"
}
```

Successful login returns a JWT:

```json
{
  "msg": "successful",
  "token": "your-jwt-token"
}
```

### Protected Routes

Protected routes require:

```http
Authorization: Bearer <your-jwt-token>
```

Requests without a valid token are rejected with:

```http
401 Unauthorized
```

## 🔐 Authentication Flow

```text
Register
   ↓
Hash Password
   ↓
Save User → MongoDB
   ↓
Login
   ↓
Find User by Email
   ↓
Verify Password
   ↓
Generate JWT
   ↓
Client Receives Token
   ↓
Protected Route
   ↓
JWT Middleware
   ↓
Allow / Reject Request
```

## 🧪 Testing

You can test the API using **Thunder Client**, Postman, or another API client.

Recommended test order:

1. Register a user
2. Login with the same credentials
3. Copy the JWT
4. Send the JWT to a protected route
5. Test the protected route without a token
6. Test with an invalid/expired token

## 📌 Status

**Core authentication system:** In progress

* [x] Registration
* [x] Password hashing
* [x] Login
* [x] Password verification
* [x] JWT generation
* [ ] Protected routes
* [ ] JWT middleware testing

## 📄 License

This project is for learning and development purposes.
