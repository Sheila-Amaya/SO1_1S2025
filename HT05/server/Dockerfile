FROM golang:1.21-alpine

WORKDIR /app

COPY proto ./proto
COPY go.mod go.sum ./
COPY server/main.go .

RUN go mod tidy
RUN go build -o server main.go

CMD ["./server"]
