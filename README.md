# Real Estate API project

A robust, clean-architecture RESTful API for a Real Estate platform built with **Go (Golang)** and the **Gin Gonic** framework. Features full CRUD functionalities, dynamic filtering, offset pagination, multiple-image uploading, and role-based access control.

---

## 🚀 Features

- **Clean Architecture:** Strict separation of layers (Models, Controllers, Services, Middleware, Routes, and Utils).
- **Authentication & Security:** JWT-based user authentication with secure password hashing.
- **Role-Based Access Control (RBAC):** Permissions managed via a dedicated user role column (`admin`, `agent`, `user`).
- **Dynamic Real Estate Search:** Advanced filtering by `city` (partial-match using `LIKE`), price ranges (`min_price`, `max_price`), and bedroom counts.
- **Reusable Pagination Helper:** A decoupled custom utility handling unified meta-response formatting.
- **One-to-Many Image Management:** Support for uploading multiple property images simultaneously with unique file-naming timestamps.
- **Hot Reload Development:** pre-configured using `air` for a fast, node-like development workflow on Windows environments.

---

## 📂 Project Structure

```text
real_estate_api/
├── cmd/
│   └── server/
│       └── main.go       # Application Entry Point
├── config/               # Database connection and environment configurations
├── controllers/          # HTTP request handlers & input validation
├── middleware/           # JWT & Role validation filters
├── models/               # GORM schemas and database definitions
├── routes/               # Route grouping and endpoint registration
├── services/             # Core business logic layer
├── uploads/              # Generated directory holding local physical images
├── utils/                # Reusable global functions (Pagination helper)
├── .air.toml             # Live-reload configuration tailored for Windows paths
├── go.mod                # Go module dependencies
└── README.md             # Project documentation

```

## 🛠️ Tech Stack & Dependencies

- **Language**: Go (Golang)

- **Web Framework**: Gin Gonic

- **ORM**: GORM

- **Database**: PostgreSQL

- **Live Reload Tool**: Air

## ⚡ Getting Started

1. **Prerequisites**
Make sure you have Go installed on your machine.

2. **Live Reload Development**
```go install github.com/air-verse/air@latest```
And then,
```air```

1. **Alternative Standard Run**
```go run cmd/server/main.go```