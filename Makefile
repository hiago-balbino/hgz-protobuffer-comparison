.PHONY: compile-protoc-win compile-protoc-macos docker-build-up docker-up docker-down docker-ps proto-bench json-bench

## compile-protoc-win:			generate Go code on Windows from proto file
compile-protoc-win:
	$$PROTOC -I=./protobuffer --go_out=./protobuffer ./protobuffer/proto/*.proto

## compile-protoc-macos:			generate Go code on MacOS from proto file
compile-protoc-macos:
	protoc -I=./protobuffer --go_out=./protobuffer ./protobuffer/proto/*.proto

## docker-build-up:			build and start all containers
docker-build-up:
	cd docker; docker-compose up -d --build

## docker-up:				start all containers
docker-up:
	cd docker; docker-compose up -d

## docker-down:				stop all containers
docker-down:
	cd docker; docker-compose down

## docker-ps:				show all containers running
docker-ps:
	cd docker; docker-compose ps

## proto-bench: 				run all benchmark tests to protocol buffer format
proto-bench:
	cd protobuffer; go test ./... -bench=. -benchtime=10s

## json-bench: 				run all benchmark tests to json format
json-bench:
	cd json; go test ./... -bench=. -benchtime=10s

## help:					show this help.
help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'