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

## Starting the Application

To run the server, navigate to the project folder and use the following commands:

- `make run`: Starts the server.
- `make stop`: Stops the server.
- `make logs`: Displays the server logs.

## Database Configuration

This project uses MongoDB as the database. Instead of manually configuring MongoDB, the database is set up using a Docker container. 

### Requirements

- **Docker**: Ensure that Docker is installed on your system.

### Docker Commands

Use the following commands to manage the MongoDB container:

- `make run`: Starts the MongoDB container along with the application.
- `make stop`: Stops the MongoDB container along with the application.
- `make logs`: Displays logs from the running container.

Make sure Docker is running on your system, and then use these commands to easily set up and manage the database.

---