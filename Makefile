run:
	go run cmd/main.go
up:
	docker compose up -d
down:
	docker compose down
rm:	#remove old data of volume
	docker volume rm manager_postgres_data
exec:
	docker exec -it postgres_manager psql -U postgres -d db