# 🔐 Go Fiber Authentication System (JWT + RBAC)

Backend authentication system built using **Golang (Fiber)** with **JWT authentication** and **Role-Based Access Control (RBAC)**.

---

## 🚀 Tech Stack

* Golang
* Fiber
* JWT

---

## 📁 Project Structure

```
go-auth-backend/
│
├── server.go
├── go.mod
│
├── src/
│   ├── app.go
│   ├── routes.go
│   ├── controllers.go
│   ├── models.go
│   ├── middleware.go
│
└── README.md
```

---

## 👥 Demo Users

| Username  | Role       |
| --------- | ---------- |
| vikasJi   | superadmin |
| diptanshu | admin      |
| ashwin    | teacher    |
| student1  | student    |

---

## 🔑 API Endpoints

### Login

**POST** `/login`

**Body:**

```json
{
  "username": "diptanshu",
  "password": "admin123"
}
```

**Response:**

```json
{
  "token": "JWT_TOKEN"
}
```

---

## 🔒 Protected Routes

| Route      | Access            |
| ---------- | ----------------- |
| /dashboard | Any authenticated |
| /admin     | admin, superadmin |
| /teacher   | teacher           |
| /student   | student           |

---

## 🔐 Authorization Header

```
Authorization: Bearer <JWT_TOKEN>
```

---

## 🔄 Flow

1. Login → validate user
2. Generate JWT (username, role, exp)
3. Send token in header
4. AuthMiddleware → verify + extract claims
5. RoleMiddleware → check access
6. Return response

---

## ▶️ Run Locally

```
go run server.go
```

Server:

```
http://localhost:3000
```

---
