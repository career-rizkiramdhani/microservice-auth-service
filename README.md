# Auth Service

A simple microservice with HTTP and gRPC support built with Go, featuring OpenAPI/Swagger documentation generation.

## Features

- **Dual Protocol Support**: Both gRPC and HTTP REST APIs
- **OpenAPI/Swagger Documentation**: Auto-generated API documentation
- **Environment Configuration**: Flexible port and settings management
- **Cross-Platform**: Build scripts for Windows, Linux, and Mac
- **Clean Architecture**: Separated proto definitions and generated code

## Getting Started

### Prerequisites
- Go 1.19 or higher
- Protocol Buffers compiler (protoc)
- Git (for cloning dependencies)

### After Cloning

1. **Install Go dependencies:**
   ```bash
   go mod download
   ```

2. **Install required protoc plugins:**
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
   go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
   ```

3. **Generate protobuf files and Swagger documentation:**
   
   **Windows:**
   ```bash
   .\generate.bat
   ```
   
   **Linux/Mac:**
   ```bash
   chmod +x generate.sh
   ./generate.sh
   ```

4. **Configure environment (optional):**
   ```bash
   cp .env.example .env
   # Edit .env file to customize ports if needed
   ```

5. **Run the service:**
   ```bash
   go run main.go
   ```

## Project Structure

```
├── docs/                      # Generated Swagger/OpenAPI documentation
│   └── auth_api.swagger.json  # API documentation file
├── proto/                     # Protocol Buffer definitions
│   ├── libs/                  # Third-party proto dependencies (auto-downloaded)
│   ├── auth_api.proto         # Service definitions with OpenAPI annotations
│   └── auth_payload.proto     # Request/Response message definitions
├── server/                    # Generated Go code
│   ├── protobuf/             # Generated files from proto
│   └── server/               # Additional server implementations
├── internal/                  # Internal application code
│   └── handlers/             # HTTP and gRPC handlers (when implemented)
├── .env.example              # Environment variables template
├── .env                      # Your local environment (gitignored)
├── .gitignore               # Git ignore rules
├── generate.bat             # Windows generation script
├── generate.sh              # Linux/Mac generation script
├── go.mod                   # Go module definition
└── main.go                  # Application entry point
```

## API Documentation

After running the generation scripts, you can access the API documentation:

- **Swagger JSON**: `./docs/auth_api.swagger.json`
- **View Documentation**: 
  - Upload the JSON file to [Swagger Editor](https://editor.swagger.io)
  - Or use any OpenAPI-compatible tool

## Development Workflow

1. **Modify proto files** in the `proto/` directory
2. **Run generation script** to update Go code and documentation:
   ```bash
   .\generate.bat  # Windows
   ./generate.sh   # Linux/Mac
   ```
3. **Implement handlers** in the `internal/handlers/` directory
4. **Test your changes** by running the service

## Default Configuration

- **HTTP Server**: `localhost:8080`
- **gRPC Server**: `localhost:9090`
- **Health Endpoint**: `GET /api/v1/health` or `GET /api/health`

You can override these settings using environment variables in your `.env` file.

## Author

This repository was created by **Rizki Ramdhani** with assistance from **GitHub Copilot** AI.
