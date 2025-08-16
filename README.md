# go-bookman-app

Simple Book Management API written in Go

## How to Use

1. Clone this repository.

2. Copy the `.env` file, then insert the required configurations

```sh
cp .env.example .env
```

3. Make sure the Redis server and Postgres database are running.

4. Start the application.

```sh
go run main.go
```

5. Test the application by sending the HTTP request.

```sh
# create a new book
curl -XPOST -H "Content-type: application/json" -d '{"title":"algorithm","description":"learn algorithm","publisher":"goodbooks"}' 'http://localhost:1323/api/v1/books'
```

## Running with Docker Compose

The application can be started using Docker Compose.

1. Make sure to modify the database and Redis host based on the service name.

2. Start the application.

```sh
docker compose up -d
```

3. Stop the application.

```sh
docker compose down
```

## List of Available Endpoints

| **Endpoint**   | **Method** | **Description**       |
| -------------- | ---------- | --------------------- |
| `/books`       | `POST`     | Create a new book     |
| `/books/batch` | `POST`     | Create multiple books |
| `/books`       | `GET`      | Get all books         |
| `/books/:id`   | `GET`      | Get book by ID        |

## Notes

- This repository uses Postgres as a main database and Redis as a cache storage.
- This repository contains basic example of caching strategies.
