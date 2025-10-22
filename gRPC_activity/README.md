* Create a routeguide directory in gRPC_activity
* Write a routeguide.proto file containing the service definition, along with the syntax declaration and the option go_package (which specifies the path where the generated Go files will be stored).
* Run following command to generate Go source files from the .proto file:
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    routeguide.proto
* This command generates two files in the directory specified in the .proto file:
  - routeguide.pb.go - contains all message (request and response) type definitions as Go structs.
  - routeguide_grpc.pb.go — contains:
   - A client interface (stub) for invoking RPC methods defined in the RouteGuide service.
   - A server interface for implementing those methods on the server side.
* Create a server package with server.go, implementing all RPC methods from the service.
* Create a client package with client.go, invoking the RPC methods exposed by the server.
* Run the following commands in separate terminals:
  - go run server/server.go
  - go run client/client.go 
* The client communicates with the server via gRPC and prints the received responses.
