# plaid-go

`plaid-go` is a Go service that connects to PostgreSQL and provides a basic HTTP endpoint for health-style verification.
The project is structured to support Plaid-related data models and repositories, with database migrations included.

## What this project includes

- A minimal HTTP server (`GET /`) returning a JSON success response.
- PostgreSQL connectivity using GORM.
- Environment-based configuration loading via `.env`.
- SQL migrations for `users` and `plaid_items` tables.
- Docker Compose setup for PostgreSQL and pgAdmin.

## Tech stack

- Go
- PostgreSQL
- GORM
- Docker Compose

## Prerequisites

Before running the project, make sure you have:

- Go installed
- Docker and Docker Compose installed
- A PostgreSQL database available (or use the provided Docker setup)

## Environment variables

Create a `.env` file in the project root with the following values:

```env
PLAID_CLIENT_ID=your_plaid_client_id
PLAID_SECRET=your_plaid_secret
PLAID_ENV=sandbox
DATABASE_URL=host=localhost user=myuser password=mypassword dbname=mydatabase port=5432 sslmode=disable
PORT=8080
```

`PLAID_ENV` supports: `sandbox`, `development`, or `production`.
If `PLAID_ENV` is missing or invalid, the app defaults to `sandbox`.

## Running with Docker (database only)

Start PostgreSQL and pgAdmin:

```bash
docker compose up -d
```

Services exposed by default:

- PostgreSQL: `localhost:5432`
- pgAdmin: `http://localhost:5050`

pgAdmin default credentials:

- Email: `admin@admin.com`
- Password: `root`

## Database migrations

Migration files are located in `migrations/` and use goose-style directives.
Apply them with your migration workflow/tool of choice before starting the app.

## Running the application

From the project root:

```bash
go run main.go
```

If startup is successful, you should see logs confirming database connectivity and server startup.

## API endpoint

### `GET /`

Response example:

```json
{
  "status": "success",
  "body": "database connected"
}
```

## Project structure

```text
config/        # environment config loading
db/            # database manager/connection
migrations/    # SQL migrations
models/        # data models
repositories/  # data access layer
services/      # business logic layer (user and plaid item services)
main.go        # entry point and HTTP server bootstrap
```

## Notes

- The current server is intentionally minimal and suitable as a starting point.
- Extend handlers, service logic, and repository usage as your Plaid integration grows.