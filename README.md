# Cache Service

A Go-based cache service built with Fiber web framework, PostgreSQL database, and Redis caching capabilities.

## Features

- 🚀 Fast HTTP API server using Fiber framework
- 🗄️ PostgreSQL database support with migrations
- 📦 Redis caching integration
- 📊 Monitoring endpoints
- ⚙️ Configuration management with Viper
- 📝 Structured logging with Zap
- 🔧 CLI commands with Cobra
- 🔄 Database migration support

## Prerequisites

- Go 1.23.3 or higher
- PostgreSQL database (or use Docker Compose)
- Docker and Docker Compose (optional, for local development)
- Redis (optional, if using Redis features)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Nattakornn/cache.git
cd cache
```

2. Install dependencies:
```bash
go mod download
```

3. Start PostgreSQL database using Docker Compose:
```bash
docker-compose up -d
```

This will start a PostgreSQL 12 container with the following default settings:
- Host: `localhost`
- Port: `5432`
- Username: `postgres`
- Password: `123456`
- Database: `cache_service`

**Note:** If you prefer to use your own PostgreSQL instance, skip this step and update the database configuration accordingly.

4. Configure the application by editing `config/config.yaml`:
```yaml
System:
  TimeZone: Asia/Bangkok

Log:
  Level: debug
  Color: true
  Json: false

Interface:
  Http:
    Host: "127.0.0.1"
    Port: 3000
    Name: "Cache Service"
    Version: "v1.0.0"
    ReadTimeout: 60
    WriteTimeout: 60
    BodyLimit: 10490000

Database:
  Host: "localhost"
  Port: 5432
  Protocol: "tcp"
  Username: "postgres"
  Password: "123456"
  Database: "cache_service"  # Match POSTGRES_DB in docker-compose.yaml
  Schema: "cache"
  SSLMode: "disable"
  MaxConnection: 50
```

**Note:** If using the Docker Compose setup, make sure the `Database.Database` value matches `POSTGRES_DB` in `docker-compose.yaml` (default: `cache_service`).

## Usage

### Start Database (Docker Compose)

Start the PostgreSQL database:
```bash
docker-compose up -d
```

Stop the database:
```bash
docker-compose down
```

View database logs:
```bash
docker-compose logs -f cache-service-postgres-db
```

### Run the Service

Start the cache service:
```bash
go run main.go serve
```

Or use a custom config file:
```bash
go run main.go serve --config /path/to/config.yaml
```

### Database Migrations

Run database migrations:
```bash
go run main.go migrate
```

Force migration (if needed):
```bash
go run main.go migrate --force
```

### Build

Build the binary:
```bash
go build -o cache-service main.go
```

Then run:
```bash
./cache-service serve
```

## Project Structure

```
.
├── cmd/                    # CLI commands
│   ├── cmd.go
│   ├── migrate.go         # Migration command
│   ├── root.go            # Root command
│   └── serve.go           # Serve command
├── config/                 # Configuration
│   ├── config.go          # Config loader
│   └── config.yaml        # Configuration file
├── modules/                # Application modules
│   ├── monitor/           # Monitoring module
│   └── servers/           # Server module
├── pkg/                    # Package libraries
│   ├── databases/         # Database connections
│   │   ├── postgressql/   # PostgreSQL implementation
│   │   └── redis/         # Redis implementation
│   └── logger/            # Logger implementation
├── docker-compose.yaml     # Docker Compose configuration for PostgreSQL
└── main.go                # Application entry point
```

## API Endpoints

The service exposes endpoints under `/api/v1`:

- Monitor endpoints (see `modules/monitor/` for details)

## Configuration

The application uses Viper for configuration management. Configuration can be provided via:

- YAML configuration file (default: `./config/config.yaml`)
- Command-line flag: `--config /path/to/config.yaml`

### Configuration Options

- **System**: Timezone settings
- **Log**: Logging level, color output, JSON format
- **Interface.Http**: Server host, port, timeouts, body limit
- **Database**: PostgreSQL connection settings

## Development

### Dependencies

Key dependencies:
- `github.com/gofiber/fiber/v2` - Web framework
- `github.com/jackc/pgx/v5` - PostgreSQL driver
- `github.com/jmoiron/sqlx` - SQL extensions
- `github.com/redis/go-redis/v9` - Redis client
- `github.com/spf13/cobra` - CLI framework
- `github.com/spf13/viper` - Configuration management
- `go.uber.org/zap` - Structured logging
- `gorm.io/gorm` - ORM

## License

[Add your license here]

## Author

Nattakorn

