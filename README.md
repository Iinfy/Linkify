# linkify

A simple URL shortening service with click analytics.

## Quick Start

Prerequisites: Docker and Docker Compose.

```bash
docker-compose -f docker-compose.dev.yml up --build
```

Services available at:
- Caddy (frontend + API proxy): http://localhost
- Backend API: http://localhost:8090
- Frontend: http://localhost:8080
- PostgreSQL: localhost:5432
- ClickHouse: localhost:9000

## Environment Setup

Copy `.env.example` to `.env` and configure:

```bash
cp .env.example .env
```

Required variables:
- `PORT` - Backend port (default: 8090)
- `POSTGRES_URL` - PostgreSQL connection string
- `CLICKHOUSE_URL` - ClickHouse connection string
- `CORS_HOST` - Allowed origin (default: http://localhost)

## API Endpoints

### POST /slink
Create a short link.

```json
{"url": "https://example.com"}
```

Response:
```json
{"url": "<12-char-hash>"}
```

### GET /s/:hash
Redirect to original URL. Records click in ClickHouse.

### GET /short/:hash
Get link stats.

Response:
```json
{"click_count": 42, "url": "https://example.com"}
```

## Architecture

- **Backend** (`app/`): Go 1.26.2 + Gin framework
- **Database**: PostgreSQL for URL storage
- **Analytics**: ClickHouse for click tracking
- **Frontend**: Static HTML/CSS/JS served via Caddy
- **Proxy**: Caddy reverse-proxy routes `/s/*`, `/slink`, `/short/*` to backend, rest to frontend

## Development

Run backend directly (requires DBs running):

```bash
cd app && go run main.go
```

Run tests:

```bash
cd app && go test ./...
```

## Stopping Services

```bash
docker-compose -f docker-compose.dev.yml down
```

To remove volumes:

```bash
docker-compose -f docker-compose.dev.yml down -v
```
