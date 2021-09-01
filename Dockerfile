# Start from the latest golang base image
FROM golang:latest

# Set Working DIR
WORKDIR /app


# Copy GO modules and expected hashes file
COPY app/. .

# # Install dependencies
RUN go mod download

# COPY . .

# # Expose port 5000 
ENV PORT 5000

# # Build
RUN go build src/main.go

ENTRYPOINT [ "./main" ]

 