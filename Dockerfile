# Use the official Golang image
FROM golang:1.23


LABEL org.opencontainers.image.source https://github.com/mikej81/mini-grpc

# Set up the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the gRPC server
RUN go build -o server .

# add grpcurl for local testing
RUN curl -sSL "https://github.com/fullstorydev/grpcurl/releases/download/v1.8.7/grpcurl_1.8.7_linux_x86_64.tar.gz" | tar -xz -C /usr/local/bin

# Expose the server's port
EXPOSE 50051

# Run the server
CMD ["./server"]
