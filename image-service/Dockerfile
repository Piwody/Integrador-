FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod tidy && go mod download

COPY . .
RUN go mod tidy && go build -o image-api ./cmd/server/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/image-api .
RUN mkdir -p uploads
EXPOSE 5000
CMD ["./image-api"]

