# Stage 1 Build
FROM golang:tip-alpine3.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o simplegin cmd/main.go


# Stage 2 Run
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/simplegin .
COPY --from=builder /app/configs/config.yaml ./configs/
EXPOSE 3000
CMD ["./simplegin"]