# Development Environment

## Local Environment  
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

## Cloud Environment (with Cloud Endpoints)

### Docker Build Configuration

1. Build the docker image:
    
       # Make sure you are in the base directory
       $ docker build -t username/repo-name[:TAG] -f docker/server.Dockerfile .
       
2. Push the docker image to the Dockerhub repository.
    
       # Make sure you are logged in to docker with docker login command
       $ docker push username/repo-name[:TAG]

### Cloud Endpoints Configuration Deployment

 1. Make sure you have separate files for gRPC APIs and gRPC APIs with HTTP/JSON Transcoding as `<file_name>.proto` and 
 `http_<file_name>.proto`.
 2. Compile the proto file using protoc compiler.
 
        $ protoc \
          --include_imports \ 
          --include_source_info \
          protos/<file_name>.proto \
          --descriptor_set_out api.pb
          
  3. Update the `PROJECT_ID` in the `api_config.yaml` file, or update the file to match your own requirements.
  4. Deploy your service's configuration to Endpoints
  
         $ gcloud endpoints services deploy api.pb api_config.yaml
         
         
### GKE Service Deployment

1. Prepare the `deployment` and `service` `.yaml` files as present in the `deployment` folder in the base directory.
2. Deploy the API backend by setting up the Kubernetes cluster on GKE.
        
     2.1 - Create a Kubernetes cluster by visiting the GCP console. 
     
     2.2 - From the GCP navigation menu, follow `Kubernetes Engine` > `Clusters` > `Create Cluster`
     to create the cluster.  
       
     2.3 - Authenticate `kubectl` to the container cluster.
       
       $ gcloud container clusters get-credentials <CLUSTER_NAME> --zone <ZONE>

     2.4 - Deploy the API Backend to the Kubernetes cluster.
     
       $ kubectl apply -f /path/to/deployment-file.yml
       $ kubectl apply -f /path/to/service-file.yml
       
 3. Test the gRPC Endpoints. 
 
    3.1 - Get the `External IP` and `Port` by running:
        
        $ kubectl get services
        
    3.2 - Hit the Endpoints from a gRPC Client (say, BloomRPC)
    
### References

1. [Cloud Endpoints - Transcoding HTTP/JSON to gRPC](https://cloud.google.com/endpoints/docs/grpc/transcoding)
2. [Cloud Endpoints - gRPC Service Configuration](https://cloud.google.com/endpoints/docs/grpc/grpc-service-config)