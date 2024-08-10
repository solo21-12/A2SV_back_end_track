# Task Management API Documentation

## Project Overview

The Task Management API is a simple, in-memory application designed to manage tasks. It allows users to perform operations such as creating, updating, retrieving, and deleting tasks. The API is intended for demonstration purposes and uses in-memory storage to manage data, meaning that all data will be lost when the application is stopped.

## How to Test Locally

To test the Task Management API locally, follow these steps:

1. **Clone the Repository**

   Clone the repository to your local machine using the following command:

   ```sh
   git clone https://github.com/solo21-12/A2SV_back_end_track
   ```

2. **Navigate to the Task Management Folder**

   Change to the `task_four` directory where the API code is located:

   ```sh
   cd task_four
   ```

3. **Install Dependencies**

   Run `go mod tidy` to install the required Go modules and clean up any dependencies:

   ```sh
   go mod tidy
   ```

4. **Run the Application**

   Start the application with the following command:

   ```sh
   go run main.go
   ```

5. **Access the API**

   Once the application is running, you can access the API at `http://localhost:8081`. You can use tools like Postman or curl to test the available endpoints.

## API Documentation

You can find the comprehensive documentation for the Task Management API at the following link:

[Task Management API Documentation](https://documenter.getpostman.com/view/22911710/2sA3s3HAzk)

This documentation includes details on all available endpoints, request payloads, response formats, and error handling.
