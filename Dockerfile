# Specify a base image
FROM golang:1.21.7

WORKDIR '/app'

# Install some dependencies
COPY go.mod .
RUN go mod download
COPY . .

# Default Command
CMD ["go", "run", "./cmd/web/main.go"]