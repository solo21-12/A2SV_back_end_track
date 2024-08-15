
# Task Management API Documentation

You can find the comprehensive documentation for the Task Management API at the following link:

[Task Management API Documentation](https://documenter.getpostman.com/view/22911710/2sA3s3HAze)

This documentation includes details on all available endpoints, request payloads, response formats, and error handling.

## Configuration Instructions

### 1. **Running the Project**

To start the project, including MongoDB, ensure Docker is installed on your system and use the following commands:

1. **Run the Project**

   Start the project and MongoDB container using:

   ```bash
   make run
   ```

2. **Stop the Project**

   Stop the project and MongoDB container with:

   ```bash
   make stop
   ```

3. **View Logs**

   Check the logs for the project and MongoDB by running:

   ```bash
   make logs
   ```

## Authentication

### 1. **User Registration**

To register a new user, send a POST request to `/register` with the following payload:

**Endpoint:** /register  
**Method:** POST  
**Headers:**  

- Content-Type: application/json

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

**Endpoint:** /login  
**Method:** POST  
**Headers:**  

- Content-Type: application/json

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

After logging in, use the access token provided to authenticate subsequent requests. Include the token in the Authorization header as a Bearer token.

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
