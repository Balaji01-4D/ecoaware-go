# 🛠️ EcoAware Complaint Tracker (Go Edition)

A robust, production-grade complaint management system built with **[Go](w)**, **[Gin Web Framework](w)**, and **[GORM](w)**. Designed for campus or organizational use to track and manage environment or facility-related issues securely with **cookie-based JWT authentication** and **role-based access control**.

---

## 📌 Features

* 🔐 Cookie-based **JWT Authentication**
* 🧑‍🎓 **User Registration & Login**
* 📝 Submit, View, and Manage Complaints
* 🗃️ Complaint Categorization
* 👥 **Role-based Access Control** (User/Admin)
* 🕒 Timestamps & Metadata Tracking
* 🖼️ Optional Image Upload Support
* ✅ Secure password hashing via **bcrypt**

---

## 🧱 Entity Relationship Diagram (ERD)

```text
+---------+       +------------+        +-------------+
|  User   |<----->|  Complaint |<------>|   Category  |
+---------+       +------------+        +-------------+
| id      |       | id         |        | id          |
| name    |       | title      |        | name        |
| email   |       | description|        +-------------+
| password|       | imagePath  |
| role    |       | status     |
+---------+       | timestamp  |
                  | created_by |
                  | category_id|
                  +-------------+
```

---

## ⚙️ Tech Stack

| Layer      | Technology                            |
| ---------- | ------------------------------------- |
| Language   | [Go](w) (Golang)                      |
| Web Server | [Gin](w)                              |
| ORM        | [GORM](w)                             |
| Database   | PostgreSQL  |
| Auth       | Cookie-based JWT                      |
| Security   | Bcrypt for hashing                    |

---

## 🌐 API Endpoints

### 🔓 Public (Authentication & Categories)

| Endpoint         | Method | Description         |
| ---------------- | ------ | ------------------- |
| `/auth/register` | POST   | Register a user     |
| `/auth/login`    | POST   | Login & set JWT     |
| `/auth/me`       | GET    | Get current user    |
| `/categories`    | GET    | List all categories |

---

### 👤 Authenticated Users

> Requires JWT (via cookie)

| Endpoint          | Method | Description            |
| ----------------- | ------ | ---------------------- |
| `/complaints`     | GET    | List user's complaints |
| `/complaints`     | POST   | Create a new complaint |
| `/complaints/:id` | GET    | Get complaint by ID    |
| `/user/password`  | PUT    | Update password        |
| `/user/profile`   | PUT    | Update user profile    |

---

### 🛠️ Admin Endpoints

> Requires `admin` role

| Endpoint                       | Method | Description             |
| ------------------------------ | ------ | ----------------------- |
| `/admin/users`                 | GET    | List all users          |
| `/admin/users/:id`             | DELETE | Delete a user           |
| `/admin/users/:id`             | PUT    | Update user data        |
| `/admin/complaints`            | GET    | List all complaints     |
| `/admin/complaints/:id/status` | PUT    | Update complaint status |

---

## 🔐 Security Overview

* ✅ JWT authentication using **HTTP-only secure cookies**
* ✅ Passwords stored securely using **bcrypt**
* ✅ Role-based access via middleware:

  * `/auth/**`: Public
  * `/complaints/**` & `/user/**`: Requires valid JWT
  * `/admin/**`: Requires `admin` role

---

## 🧪 Sample JSON Payloads

### 🔸 Register

```json
{
  "name": "Alice",
  "email": "alice@example.com",
  "password": "alice@123"
}
```

### 🔸 Login

```json
{
  "email": "alice@example.com",
  "password": "alice@123"
}
```

### 🔸 Create Complaint

```json
{
  "title": "Water leak in bathroom",
  "description": "Continuous water leakage in hostel block B",
  "imagePath": "uploads/image.jpg",
  "categoryId": 1
}
```

### 🔸 Admin - Update Complaint Status

```json
{
  "status": "IN_PROGRESS"
}
```

---

## 🚀 Running the Project

### 📋 Prerequisites

* Go 1.21+
* PostgreSQL

### 🛠 Setup Instructions

```bash
git clone https://github.com/yourusername/ecoaware-go.git
cd ecoaware-go

# Set env variables or .env file
DB_URL="host=localhost user=YOUR_USERNAME password=YOUR_PASSWORD port=5432 dbname=YOUR_DATABASE_NAME"
SECRET_KEY = "" # ANY RANDOM TEXT OR GET IT SECRET KEY GENERATOR


# Run the app
go run main.go
```

---

## 📌 Models Overview

### User

* `id`, `name`, `email`, `password`, `role`
* Role: `user` or `admin`

### Complaint

* `id`, `title`, `description`, `status`, `imagePath`, `createdBy`, `categoryId`

### Category

* `id`, `name`

### Enums

* Status: `PENDING`, `IN_PROGRESS`, `RESOLVED`, `REJECTED`

---

## 🔮 Future Enhancements

* 🧾 Swagger/OpenAPI Docs
* ⏳ Pagination & Filtering
* 📂 Image Storage via AWS S3
* 🧑‍💼 Admin Dashboard (React)
* 🛡️ Rate Limiting & Brute-force Protection

---

## 📄 License

This project is licensed under the MIT License.

---
