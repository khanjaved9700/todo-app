FROM golang:1.24.4

WORKDIR /app

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Now copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o main main.go

EXPOSE 8080

CMD ["./main"]
