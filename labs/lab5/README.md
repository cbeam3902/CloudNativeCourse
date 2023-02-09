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

## Python version
This part wasn't intended for the lab, but I decided to include a python version that used the gRPC movie API. The python version is just a quick recreation of the Go code. It doesn't include an error checking mainly because I don't know how it's implemented in Python because it's obvious in Go.

To compile the proto code to generate the python code, execute the following line:

    $ pip3 install grpcio-tools
    $ python3 -m grpc_tools.protoc -Imovieapi --python_out=. --pyi_out=. --grpc_python_out=. movieapi/movieapi.proto

(The python code was made with python3 in mind.)

The client and server python code are located in the movieclient and movieserver directories respectively. This repo includes the generated python code from the line above.
