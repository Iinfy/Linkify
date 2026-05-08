# linkify

A URL shortening service with click analytics.

## Quick Start

Prerequisites: Docker and Docker Compose.

```bash
docker-compose -f docker-compose.dev.yml up --build
```

Services available at:

| Service | URL |
|---|---|
| Frontend (via Caddy) | http://localhost |
| Backend API (direct) | http://localhost:8090 |
| Frontend (direct) | http://localhost:8080 |
| PostgreSQL | localhost:5432 |

## Configuration

Copy `.env.example` to `.env`:

```bash
cp .env.example .env
```

| Variable | Required | Default | Description |
|---|---|---|---|
| `POSTGRES_URL` | Yes | — | PostgreSQL connection string |
| `CORS_HOST` | Yes | — | Allowed CORS origin (also used for `/not_found` redirect) |
| `PORT` | No | 8090 | Backend port |
| `RATE_LIMIT` | No | 100 | Max requests per IP per minute |

Note: the backend's listen port is hardcoded to `:8090` regardless of the `PORT` env var.

## API

### POST /slink

Create a short link.

```
POST /slink
Content-Type: application/json

{"url": "https://example.com"}
```

Validation:
- Max URL length: 1024 characters
- Must start with `http://` or `https://`

**200** — OK:
```json
{"url": "<12-char-hash>"}
```

**400** — Malformed request body:
```json
{"error": "internal server error"}
```

**422** — Validation error:
```json
{"error": "URL too long"}
```

The same URL always produces the same hash. Duplicates are silently handled (HTTP 200, existing hash returned).

### GET /s/:hash

Redirect to the original URL. Records click analytics (timestamp + user agent).

- **302 Found** — hash exists, redirects to original URL
- **301 Moved Permanently** — hash not found, redirects to `{CORS_HOST}/not_found`

### GET /short/:hash

Get click statistics for a short link.

```json
{"click_count": 42, "url": "https://example.com"}
```

Returns `""` for `url` and `-1` for `click_count` if the hash is not found or the query fails.

## Architecture

```
Internet → Caddy (:80)
  ├── /s/*, /slink*  → Go backend (:8090)
  └── /* (everything else) → Vue SPA (:8080)
```

- **Backend** (`app/`): Go 1.26.2 + Gin framework. In-memory per-IP rate limiter.
- **Database**: PostgreSQL — stores URLs (`urls` table) and click analytics (`stats` table).
- **Frontend** (`web/`): Vue 3 + TypeScript + Vite SPA, built to static files and served by `devforth/spa-to-http`.
- **Proxy**: Caddy reverse-proxies API paths to the backend; all other paths serve the frontend.

## Development

Run backend directly (requires PostgreSQL running):

```bash
cd app && go run main.go
```

Run tests:

```bash
cd app && go test ./...
```

## Stopping

```bash
docker-compose -f docker-compose.dev.yml down
```

To remove volumes (deletes all data):

```bash
docker-compose -f docker-compose.dev.yml down -v
```
