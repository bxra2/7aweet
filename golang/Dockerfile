FROM golang:1.23

WORKDIR /app

# Install sqlite3 headers for CGO
RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Enable CGO when building/running
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

EXPOSE 5000
CMD ["go", "run", "./main.go"]
