# Stage 1: Build
FROM golang:alpine as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ottotest_backend ./cmd/api

# Stage 2: Run
FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/ottotest_backend .
COPY migrations ./migrations
COPY docs ./docs
EXPOSE 8080

CMD ["./ottotest_backend"]
