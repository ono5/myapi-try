# マルチステージビルド
FROM golang:latest AS builder

WORKDIR /app
RUN go install github.com/cosmtrek/air@latest