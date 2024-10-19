FROM golang:1.20.0-alpine

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN go build -o bin .

ENTRYPOINT ["/app/bin"]
