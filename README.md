# Auth Service

A production-ready microservice with HTTP and gRPC support built with Go, featuring automatic HTTP-to-gRPC gateway, OpenAPI/Swagger documentation generation, and clean architecture.

## Features

- **Dual Protocol Support**: Both gRPC and HTTP REST APIs with automatic conversion
- **gRPC Gateway**: Auto-generated HTTP REST API from gRPC definitions
- **OpenAPI/Swagger Documentation**: Complete API documentation with interactive UI
- **Clean Architecture**: Separated concerns with service, server, and proto layers
- **Environment Configuration**: Flexible port and settings management
- **Cross-Platform**: Build scripts for Windows, Linux, and Mac
- **Production Ready**: Graceful shutdown, error handling, and logging

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

3. **Generate protobuf files, gRPC Gateway, and Swagger documentation:**
   
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
├── api/                       # API definitions
│   └── proto/                 # Protocol Buffer definitions
│       ├── libs/              # Third-party proto dependencies (auto-downloaded)
│       ├── auth_api.proto     # Service definitions with OpenAPI annotations
│       └── auth_payload.proto # Request/Response message definitions
├── pkg/                       # Generated code packages
│   └── protobuf/             # Generated Go files from proto
│       ├── auth_api.pb.go    # Proto messages
│       ├── auth_api_grpc.pb.go # gRPC server/client
│       └── auth_api.pb.gw.go # HTTP gateway (auto-generated)
├── internal/                  # Internal application code
│   ├── service/              # Business logic layer
│   │   └── auth_service.go   # Authentication service implementation
│   └── server/               # Server management layer
│       └── server.go         # HTTP + gRPC server orchestration
├── docs/                     # Generated documentation
│   ├── auth_api.swagger.json # API documentation file
│   └── auth_payload.swagger.json # Payload documentation
├── .env.example              # Environment variables template
├── .env                      # Your local environment (gitignored)
├── .gitignore               # Git ignore rules
├── generate.bat             # Windows generation script
├── generate.sh              # Linux/Mac generation script
├── go.mod                   # Go module definition
└── main.go                  # Application entry point
```

## Architecture

### Clean Architecture Layers

1. **API Layer** (`api/proto/`): Protocol definitions and contracts
2. **Service Layer** (`internal/service/`): Business logic and domain rules
3. **Server Layer** (`internal/server/`): Infrastructure and server management
4. **Generated Layer** (`pkg/`): Auto-generated code from proto definitions

### gRPC Gateway Approach

This service uses **gRPC Gateway** to automatically generate HTTP REST endpoints from gRPC service definitions. Benefits:

- ✅ **Single source of truth**: Proto files define both gRPC and HTTP APIs
- ✅ **Automatic conversion**: HTTP requests are converted to gRPC calls
- ✅ **Type safety**: Shared message types between protocols
- ✅ **Documentation**: OpenAPI/Swagger generated from annotations

## API Documentation

After running the generation scripts, you can access the API documentation:

- **Swagger JSON**: `./docs/auth_api.swagger.json`
- **View Documentation**: 
  - Upload the JSON file to [Swagger Editor](https://editor.swagger.io)
  - Access via HTTP: `http://localhost:8080/docs` (when server is running)
  - Or use any OpenAPI-compatible tool

## Development Workflow

1. **Modify proto files** in the `api/proto/` directory
2. **Add OpenAPI annotations** for HTTP endpoints and documentation
3. **Run generation script** to update Go code and documentation:
   ```bash
   .\generate.bat  # Windows
   ./generate.sh   # Linux/Mac
   ```
4. **Implement business logic** in `internal/service/`
5. **Test your changes** by running the service

## Default Configuration

- **gRPC Server**: `localhost:9090`
- **HTTP Gateway**: `localhost:8080`
- **Health Endpoint**: `GET /api/v1/health`
- **API Documentation**: `GET /docs`

You can override these settings using environment variables in your `.env` file:

```env
GRPC_PORT=9090
HTTP_PORT=8080
```

## Available Endpoints

### Health Check
- **gRPC**: `AuthService.Health`
- **HTTP**: `GET /api/v1/health`
- **Response**: Service status and health information

## Development Features

### Auto-Generation
- **Protobuf compilation**: Go structs and gRPC code
- **HTTP Gateway**: REST endpoints from gRPC definitions  
- **OpenAPI Documentation**: Interactive API documentation
- **Type-safe clients**: Both gRPC and HTTP clients

### Production Features
- **Graceful shutdown**: Proper cleanup on SIGINT/SIGTERM
- **Concurrent servers**: gRPC and HTTP running simultaneously
- **Error handling**: Proper error propagation and logging
- **Health checks**: Built-in health monitoring

## Adding New Endpoints

1. **Define RPC in proto file** (`api/proto/auth_api.proto`):
   ```protobuf
   rpc Login(LoginRequest) returns (LoginResponse) {
     option (google.api.http) = {
       post: "/api/v1/auth/login"
       body: "*"
     };
   }
   ```

2. **Define messages** (`api/proto/auth_payload.proto`):
   ```protobuf
   message LoginRequest { ... }
   message LoginResponse { ... }
   ```

3. **Run generation**: `.\generate.bat`

4. **Implement in service** (`internal/service/auth_service.go`):
   ```go
   func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
     // Business logic here
   }
   ```

## Author

This repository was created by **Rizki Ramdhani** with assistance from **GitHub Copilot** AI.
