# 🔐 Go Fiber Auth Backend (JWT + RBAC)

This project is a backend authentication system built using **Go + Fiber**.

It implements:
- JWT-based Authentication
- Role-Based Access Control (RBAC)
- Middleware architecture
- Protected routes

---

## 🚀 Tech Stack

- Go (Golang)
- Fiber Framework
- JWT (Authentication)

---

## 📌 Features

- Login API with JWT token generation
- Middleware for authentication
- Role-based route protection
- Clean project structure

---

## 👥 Test Users

| Username   | Role        |
|-----------|------------|
| vikasJi   | superadmin |
| diptanshu | admin      |
| ashwin    | teacher    |
| student1  | student    |

---

## 🔑 API Endpoints

### 1. Login

POST `/login`

Body:
```json
{
  "username": "diptanshu",
  "password": "admin123"
}