@echo off
if not exist "server\" mkdir server\
if not exist "docs\" mkdir docs\
echo Generating Go code from proto files...
protoc --go_out=./server/ --go-grpc_out=./server/ --proto_path=proto --proto_path=proto/libs proto/*.proto
echo Generating Swagger/OpenAPI documentation...
protoc --openapiv2_out=./docs/ --openapiv2_opt=logtostderr=true --proto_path=proto --proto_path=proto/libs proto/auth_api.proto
echo Generation completed!
echo Swagger JSON available at: ./docs/auth_api.swagger.json
