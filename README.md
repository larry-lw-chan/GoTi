# GoTi
GoTi is a social networking app written in Go designed to be light and easy to host.  

## How to run...

### Via Docker
Assuming you have Docker running on the host system, just cd into the project directory
and type:
```
docker-compose up
```

### Via Go
Alternatively, if you have [Go](https://go.dev/) installed, just type
```
go run ./cmd/web/main.go
```

## Development Mode
To run the development environment, simply install [Air](https://github.com/cosmtrek/air)
and type the following command inside the project root directory:
```
air
```

## Database & Migration
Goti uses a forked version of [Goose](https://github.com/pressly/goose) for database migrations.  Migration files are located in db folder.  SQlite3 is currently the default database.  To generate a develoopment database, type

```
make db
```

Other development related database commands include

```
make migrate-up
```

```
make migrate-down
```

```
make seed
```

Goti also makes heavy use of [SQLC](https://sqlc.dev/) for generating database query code.
