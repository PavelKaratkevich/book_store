# syntax=docker/dockerfile:1

FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api/myapp.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/.env .
COPY --from=builder /app/internal/repositoryDB/postgresDB/instructions.sql .
COPY --from=builder /app/prometheus/prometheus.yml .


EXPOSE 8080

CMD ["./app"]