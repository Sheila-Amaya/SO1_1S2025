# builder
FROM golang:1.21-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum vendor/ ./
ENV GOFLAGS="-mod=vendor"
COPY . .
RUN go build -o kafka-consumer .

# final
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
WORKDIR /root/
COPY --from=builder /app/kafka-consumer .
CMD ["./kafka-consumer"]
