## Development Environment

### Setup gRPC Gateway

#### First time dependencies

1. Install the Protocol Buffer compiler.

        $ mkdir tmp
        $ cd tmp
        $ git clone https://github.com/google/protobuf
        $ cd protobuf
        $ ./autogen.sh
        $ ./configure
        $ make
        $ make check
        $ sudo make install

2. Make sure the `$PROTOC_HOME` environment variable is set and added to the `$PATH` environment variable.
3. Download and install the [go library](https://golang.org/dl/).
4. Make sure you have the `$GOBIN` environment variable set and added to the `$PATH` environment variable.
5. Install the protoc plugins:
        
        $ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
        $ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
        $ go get -u github.com/golang/protobuf/protoc-gen-go

#### gRPC Gateway dependencies

To access the protoc generated dependencies (gRPC stubs and reverse-proxy) in the grpc-gateway initializer file, `main.go`, 
Install it as a plugin enabling installation in the go directory's `/src` package:

    $ go get -u github.com/your-username/repo-name

##### Bind Protobuf compiler to gRPC

1. Generate gRPC stubs and reverse proxy using the protoc compiler:

        # Provided the genproto.sh bash file isn't executable
        $ chmod +x genproto.sh
    
        $ ./genproto.sh

2. Update the dependencies in the go directory's `/src` package with the generated stubs and reverse-proxy in the above step:
        
        $ cp -r protos/*.go ${GOPATH}/src/github.com/your-username/repo-name/path/to/protos/directory

### Start the services

From the base directory of the project,

    # Starts the gRPC Server (at port 9090)    
    $ node server.js
    
    # Starts the gRPC Gateway (at port 8080)
    $ go run main.go