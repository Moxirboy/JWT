.SILENT:
build:
    go mod dowload && CGO_ENABLED=0 GOOS=linux go build -o ./cmd
run: build
   docker -compose up --remove-orphans --build server

migrate:
    migrate --path ./migrations/postgres --database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up
migrate-down:
      migrate --path ./migrations/postgres --database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' down 1
