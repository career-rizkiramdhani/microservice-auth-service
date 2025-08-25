# Auth Service

A simple microservice with HTTP and gRPC support built with Go.

## Author

This repository was created by **Rizki Ramdhani** with assistance from **GitHub Copilot** AI.

## Getting Started

### Prerequisites
- Go 1.19 or higher
- Protocol Buffers compiler (protoc)

### After Cloning

1. **Install Go dependencies:**
   ```bash
   go mod download
   ```

2. **Install protoc plugins (if not already installed):**
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

3. **Generate protobuf files:**
   
   **Windows:**
   ```bash
   .\generate.bat
   ```
   
   **Linux/Mac:**
   ```bash
   chmod +x generate.sh
   ./generate.sh
   ```

## Development

- Modify `.proto` files in the `proto/` directory
- Run generation script after proto changes
- Generated files will be placed in `server/protobuf/`
