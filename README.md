# Manager
A manager tasks service build with Go, following clean architecture principles and microservice patterns

## 1. Tech stack
### Core framework and language
- Go 1.23+: Primary programming language
- Gin: HTTP web framework for Restful APIs
- Postgres 15: Primary relational database
- GORM: ORM library for PostgresSQL database operations
- Zaplog: Structured, high performance logging library
- godotenv: Load environment variables from .env files
- env: Parse environment variables into structs
- DBeaver 25.2.5: Free Open-Source Database Management Tool
- Postman: Platform supports build APIs.

## 2. How to run project
### Pre-requisites
- Go 1.23+
- Docker and docker compose
### Step 1: Clone the repository
```
git clone <repo-url>
cd manager
```
### Step 2: Install Go dependencies
```
go mod download
```
### Step 3: Start Infrastructure Services
Start all required infrastructure services (PostgresSQL, maybe Redis, Prometheus, Grafana) using docker compose
```
make up
# or
docker-compose up -d
```
This will start:
- PostgresSQL on port 3002
Verify services are running:
```
docker ps
```
### Step 4: Configure environment variables
Create a .env file in the project root with the following configuration
```
ENV=local
LOG_TYPE=zap
LOG_LEVEL=debug
LOG_OUTPUT=console
LOG_USE_JSON=false
DB_USER_NAME=postgres
DB_PASSWORD=1
DB_HOST=localhost
DB_PORT=5432
DB_NAME=db
RUN_MODE=dev
PORT=3002
```

### Step 5: Start application services
```
make run
# or
go run cmd/main.go
```
### Step 6: Demo
- Currently, this project complete to create user, login, verify user api. In the future, it'll support for additional features
- Demo files: manager/2025-11-30 17-05-52.mkv
