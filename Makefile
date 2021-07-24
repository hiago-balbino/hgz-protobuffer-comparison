.PHONY: compile-protoc-win compile-protoc-macos

compile-protoc-win:
	$$PROTOC -I=./protobuffer --go_out=./protobuffer ./protobuffer/proto/*.proto

compile-protoc-macos:
	protoc -I=./protobuffer --go_out=./protobuffer ./protobuffer/proto/*.proto