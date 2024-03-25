# Description: Makefile for the project

# Compile the app for multiple OS and Platforms
compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 ./cmd/web/* 
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 ./cmd/web/* 
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 ./cmd/web/* 

# Database Migrations, creation and seeding
migrate-up:
	GOOSE_DRIVER=sqlite3 \
	GOOSE_DBSTRING=./database/sqlite3.db \
	GOOSE_MIGRATION_DIR=./database/migrations/ \
	goose up

migrate-down:
	GOOSE_DRIVER=sqlite3 \
	GOOSE_DBSTRING=./database/sqlite3.db \
	GOOSE_MIGRATION_DIR=./database/migrations/ \
	goose down

seed:
	GOOSE_DRIVER=sqlite3 \
	GOOSE_DBSTRING=./database/sqlite3.db \
	GOOSE_MIGRATION_DIR=./database/seed/ \
	goose -no-versioning up

db:
	make migrate-up
	make seed

# Generate SQLC for selected package.  Example: make package=users sqlc
sqlc:
	sqlc generate -f ./internal/$(internal)/sqlc/sqlc.yaml

# Run the app
run:
	go run ./cmd/web/main.go

# Used by Dockerfile to create new db if not exists and run app
migrate-up-then-run:
	make migrate-up
	make run