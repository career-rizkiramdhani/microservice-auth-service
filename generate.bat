@echo off
if not exist "pkg\" mkdir pkg\
if not exist "docs\" mkdir docs\
echo Generating Go code from proto files...
protoc --go_out=./pkg/ --go-grpc_out=./pkg/ --proto_path=api/proto --proto_path=api/proto/libs api/proto/*.proto
echo Generating Swagger/OpenAPI documentation...
protoc --openapiv2_out=./docs/ --openapiv2_opt=logtostderr=true --proto_path=api/proto --proto_path=api/proto/libs api/proto/auth_api.proto
echo Generation completed!
echo Swagger JSON available at: ./docs/auth_api.swagger.json
