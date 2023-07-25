FROM golang:latest
COPY ./ ./
RUN go build -o ./cmd/main .
CMD "./main"
