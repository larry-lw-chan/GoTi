# Specify a base image
FROM golang:1.21.7

WORKDIR '/app'

# Install some dependencies
ADD https://github.com/pressly/goose/releases/download/v3.7.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose
COPY go.mod .
RUN go mod download
COPY . .

# Default Command
CMD ["make", "up-and-run"]