# go-tiny
A social networking app written in Go designed to be light and cheap to host


## Development Mode
To run the development environment, simply install [Air](https://github.com/cosmtrek/air)
and type the following command inside the installed directory:
```
air
```

## Database & Migration
Goti uses a forked version of [Goose](https://github.com/pressly/goose) for database migrations.  Migration files are located in db folder.  SQlite3 is currently the default database and is also located in the db folder.

```
goose sqlite3 ./sqlite3.db up 
```

```
goose sqlite3 ./sqlite3.db down
```