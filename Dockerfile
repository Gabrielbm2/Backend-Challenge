FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

ENTRYPOINT ["/app/server"]
