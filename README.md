# GoTi
GoTi is a distributed social networking app written in Go, designed to be light, customizable, and free of vendor lock ins.

## Getting Started

### Prerequisites
Make sure you have the following installed before running GoTi:

- Git
- Docker (optional, for Docker setup)
- Go (optional, for direct Go setup)
- Make (optional, for direct Go setup)
- Node.js (optional, for development environment)

## Installation
1. Clone project and cd into project directory
```
git clone git@github.com:larry-lw-chan/GoTi.git
cd goti
```

#### Docker
To run GoTi using Docker, simply execute:
```
docker-compose up
```

#### Direct Go Setup
If you prefer to run GoTi directly, follow these steps:

1. Make sure you have [Go](https://go.dev/) and [Make](https://www.gnu.org/software/make/manual/make.html) installed.
2. Execute the following commands:

```
make migrate-up
make run
```

## Configuration
Goti simplifies configuration by leveraging environment variables. Begin by copying the provided sample.env file into a .env file. Then, customize the variables within to align with your specific requirements.

#### Mac or Linux Terminal

```
cp .env.example .env
```

## Development Mode
To run the development environment, install [Air](https://github.com/cosmtrek/air)
and execute the following command in the project root directory:

```
air
```

## Coding Conventions
GoTi loosely follows BEM (Block, Element, Modifier) naming conventions for HTML/CSS, ensuring consistency and maintainability. All templates, JavaScript, and CSS formatting are standardized using Prettier. Individual page overrides can be done on a per template basis by defining a 'style' or 'script' block.  To install development dependencies, run:

```
npm install
```

## Database & Migration
Goti uses a forked version of [Goose](https://github.com/pressly/goose) for database migrations.  Migration files are located in database folder. To generate a development database, type

```
make db
```

Other development related database commands include:

```
make migrate-up
```

```
make migrate-down
```

```
make seed
```

Goti also makes heavy use of [SQLC](https://sqlc.dev/) for generating database query code.  Use the following Make shortcut for easy SQLC generation (fill ? with appropriate internal package name):

```
make internal=? sqlc
```

## License
This project is licensed under the Server Side Public License - see the LICENSE file for details

## Copyright
(C) 2024 Larry LW Chan

