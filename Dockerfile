# syntax=docker/dockerfile:1

FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/
COPY .env /app/

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/myapp.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/ /app/

EXPOSE 8080
EXPOSE 5432

CMD ["./app"]