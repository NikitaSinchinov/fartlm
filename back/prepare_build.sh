# Download dependencies
go mod download

# Generate Client OpenAPI
mkdir -p api_client
go generate api_client/generate.go

# Generate Server OpenAPI
mkdir -p api_server
go generate api_server/generate.go

