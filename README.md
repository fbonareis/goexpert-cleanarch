# GoExpert - Clean Arch

## Startup

### Pre-run
Before starting the application, ensure that Docker is installed on your system. Then, run the following command to start necessary services in Docker containers:

```shell
docker-compose up -d
```

### Run application
To start the application, navigate to the `cmd/ordersystem` directory and execute the following command to start the application:
```shell
cd cmd/ordersystem
go run main.go ./wire_gen.go  
```

## Testing

### GraphQL
Access `http://localhost:8080` in your browser to interact with the GraphQL API. Ensure you have a basic understanding of GraphQL queries and mutations for effective testing.

### REST
Install the [REST Client externsion](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) in VSCode to test the REST API. Use the `create_order.http` and `list_order.http` files located in the `/api` folder to perform tests. Familiarize yourself with HTTP request methods (GET, POST, etc.) and RESTful API concepts.

### gRPC

## Evans client
To test gRPC service calls locally, install the [Evans Client](https://github.com/ktr0731/evans) and run the following command:
```shell
evans -r repl
```

## Development

### gRPC

## Protobuf
To compile the Protobuf files, use the following command:
```shell
protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
```
Make sure to install the Protocol Buffers compiler (protoc) and the Go protocol buffers plugin (protoc-gen-go) before running this command.



