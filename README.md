# Go CRUD API with Gin and Supabase

This project is a RESTful CRUD API built with Go, Gin framework, and Supabase PostgreSQL database. It provides endpoints to manage blog posts with full create, read, update, and delete operations.

## Features

- Full CRUD Operations: Create, Read, Update, and Delete blog posts
- RESTful Design: Clean and intuitive API endpoints
- Database Integration: Uses Supabase PostgreSQL with GORM ORM
- Environment Configuration: Easy setup with .env file
- Automatic Schema Migration: Database tables automatically created on first run

## Technologies Used

- **Go** - Programming language
- **Gin** - HTTP web framework
- **GORM** - ORM library for Go
- **Supabase** - PostgreSQL database service
- **godotenv** - Environment variable loader

## API Endpoints

| Method | Endpoint          | Description                |
|--------|-------------------|----------------------------|
| POST   | /createPost       | Create a new blog post     |
| GET    | /getPosts         | Get all blog posts         |
| GET    | /getPost/:id      | Get a single post by ID    |
| PUT    | /updatePost/:id   | Update a post by ID        |
| DELETE | /deletePost/:id   | Delete a post by ID        |

## Setup Instructions

### Prerequisites

- Go 1.19+ installed
- Supabase account (free tier works)

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/go-crud-api.git
    cd go-crud-api
    ```

2. Install dependencies:

    ```bash
    go mod download
    ```

3. Create a `.env` file in the root directory:

    ```env
    SUPABASE_DB_URL=your-supabase-connection-string
    PORT=8000
    ```

4. Get your Supabase connection string:

    - Create a new project in Supabase
    - Go to Database → Connection String
    - Copy the connection string and replace `your-supabase-connection-string`

### Running the Application

Start the server:

```bash
go run main.go
```

The server will run on [http://localhost:8000](http://localhost:8000)

## API Usage Examples

### Create a Post

```bash
curl -X POST http://localhost:8000/createPost \
  -H "Content-Type: application/json" \
  -d '{
    "title": "My First Post",
    "body": "This is the content of my first post"
  }'
```

Response:

```json
{
  "post": {
    "ID": 1,
    "CreatedAt": "2025-08-01T12:34:56.789Z",
    "UpdatedAt": "2025-08-01T12:34:56.789Z",
    "DeletedAt": null,
    "title": "My First Post",
    "body": "This is the content of my first post"
  }
}
```

### Get All Posts

```bash
curl http://localhost:8000/getPosts
```

Response:

```json
{
  "posts": [
    {
      "ID": 1,
      "CreatedAt": "2025-08-01T12:34:56.789Z",
      "UpdatedAt": "2025-08-01T12:34:56.789Z",
      "DeletedAt": null,
      "title": "My First Post",
      "body": "This is the content of my first post"
    },
    {
      "ID": 2,
      "CreatedAt": "2025-08-01T13:45:22.123Z",
      "UpdatedAt": "2025-08-01T13:45:22.123Z",
      "DeletedAt": null,
      "title": "Another Post",
      "body": "More content here"
    }
  ]
}
```

### Get Single Post

```bash
curl http://localhost:8000/getPost/1
```

Response:

```json
{
  "post": {
    "ID": 1,
    "CreatedAt": "2025-08-01T12:34:56.789Z",
    "UpdatedAt": "2025-08-01T12:34:56.789Z",
    "DeletedAt": null,
    "title": "My First Post",
    "body": "This is the content of my first post"
  }
}
```

### Update a Post

```bash
curl -X PUT http://localhost:8000/updatePost/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Updated Title",
    "body": "Updated content"
  }'
```

Response:

```json
{
  "post": {
    "ID": 1,
    "CreatedAt": "2025-08-01T12:34:56.789Z",
    "UpdatedAt": "2025-08-01T14:20:30.456Z",
    "DeletedAt": null,
    "title": "Updated Title",
    "body": "Updated content"
  }
}
```

### Delete a Post

```bash
curl -X DELETE http://localhost:8000/deletePost/1
```

Response:

```json
{
  "message": "Post deleted successfully"
}
```

## Project Structure

```
go-crud-api/
├── controllers/       # API controllers
│   └── postControllers.go
├── internals/         # Internal utilities
│   └── database.go    # Database connection
├── models/            # Data models
│   └── post.go
├── main.go            # Application entry point
├── go.mod             # Go module file
├── go.sum             # Go checksums
└── .env.example       # Environment variables example
```

## Database Schema

The application uses the following database schema:

```sql
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
    title VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
);
```