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
