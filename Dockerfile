# Use the official Golang image
FROM golang:1.19

# Set up the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the gRPC server
RUN go build -o server .

# Expose the server's port
EXPOSE 50051

# Run the server
CMD ["./server"]
