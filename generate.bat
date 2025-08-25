@echo off
if not exist "server\" mkdir server\
protoc --go_out=./server/ --go-grpc_out=./server/ --proto_path=proto proto/*.proto
