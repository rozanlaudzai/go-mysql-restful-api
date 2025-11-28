# Go-MySQL RESTful API

A RESTful API designed for category management, built with Go and MySQL. This project demonstrates clean architecture patterns, secure API key authentication, and CRUD operations.

## 🚀 Features

* **CRUD Operations:** Create, Read, Update, and Delete categories.
* **Security:** Middleware-based authentication using API Keys.

## 🛠️ Tech Stack

* **Language:** Go (Golang)
* **Database:** MySQL
* **Dependencies:** `go-sql-driver/mysql`, `julienschmidt/httprouter`, `joho/godotenv`, `go-playground/validator/v10`

## ⚙️ Configuration

The application requires the following environment variables to be set. You can export them in your terminal or use a `.env` file if supported.

| Variable      | Description                                  | Example Value      |
| :------------ | :------------------------------------------- | :----------------- |
| `DB_USERNAME` | MySQL database username                      | `root`             |
| `DB_PASSWORD` | MySQL database password                      | `password123`      |
| `DB_HOST`     | Database host address                        | `localhost`        |
| `DB_PORT`     | Database port                                | `3306`             |
| `DB_NAME`     | Name of the database schema                  | `go_restful_db`    |
| `SERVER_PORT` | The port the API server will listen on       | `3000`             |
| `API_KEY`     | The secret key required for request headers  | `secret-api-key`   |

## 💾 Database Setup

Before running the application, ensure your MySQL instance is running and the database is initialized.

1.  Create a database matching your `DB_NAME`.
2.  Run the initialization script to create the required table:
    * [initial_query.sql](initial_query.sql)

## 🏁 Getting Started

### 1. Install required go modules

Install the required Go modules.

```bash
go mod tidy
```

### 2. Run the Application

You can run the application directly or build a binary.

**Development Mode:**

```bash
go run main.go
```

**Production Build:**

```bash
go build -o go-mysql-restful-api
./go-mysql-restful-api
```

The server should now be running at `http://localhost:3000` (or your defined port).

## 🔐 Authentication

This API uses **API Key Authentication**. Every request to protected endpoints must include the `X-API-Key` header.

**Header Format:**

```http
X-API-Key: <YOUR_ENV_API_KEY>
```

> **Note:** If the key is missing or incorrect, the API will return `401 Unauthorized`.

## 📡 API Documentation

For full details on endpoints, parameters, and responses, please refer to the OpenAPI Specification:

👉 **[View OpenAPI Documentation](apispec.json)**

### Quick Usage Example (Create Category)

```bash
curl -X POST http://localhost:3000/api/categories \
  -H "Content-Type: application/json" \
  -H "X-API-Key: secret-api-key" \
  -d '{"name": "Electronics"}'
```
