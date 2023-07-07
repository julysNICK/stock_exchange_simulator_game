DB_URL_DEV=postgresql://root:secret@localhost:5432/stock_exchange?sslmode=disable

sqlc:
	sqlc generate

server:
	go run main.go
network:
	sudo docker network create stock_exchange_network

postgres:
	sudo docker run --name postgres_stock_exchange --network stock_exchange_network -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/stock_exchange?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/stock_exchange?sslmode=disable" -verbose down

createdb:
	sudo docker exec -it postgres_stock_exchange createdb --username=root --owner=root stock_exchange


dropdb:
	sudo docker exec -it postgres_stock_exchange dropdb stock_exchange


initdocker:
	sudo systemctl start docker && sudo docker start postgres_stock_exchange

stopdocker:
	sudo systemctl stop docker && sudo docker stop postgres_stock_exchange


startdocker:
	sudo systemctl start docker && sudo docker start postgres_stock_exchange

	
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/julysNICK/stock_exchange_simulator_game/db/sqlc StoreDB



.PHONY: sqlc server network postgres createdb dropdb initdocker stopdocker migrateup migratedown
