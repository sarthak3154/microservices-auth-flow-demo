FROM golang:1.12.9-alpine

# Install go protoc packages
RUN apk update && apk add --no-cache git && \
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && \
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && \
    go get -u github.com/golang/protobuf/protoc-gen-go && \
    go get -u google.golang.org/grpc

ENV GITHUB_REP_PATH=github.com/sarthak3154/microservices-auth-flow-demo
RUN mkdir -p /go/src/${GITHUB_REP_PATH}/backend/protos

# Copy the probuf go stubs and reverse proxy
COPY protos/*.go /go/src/${GITHUB_REP_PATH}/backend/protos/

COPY main.go /usr/src/

EXPOSE 80

CMD [ "go", "run", "/usr/src/main.go" ]