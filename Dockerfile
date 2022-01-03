# syntax=docker/dockerfile:1

FROM golang:alpine AS builder

WORKDIR /go/src/app
COPY . ./

RUN go build -o book_store ./cmd/myapp.go

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/app/ /app/

CMD [ "/book_store" ]