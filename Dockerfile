# Stage 1: Build Svelte frontend
FROM node:20 AS frontend-builder

WORKDIR /app
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ ./
RUN npm run build

# Stage 2: Build Go backend
FROM golang:1.21 AS backend-builder

WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ ./

# Copy the built frontend from the previous stage
COPY --from=frontend-builder /app/build ./frontend

# Build the Go app
RUN go build -o server .

# Stage 3: Production image
FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=backend-builder /app/server .
COPY --from=backend-builder /app/frontend ./frontend

EXPOSE 8080

ENTRYPOINT ["/app/server"]
