# Introduction to gRPC

Here are an implementation of a NodeJs client and a Go server.

1) First, create in `pb` path the messeges model based on buffer protocol `messages.proto`
2) Compile the model to a go implementation
   `$ protoc pb/messages.proto --go_out=plugins=grpc:/home/alex/go/src`
3) Instantiate a server to list the port `50000` and return some response data
4) Run the server
5) Instantiate the node client to send data using a method
6) Run the client

## Go dependencies

Install this dependencies
```
$ go get golang.org/x/net/context
$ go get google.golang.org/grpc
```

## Node dependencies

Install this dependencies
```
$ npm install grpc
```