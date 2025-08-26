@echo off
if not exist "pkg\protobuf" mkdir pkg\protobuf
if not exist "docs\" mkdir docs\

echo Generating Go code from proto files...
protoc --proto_path=./api/proto ^
    --proto_path=./api/proto/libs ^
    --plugin=%GOPATH%\bin\protoc-gen-go.exe ^
    --plugin=%GOPATH%\bin\protoc-gen-go-grpc.exe ^
    --plugin=%GOPATH%\bin\protoc-gen-grpc-gateway.exe ^
    --plugin=%GOPATH%\bin\protoc-gen-openapiv2.exe ^
    --go_out=./pkg/protobuf --go_opt=paths=source_relative ^
    --go-grpc_out=./pkg/protobuf --go-grpc_opt=paths=source_relative ^
    --grpc-gateway_out=./pkg/protobuf ^
    --grpc-gateway_opt=allow_delete_body=true,logtostderr=true,paths=source_relative,repeated_path_param_separator=ssv ^
    --openapiv2_out=./docs ^
    --openapiv2_opt=logtostderr=true,repeated_path_param_separator=ssv ^
    ./api/proto/auth_api.proto ./api/proto/auth_payload.proto
echo Swagger JSON available at: ./docs/auth_api.swagger.json
