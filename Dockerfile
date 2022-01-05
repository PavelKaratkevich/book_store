# syntax=docker/dockerfile:1

FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/myapp.go

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/app/ /app/

CMD ["./app"]

# docker run -dp 8080:8080 book_store
