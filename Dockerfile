FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app

COPY . /app

RUN go mod tidy
RUN go build -o main /app/cmd/main.go