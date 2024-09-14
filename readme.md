# Co-working-space-management

## Pre-requisites
1. protoc
   1. download zip file from [here](https://github.com/protocolbuffers/protobuf/releases)
   2. extract the zip file
   3. configure the path to the bin folder in the extracted folder
   4. run `protoc --version` to verify the installation (restart the terminal if it doesn't work)
2. protoc-gen-go
   1. run `go get -u github.com/golang/protobuf/protoc-gen-go`

## To run the project
1. Clone the project
2. Copy the `.env.example` file to `.env` and fill in the required fields
3. Run `docker compose up --build -d` to start the project