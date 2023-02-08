# Lab 5
This lab introduces gRPC and how to use it to create a server/client movie database.

The github link to the gRPC project is [here](https://github.com/grpc/grpc).
It is mostly written in C/C++, however can be used from a number of different languages.
This lab is mostly based on the official gRPC documentation available at [grpc.io](https://grpc.io/).

## Prerequisites:
- Install Go
- Protocol buffer compiler, protoc (version 3+)
- Go plugins for the protocol complier:
1. Install the protocol compiler plugins for GO


        $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
        $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

2. Update your PATH so the protoc compiler can find the plugins:

        $ export PATH="$PATH:$(go env GOPATH)/bin"
    
## Run the example
1. Clone the repo to get the example code

    $ git clone -b v1.52.0 --depth 1 https://github.com/grpc/grpc-go

2. Change to the example directory

    $ cd grpc-go/examples/helloworld

3. Run the example code
    
    From one terminal execute the following line 

        $ go run greeter_server/main.go

    From another terminal execute the following line and see the output

        $ go run greeter_client/main.go
        Greeting: Hello world

Congrats, gRPC works

## Client-server API with gRPC

The following tutorial will go over how to create the API using gRPC for the movie database.
It's already pre-generated for this repo.

The proto file is in the movieapi directory with the structure of the gRPC message in there.

To compile the code, execute the following line:

    $ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative movieapi/movieapi.proto

This will create the two files in the movieapi directory *movieapi_grpc.pb.go* and *movieapi.pb.go* that contains the server and client code.

To run the server code 

    $ go run movieserver/server.go

To run the client code in a different terminal

    $ go run movieclient/client.go

If you encounter an error while trying to run the server code, run:

    $ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    $ go get -u google.golang.org/grpc

and recompile the protoc.
