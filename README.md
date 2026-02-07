# Go API Boilerplate

A production-ready Go API boilerplate with Echo v5, PostgreSQL, and sqlc.

## Stack

- **Router**: [Echo v5](https://echo.labstack.com/)
- **Database**: PostgreSQL with [sqlc](https://sqlc.dev/) for type-safe queries
- **Configuration**: [Viper](https://github.com/spf13/viper)
- **Logging**: slog (stdlib)
- **Validation**: [go-playground/validator](https://github.com/go-playground/validator)
- **Migrations**: [golang-migrate](https://github.com/golang-migrate/migrate)

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go           # Application entrypoint
├── internal/
│   ├── config/               # Configuration management
│   ├── database/             # Database connection
│   ├── handler/              # HTTP handlers
│   ├── middleware/           # Custom middleware
│   ├── repository/
│   │   └── sqlc/             # Generated sqlc code
│   └── validator/            # Request validation
├── db/
│   ├── migrations/           # SQL migrations
│   └── queries/              # sqlc queries
├── config.yaml.example       # Example configuration
├── docker-compose.yml        # Local development setup
├── Dockerfile                # Production Docker image
├── Makefile                  # Development commands
└── sqlc.yaml                 # sqlc configuration
```

## Getting Started

### Prerequisites

- Go 1.22+
- Docker & Docker Compose
- Make

### Install Development Tools

```bash
make install-tools
```

This installs:
- `sqlc` - SQL compiler
- `migrate` - Database migrations
- `golangci-lint` - Linter

### Setup

1. Clone the repository

2. Copy the example config:
   ```bash
   cp config.yaml.example config.yaml
   ```

3. Start PostgreSQL:
   ```bash
   make docker-up
   ```

4. Run migrations:
   ```bash
   make migrate-up
   ```

5. Generate sqlc code:
   ```bash
   make sqlc
   ```

6. Run the application:
   ```bash
   make run
   ```

The API will be available at `http://localhost:8080`.

## Available Commands

```bash
make help           # Show all commands
make build          # Build the application
make run            # Run the application
make test           # Run tests
make lint           # Run linter
make fmt            # Format code
make tidy           # Tidy go modules
make docker-up      # Start Docker services
make docker-down    # Stop Docker services
make migrate-up     # Run migrations
make migrate-down   # Rollback last migration
make migrate-create name=<name>  # Create new migration
make sqlc           # Generate sqlc code
```

## Configuration

Configuration can be provided via:

1. **Config file** (`config.yaml`)
2. **Environment variables** (uppercase with underscores, e.g., `DATABASE_HOST`)

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `SERVER_HOST` | Server host | `0.0.0.0` |
| `SERVER_PORT` | Server port | `8080` |
| `DATABASE_HOST` | PostgreSQL host | `localhost` |
| `DATABASE_PORT` | PostgreSQL port | `5432` |
| `DATABASE_USER` | PostgreSQL user | `postgres` |
| `DATABASE_PASSWORD` | PostgreSQL password | `postgres` |
| `DATABASE_NAME` | Database name | `go_api_boilerplate` |
| `DATABASE_SSLMODE` | SSL mode | `disable` |
| `LOG_LEVEL` | Log level (debug/info/warn/error) | `info` |
| `LOG_FORMAT` | Log format (json/text) | `json` |

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/health` | Liveness check |
| GET | `/health/ready` | Readiness check (includes DB) |

## Adding New Features

### 1. Create a Migration

```bash
make migrate-create name=create_products
```

Edit the generated files in `db/migrations/`.

### 2. Write SQL Queries

Add queries to `db/queries/` following [sqlc syntax](https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html).

### 3. Generate Code

```bash
make sqlc
```

### 4. Create Handler

Add handlers in `internal/handler/` and register routes in `internal/server/server.go`.

## License

MIT
