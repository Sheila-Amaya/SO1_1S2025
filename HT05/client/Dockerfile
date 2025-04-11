FROM golang:1.21-alpine

WORKDIR /app

COPY proto ./proto
COPY go.mod go.sum ./
COPY client/main.go .

RUN go mod tidy
RUN go build -o client main.go

CMD ["./client"]
