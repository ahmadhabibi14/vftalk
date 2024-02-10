setup:
	go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate:
	migrate create -ext sql -dir database/migration $(state)

migrate-up:
	migrate -path database/migration -database "mysql://habi:habi123@tcp(localhost:3306)/vftalk" -verbose up

migrate-down:
	migrate -path database/migration -database "mysql://habi:habi123@tcp(localhost:3306)/vftalk" -verbose down

build:
	go build -o vftalk

docker-prod:
	docker-compose -f docker-compose.prod.yml up -d

swagger:
	go-swagger3 --module-path . --output docs/swagger.json --schema-without-pkg=true