# hgz-protobuffer-comparison

Sample project to compare Protobuffer versus JSON.

* [Protocol Buffer](https://developers.google.com/protocol-buffers)
* [JSON](https://www.json.org/json-en.html)

## Configure environment variable

To run the project you need to install the **GNU Make**, **protobuffer compiler** and **go code generator**.

* [Protobuffer compiler](https://developers.google.com/protocol-buffers/docs/downloads)
* Go code generator
  * `go install google.golang.org/protobuf/cmd/protoc-gen-go`
* [GNU Make](https://www.gnu.org/software/make/)
    * The way to install will depend on your system environment 

## Commands

* `make compile-protoc-win`
  * For that, you need to have `$PROTOC` name as an environment variable to **protoc**    
  * generates go code from proto file to Windows

* `make compile-protoc-macos`
    * generates go code from proto file to MacOS