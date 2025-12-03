# Go-MySQL RESTful API

A RESTful API designed for category management, built with Go and MySQL. This project demonstrates clean architecture patterns, secure API key authentication, comprehensive error handling, and full CRUD operations.

## 📑 Table of Contents

- [Features](#-features)
- [Tech Stack](#️-tech-stack)
- [Project Structure](#-project-structure)
- [Environment Variables](#️-environment-variables)
- [Database Setup](#-database-setup)
- [Getting Started](#-getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Authentication](#-authentication)
- [API Documentation](#-api-documentation)
  - [Base URL](#base-url)
  - [Endpoints](#endpoints)
  - [Error Responses](#error-responses)
  - [OpenAPI Specification](#openapi-specification)
- [Testing](#-testing)
  - [Test Setup](#test-setup)
  - [Running Tests](#running-tests)
  - [Test Coverage](#test-coverage)
- [Usage Examples](#-usage-examples)
  - [Using cURL](#using-curl)
  - [Using HTTP File](#using-http-file)
- [Architecture](#️-architecture)
  - [Request Flow](#request-flow)
  - [Error Handling](#error-handling)
- [License](#-license)

## 🚀 Features

* **Full CRUD Operations:** Create, Read, Update, and Delete categories with proper validation
* **Security:** Middleware-based API Key authentication for all endpoints
* **Clean Architecture:** Separation of concerns with Controller, Service, and Repository layers
* **Input Validation:** Request validation using `go-playground/validator`
* **Error Handling:** Comprehensive error handling with custom exceptions and panic recovery
* **Database Connection Pooling:** Optimized database connections with configurable pool settings
* **Unit Tests:** Unit tests covering all endpoints and edge cases
* **OpenAPI Specification:** Complete API documentation in OpenAPI 3.0 format

## 🛠️ Tech Stack

* **Language:** Go 1.25.1
* **Database:** MySQL
* **Router:** `julienschmidt/httprouter` - HTTP request router
* **Database Driver:** `go-sql-driver/mysql` - MySQL driver for Go
* **Environment Variables:** `joho/godotenv` - Environment variable management
* **Validation:** `go-playground/validator/v10` - Struct validation
* **Testing:** `stretchr/testify` - Testing toolkit

## 📁 Project Structure

This project follows clean architecture principles with clear separation of concerns:

```
go-mysql-restful-api/
├── app/                    # Application configuration
│   ├── database.go        # Database connection and pooling
│   └── router.go          # HTTP router setup
├── controller/            # HTTP request handlers
│   ├── category_controller.go
│   └── category_controller_impl.go
├── service/               # Business logic layer
│   ├── category_service.go
│   └── category_service_impl.go
├── repository/            # Data access layer
│   ├── category_repository.go
│   └── category_repository_impl.go
├── model/                 # Data models
│   ├── domain/           # Domain entities
│   │   └── category.go
│   └── web/              # Request/Response DTOs
│       ├── category_create_request.go
│       ├── category_update_request.go
│       ├── category_response.go
│       └── web_response.go
├── middleware/            # HTTP middleware
│   └── auth_middleware.go
├── exception/             # Error handling
│   ├── error_handler.go
│   ├── not_found_error.go
│   └── write_error_response.go
├── test/                  # Unit tests
│   └── category_controller_test.go
├── main.go               # Application entry point
├── initial_query.sql     # Database schema
├── apispec.json          # OpenAPI specification
├── test.http             # HTTP request examples
└── README.md
```

## ⚙️ Environment Variables

The application requires the following environment variables to be set. You can export them in your terminal or use a `.env` file (loaded automatically via `godotenv`).

| Variable      | Description                                  | Example Value      |
| :------------ | :------------------------------------------- | :----------------- |
| `DB_USERNAME` | MySQL database username                      | `root`             |
| `DB_PASSWORD` | MySQL database password                      | `password123`      |
| `DB_HOST`     | Database host address                        | `localhost`        |
| `DB_PORT`     | Database port                                | `3306`             |
| `DB_NAME`     | Name of the database schema                  | `go_restful_api`   |
| `SERVER_PORT` | The port the API server will listen on       | `3000`             |
| `API_KEY`     | The secret key required for request headers  | `secret-api-key`   |

### Example `.env` file:

```env
DB_USERNAME=root
DB_PASSWORD=password123
DB_HOST=localhost
DB_PORT=3306
DB_NAME=go_restful_api
SERVER_PORT=3000
API_KEY=secret-api-key
```

## 💾 Database Setup

Before running the application, ensure your MySQL instance is running and the database is initialized.

1. **Create the database:**
   ```sql
   CREATE DATABASE go_restful_api;
   ```

2. **Run the initialization script:**
   ```bash
   mysql -u root -p go_restful_api < initial_query.sql
   ```
   
   Or manually execute the SQL from [initial_query.sql](initial_query.sql):
   ```sql
   CREATE TABLE category (
       id INT PRIMARY KEY AUTO_INCREMENT,
       name VARCHAR(200) NOT NULL
   ) ENGINE = InnoDB;
   ```

### Database Connection Pooling

The application uses optimized database connection pooling:
- **Max Idle Connections:** 5
- **Max Open Connections:** 20
- **Connection Max Idle Time:** 10 minutes
- **Connection Max Lifetime:** 1 hour

## 🏁 Getting Started

### Prerequisites

- Go 1.25.1 or higher
- MySQL 5.7+ or MySQL 8.0+
- Git (optional)

### Installation

1. **Clone the repository** (if applicable):
   ```bash
   git clone <repository-url>
   cd go-mysql-restful-api
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Set up environment variables:**
   Create a `.env` file in the root directory with your configuration (see [Environment Variables](#-environment-variables) section).

4. **Initialize the database:**
   Follow the steps in [Database Setup](#-database-setup).

5. **Run the application:**

   **Development Mode:**
   ```bash
   go run main.go
   ```

   **Production Build:**
   ```bash
   go build -o go-mysql-restful-api
   ./go-mysql-restful-api
   ```

   The server will start and listen at `http://localhost:3000` (or your configured `SERVER_PORT`).

## 🔐 Authentication

This API uses **API Key Authentication** for all endpoints. Every request must include the `X-API-Key` header with a valid API key.

**Header Format:**
```http
X-API-Key: <YOUR_ENV_API_KEY>
```

**Example:**
```http
X-API-Key: secret-api-key
```

> **⚠️ Security Note:** If the API key is missing or incorrect, the API will return `401 Unauthorized`. Make sure to keep your API key secure and never commit it to version control.

## 📡 API Documentation

### Base URL

```
http://localhost:<your-port>/api
```

### Endpoints

#### 1. Get All Categories

Retrieve a list of all categories.

**Request:**
```http
GET /api/categories
X-API-Key: <your-api-key>
```

**Response:**
```json
{
  "code": 200,
  "status": "OK",
  "data": [
    {
      "id": 1,
      "name": "Electronics"
    },
    {
      "id": 2,
      "name": "Fashion"
    }
  ]
}
```

#### 2. Get Category by ID

Retrieve a specific category by its ID.

**Request:**
```http
GET /api/categories/{categoryId}
X-API-Key: <your-api-key>
```

**Response (Success):**
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "id": 1,
    "name": "Electronics"
  }
}
```

**Response (Not Found - 404):**
```json
{
  "code": 404,
  "status": "NOT FOUND",
  "data": "category not found"
}
```

#### 3. Create Category

Create a new category.

**Request:**
```http
POST /api/categories
X-API-Key: <your-api-key>
Content-Type: application/json

{
  "name": "Electronics"
}
```

**Response (Success):**
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "id": 1,
    "name": "Electronics"
  }
}
```

**Response (Bad Request - 400):**
```json
{
  "code": 400,
  "status": "BAD REQUEST",
  "data": "invalid fields"
}
```

#### 4. Update Category

Update an existing category by ID.

**Request:**
```http
PUT /api/categories/{categoryId}
X-API-Key: <your-api-key>
Content-Type: application/json

{
  "name": "Updated Category Name"
}
```

**Response (Success):**
```json
{
  "code": 200,
  "status": "OK",
  "data": {
    "id": 1,
    "name": "Updated Category Name"
  }
}
```

**Response (Not Found - 404):**
```json
{
  "code": 404,
  "status": "NOT FOUND",
  "data": "category not found"
}
```

#### 5. Delete Category

Delete a category by ID.

**Request:**
```http
DELETE /api/categories/{categoryId}
X-API-Key: <your-api-key>
```

**Response (Success):**
```json
{
  "code": 200,
  "status": "OK"
}
```

**Response (Not Found - 404):**
```json
{
  "code": 404,
  "status": "NOT FOUND",
  "data": "category not found"
}
```

### Error Responses

The API uses consistent error response format:

```json
{
  "code": <HTTP_STATUS_CODE>,
  "status": "<STATUS_TEXT>",
  "data": "<ERROR_MESSAGE>"
}
```

**Common HTTP Status Codes:**
- `200` - OK (Success)
- `400` - Bad Request (Validation errors)
- `401` - Unauthorized (Invalid or missing API key)
- `404` - Not Found (Resource not found)
- `500` - Internal Server Error (Server errors)

### OpenAPI Specification

For complete API documentation including request/response schemas, refer to the OpenAPI 3.0 specification:

👉 **[View OpenAPI Documentation](apispec.json)**

You can use tools like [Swagger UI](https://swagger.io/tools/swagger-ui/) or [Postman](https://www.postman.com/) to import and explore the API specification.

## 🧪 Testing

The project includes comprehensive unit tests covering all endpoints and edge cases.

### Test Setup

1. **Create a test environment file:**
   Create a `.env.test` file in the root directory with test database configuration:
   ```env
   DB_USERNAME=root
   DB_PASSWORD=password123
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=go_restful_api_test
   SERVER_PORT=3000
   API_KEY=test-api-key
   ```

2. **Create a separate test database:**
   ```sql
   CREATE DATABASE go_restful_api_test;
   ```
   
   Then run the schema:
   ```bash
   mysql -u root -p go_restful_api_test < initial_query.sql
   ```

   > **⚠️ Important:** Always use a separate database for testing to avoid data loss or corruption in your development database.

### Running Tests

Run all tests:
```bash
go test ./test/...
```

Run tests with verbose output:
```bash
go test -v ./test/...
```

Run a specific test:
```bash
go test -v ./test/... -run TestCreateCategorySuccess
```

### Test Coverage

The test suite includes:
- ✅ Create category (success and validation errors)
- ✅ Get all categories
- ✅ Get category by ID (success and not found)
- ✅ Update category (success, validation errors, and not found)
- ✅ Delete category (success and not found)
- ✅ Authentication (unauthorized access)

## 📝 Usage Examples

### Using cURL

**Create a category:**
```bash
curl -X POST http://localhost:3000/api/categories \
  -H "Content-Type: application/json" \
  -H "X-API-Key: secret-api-key" \
  -d '{"name": "Electronics"}'
```

**Get all categories:**
```bash
curl -X GET http://localhost:3000/api/categories \
  -H "X-API-Key: secret-api-key"
```

**Get category by ID:**
```bash
curl -X GET http://localhost:3000/api/categories/1 \
  -H "X-API-Key: secret-api-key"
```

**Update category:**
```bash
curl -X PUT http://localhost:3000/api/categories/1 \
  -H "Content-Type: application/json" \
  -H "X-API-Key: secret-api-key" \
  -d '{"name": "Updated Electronics"}'
```

**Delete category:**
```bash
curl -X DELETE http://localhost:3000/api/categories/1 \
  -H "X-API-Key: secret-api-key"
```

### Using HTTP File

The project includes a `test.http` file with example requests that can be used with REST Client extensions in VSCode or JetBrains IDEs.

## 🏗️ Architecture

This project follows **Clean Architecture** principles:

1. **Controller Layer:** Handles HTTP requests and responses
2. **Service Layer:** Contains business logic and validation
3. **Repository Layer:** Manages database operations
4. **Domain Layer:** Defines core business entities
5. **Middleware:** Handles cross-cutting concerns (authentication, error handling)

### Request Flow

```
HTTP Request
    ↓
Auth Middleware (API Key validation)
    ↓
Router
    ↓
Controller (Request parsing, response formatting)
    ↓
Service (Business logic, validation)
    ↓
Repository (Database operations)
    ↓
Database
```

### Error Handling

The application includes comprehensive error handling:
- **Panic Recovery:** Global panic handler for unexpected errors
- **Custom Exceptions:** `NotFoundError` for resource not found scenarios
- **Validation Errors:** Automatic handling of validation failures
- **Consistent Error Responses:** Standardized error response format

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.