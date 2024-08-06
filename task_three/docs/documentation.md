# Library Management System

## Overview

A simple Library Management System implemented in Go. This application allows users to manage books and members, including functionalities for adding, removing, borrowing, and returning books.

## Features

- Add new books and members.
- Remove books from the library.
- Borrow and return books.
- List available and borrowed books.

## Setup

1. **Clone the Repository**:
   ```sh
   git clone git@github.com:solo21-12/A2SV_back_end_track.git
   ```

2. **Navigate to the Project Directory**:
   ```sh
   cd task_three
   ```

3. **Install Dependencies**:
   ```sh
   go mod tidy
   ```

4. **Run the Application**:
   ```sh
   go run main.go
   ```

## Usage

- **Add Book**: Follow prompts to enter book title and author.
- **Add Member**: Enter the member's name when prompted.
- **Remove Book**: Enter the Book ID to remove a book.
- **Borrow Book**: Enter Book ID and Member ID to borrow a book.
- **Return Book**: Enter Book ID and Member ID to return a book.
- **List Available Books**: Displays all available books.
- **List Borrowed Books**: Enter Member ID to view borrowed books.

## Code Structure

- **controllers**: Manages user interactions.
- **models**: Defines data structures.
- **services**: Contains business logic.
