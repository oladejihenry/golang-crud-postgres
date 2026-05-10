# golang-crud-postgres

A small REST API for posts built with **Go**, **Gin**, **GORM**, and **PostgreSQL**. Postgres runs in **Docker**; the app reads database settings from environment variables (optionally via a `.env` file using [godotenv](https://github.com/joho/godotenv)).

## Stack

- [Go](https://go.dev/) 1.26+
- [Gin](https://github.com/gin-gonic/gin) web framework
- [GORM](https://gorm.io/) with the [PostgreSQL driver](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)
- PostgreSQL 16 (Docker Compose)

## Prerequisites

- Go 1.26 or newer
- Docker and Docker Compose

## Quick start

### 1. Start PostgreSQL

From the project root:

```bash
docker compose up -d
```

This starts Postgres on port `5432` with user, password, and database name `postgres` (see `docker-compose.yml`).

### 2. Configure the app

Create a `.env` file in the project root (this file is gitignored). Example values that match the default Compose setup when the API runs on your host machine:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=postgres
DB_PORT=5432
DB_TIMEZONE=Europe/London
```

If `.env` is missing, the app still loads variables from your shell environment.

### 3. Run the API

```bash
go run .
```

The server listens on **http://localhost:8080**.

On first connect, GORM runs `AutoMigrate` for the `Post` model.

## HTTP API

| Method | Path | Description |
|--------|------|-------------|
| GET | `/` | Simple hello message |
| GET | `/health` | Health check |
| GET | `/posts` | List all posts |
| GET | `/posts/:id` | Get one post by ID |
| POST | `/posts` | Create a post (JSON body) |
| PUT | `/posts/:id` | Update title and content |
| DELETE | `/posts/:id` | Delete a post |

### Examples

Create a post:

```bash
curl -s -X POST http://localhost:8080/posts \
  -H "Content-Type: application/json" \
  -d '{"title":"Hello","content":"First post"}'
```

List posts:

```bash
curl -s http://localhost:8080/posts
```

## Project layout

```
.
├── main.go                 # Gin routes and server entrypoint
├── database/
│   └── database.go         # DSN from env, connect, AutoMigrate
├── handlers/
│   └── post_handler.go     # Post CRUD handlers
├── models/
│   └── post.go             # Post GORM model
├── docker-compose.yml      # PostgreSQL service
├── go.mod
└── .env                    # Local secrets (not committed; see .gitignore)
```

## License

Use and modify freely for learning or your own projects.
