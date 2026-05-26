FROM node:20-alpine AS frontend
WORKDIR /svelte
COPY svelte/package*.json ./
RUN npm install
COPY svelte/ ./
RUN mkdir -p /golang/public && npm run build

FROM golang:1.23 AS backend
WORKDIR /app
RUN apt-get update && apt-get install -y --no-install-recommends sqlite3 libsqlite3-dev \
    && rm -rf /var/lib/apt/lists/*
COPY golang/go.mod golang/go.sum ./
RUN go mod download
COPY golang/ ./
ENV CGO_ENABLED=1
RUN go build -o /app/server ./main.go

FROM debian:bookworm-slim
WORKDIR /app
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates libsqlite3-0 \
    && rm -rf /var/lib/apt/lists/*
COPY --from=backend /app/server /app/server
COPY --from=backend /app/words.db /app/words.db
COPY --from=frontend /golang/public /app/public
EXPOSE 5000
CMD ["/app/server"]
