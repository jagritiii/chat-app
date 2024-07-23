# Stage 1: Build
FROM golang:1.20 AS builder

# Set working directory
WORKDIR /app

# Copy go mod file
COPY go.mod ./

# Download all dependencies and generate go.sum
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Run
FROM alpine:latest  

# Set working directory
WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"]