postgres:
	sudo docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	sudo docker exec -it postgres16 createdb --username=root --owner=root db_SeeCV

dropdb:
	sudo docker exec -it postgres16 dropdb db_SeeCV


migrateup:
	migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5432/db_SeeCV?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/db/migration -database "postgresql://root:secret@localhost:5432/db_SeeCV?sslmode=disable" -verbose down

sqlc:
	sqlc generate

run:
	go run cmd/SeeCV/main.go      

test:
	go test -v  -cover ./...

tidy:
	go mod tidy

dockerbuild:
	sudo docker build --build-arg OPENAI_API_KEY=$OPENAI_API_KEY -t seecv:latest .  

apprun:
	sudo docker run --name seecv -e OPENAI_API_KEY=$OPENAI_API_KEY -p 8080:8080 seecv:latest

rma: 
	sudo docker rm -vf $(sudo docker ps -aq)

rmia: 
	sudo docker rmi -f $(docker images -aq)

.PHONY: postgres createdb dropdb migrateup migratedown sqlc run dockerbuild tidy test rma rmia