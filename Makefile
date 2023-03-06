DB_URL=postgresql://root:secret@localhost:5431/e_commerce?sslmode=disable

network:
	docker network create commerce-network

postgres:
	docker run --name postgres_e --network commerce-network -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres_e createdb --username=root --owner=root e_commerce

dropdb:
	docker exec -it postgres dropdb e_commerce

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

mock:
	mockgen -package mockdb -destination db/mock/store.go project/e-commerce/db/sqlc Store


.PHONY: network postgres createdb dropdb migrateup migrateup1 migratedown mock