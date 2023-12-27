setup:
	go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-up:
	migrate -path database/migration -database "mysql://habi:habi123@tcp(localhost:3306)/vftalk" -verbose up

migrate-down:
	migrate -path database/migration -database "mysql://habi:habi123@tcp(localhost:3306)/vftalk" -verbose down

build:
	go build -o vftalk