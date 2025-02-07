FROM golang:1.23-alpine AS go-builder

WORKDIR /back

# Bass boost
RUN apk update && apk add ffmpeg

# Copy only go.mod and go.sum to leverage caching
COPY /back/go.mod /back/go.sum ./
COPY /back/go-shared/go.mod /back/go-shared/go.sum ./go-shared/

# Download Go dependencies
RUN go mod download

# Copy the rest of the code
COPY back .

# Generate Client OpenAPI
RUN go generate api_client/generate.go

# Generate Server OpenAPI
RUN go generate api_server/generate.go

# Build the Go application
RUN go build -o main .

# Run the Go application
CMD ["./main"]
