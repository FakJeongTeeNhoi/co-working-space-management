# Co-working-space-management

## Pre-requisites
1. protoc
   1. download zip file from [here](https://github.com/protocolbuffers/protobuf/releases)
   2. extract the zip file
   3. configure the path to the bin folder in the extracted folder
   4. run `protoc --version` to verify the installation (restart the terminal if it doesn't work)
2. protoc-gen-go
   1. run `go install github.com/golang/protobuf/protoc-gen-go`

## To run the project
1. Clone the project
2. Copy the `.env.example` file to `.env` and fill in the required fields
3. Run `docker compose up --build -d` to start the project

## To compile the proto file
1. Configure the proto header as shown below
```proto
syntax = "proto3";

package <PACKAGE_NAME>;

option go_package = "generated/<PACKAGE_NAME>";
```
2. Run `protoc --go_out=. --go-grpc_out=. <PATH_TO_PROTO_FILE>` to compile the proto file