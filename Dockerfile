# Stage 1: Building the Go application
FROM golang:1.23.3-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o receipt-processor-solution-api-maaz .

# Stage 2: Create a minimal image to run the Go binary
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/receipt-processor-solution-api-maaz .

EXPOSE 8080

# Command to run the Go binary when the container starts
CMD ["./receipt-processor-solution-api-maaz"]
