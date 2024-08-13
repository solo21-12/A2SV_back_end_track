# Task Management API Documentation

You can find the comprehensive documentation for the Task Management API at the following link:

[Task Management API Documentation](https://documenter.getpostman.com/view/22911710/2sA3s3HAze)

This documentation includes details on all available endpoints, request payloads, response formats, and error handling.

## Authentication

### 1. **User Registration**

To register a new user, send a POST request to `/register` with the following payload:

**Endpoint:** `/register`  
**Method:** POST  
**Headers:**  
- `Content-Type: application/json`

**Request Body:**

```json
{
    "email": "user@example.com",
    "password": "your_password"
}
```

**Response:**

- **200 OK**: Registration successful.
- **400 Bad Request**: Invalid input or user already exists.
- **500 Internal Server Error**: Server error.

### 2. **User Login**

To log in and receive an access token, send a POST request to `/login` with the following payload:

**Endpoint:** `/login`  
**Method:** POST  
**Headers:**  
- `Content-Type: application/json`

**Request Body:**

```json
{
    "email": "user@example.com",
    "password": "your_password"
}
```

**Response:**

- **200 OK**: Returns an access token in the response body.
  
  **Response Body:**

  ```json
  {
      "access_token": "your_jwt_token_here"
  }
  ```

- **401 Unauthorized**: Invalid email or password.
- **500 Internal Server Error**: Server error.

### 3. **Authenticated Requests**

After logging in, use the access token provided to authenticate subsequent requests. Include the token in the `Authorization` header as a Bearer token.

**Header Format:**

```
Authorization: Bearer <access_token>
```

**Example:**

```http
GET /api/protected-resource
Host: example.com
Authorization: Bearer your_jwt_token_here
```

**Response:**

- **200 OK**: Authorized request.
- **401 Unauthorized**: Invalid or missing token.
- **403 Forbidden**: Insufficient permissions.

## Configuration Instructions

### 1. **MongoDB Configuration**

To configure MongoDB for your Task Management API, follow these steps:

1. **Create a `.env` File**

   In the `bootstrap` directory of the project, create a file named `.env`.

2. **Add Configuration Variables**

   In the `.env` file, add the following lines with your configuration values:

   ```plaintext
   MONGO_URL=<your-mongodb-connection-string>
   MONGO_DATABASE=task_manager
   SERVER_ADDRESS=:8081
   USER_COLLECTION=users
   JWT_SECRET=<your-jwt-secret>
   ALLOWED_USERS=admin
   TASK_COLLECTION=tasks
   ```

   Replace `<your-mongodb-connection-string>` with the actual connection string for your MongoDB instance. This string typically looks like `mongodb://username:password@host:port/database`. Replace `<your-jwt-secret>` with your secret key for JWT authentication.

### 2. **Install Required Modules**

To install the required Go modules for the Task Management API, run the following command in your terminal:

```bash
go get .
```

This command will fetch and install all dependencies specified in your `go.mod` file.

## Example `.env` File

Here is an example of what your `.env` file might look like:

```plaintext
MONGO_URL=mongodb://username:password@localhost:27017/task_management
MONGO_DATABASE=task_manager
SERVER_ADDRESS=:8081
USER_COLLECTION=users
JWT_SECRET=your_jwt_secret_key
ALLOWED_USERS=admin
TASK_COLLECTION=tasks
```

## Starting the Application

Ensure that you have your `.env` file configured correctly and all required Go modules installed. You can then start your application using the following command:

```bash
go run main.go
```

## Folder Structure

The project follows the Clean Architecture pattern and has the following folder structure:

```
bootstrap
│   ├── app.go
│   ├── database.go
│   └── env.go
├── Delivery
│   ├── controllers
│   │   ├── login.controller.go
│   │   ├── promote_user.controller.go
│   │   ├── sign_up.controller.go
│   │   └── task.controller.go
│   ├── main.go
│   ├── routers
│   │   ├── login.router.go
│   │   ├── promote.route.go
│   │   ├── router.go
│   │   ├── sign_up.router.go
│   │   └── task.router.go
│   └── tmp
│       └── build-errors.log
├── doc
│   └── api_documentation.md
├── Domain
│   ├── error_response.go
│   ├── jwt_custome.go
│   ├── login.go
│   ├── promote.go
│   ├── sign_up.go
│   ├── task.go
│   ├── user.go
│   └── validate.go
├── go.mod
├── go.sum
├── Infrastructure
│   ├── auth.middleware.go
│   ├── get_objectID.go
│   ├── jwt.service.go
│   └── password.service.go
├── Repositories
│   ├── task_repository.go
│   └── user_repository.go
└── UseCases
    ├── login.usecase.go
    ├── promote.usecase.go
    ├── sign_up.usecase.go
    └── task.usecase.go
```