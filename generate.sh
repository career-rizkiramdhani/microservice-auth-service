#!/bin/bash
mkdir -p pkg/protobuf
mkdir -p docs/

echo "Generating Go code from proto files..."
protoc --proto_path=./api/proto \
    --proto_path=./api/proto/libs \
    --plugin=$GOPATH/bin/protoc-gen-go \
    --plugin=$GOPATH/bin/protoc-gen-go-grpc \
    --plugin=$GOPATH/bin/protoc-gen-grpc-gateway \
    --plugin=$GOPATH/bin/protoc-gen-openapiv2 \
    --go_out=./pkg --go_opt=paths=source_relative \
    --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out=./pkg \
    --grpc-gateway_opt=allow_delete_body=true,logtostderr=true,paths=source_relative,repeated_path_param_separator=ssv \
    --openapiv2_out=./docs \
    --openapiv2_opt=logtostderr=true,repeated_path_param_separator=ssv \
    ./api/proto/auth_api.proto ./api/proto/auth_payload.proto

echo "Updating handler implementations..."
go run scripts/update_handlers.go
echo "Generation completed!"
echo "Swagger JSON available at: ./docs/auth_api.swagger.json"
echo "gRPC Gateway available at: ./pkg/protobuf/auth_api.pb.gw.go"
