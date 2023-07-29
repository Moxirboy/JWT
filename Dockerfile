# Build stage
FROM golang:1.20.1-alpine3.16 AS builder
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project directory into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o JWT.exe ./cmd

# Final stage
FROM alpine:3.16

# Set environmental variables for PostgreSQL
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=qwerty
ENV POSTGRES_DB=postgres

# Install PostgreSQL client (optional, only needed if you want to interact with a remote PostgreSQL server)
RUN apk add --no-cache postgresql-client

# Copy the Go executable from the build stage
COPY --from=builder /app/JWT.exe /JWT.exe

# Expose the desired port for your Go application (if applicable)
EXPOSE 5436

# Expose the PostgreSQL default port (if needed)
EXPOSE 5432

# (Optional) Copy any SQL scripts or initialization files to the container
# COPY init.sql /docker-entrypoint-initdb.d/

# Run the Go application
CMD ["/JWT.exe"]
