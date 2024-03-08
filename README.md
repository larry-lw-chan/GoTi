# GoTi
GoTi is a social networking app written in Go designed to be light and easy to host.  

## How to run...

### First: Clone project and cd into project directory
```
git clone git@github.com:larry-lw-chan/GoTi.git
cd goti
```

### To Run via Docker
Just type:
```
docker-compose up
```

### To Run via Go directly
Alternatively, if you have [Go](https://go.dev/) and [Make](https://www.gnu.org/software/make/manual/make.html) installed, just type
```
make migrate-up
make run
```

### Todo - Binary Builds
For future release, release binary builds...

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

Goti also makes heavy use of [SQLC](https://sqlc.dev/) for generating database query code.  The following make shortcut below allows easy SQLC generation.
```
make package=users sqlc
```
