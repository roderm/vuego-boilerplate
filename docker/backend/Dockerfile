FROM roderm/go-protoc

ENV PKG_PATH "github.com/roderm/vuego-boilerplate"

ENV PROTOC_WEB https://github.com/grpc/grpc-web/releases/download/1.0.7/protoc-gen-grpc-web-1.0.7-linux-x86_64
## deps
RUN go get github.com/GeertJohan/go.rice/rice && \
    go get -u github.com/golang/protobuf/protoc-gen-go

COPY . $GOPATH/src/${PKG_PATH}
WORKDIR $GOPATH/src/${PKG_PATH}
## generate proto
RUN find ./api/ -type f -name *.proto -exec \
        protoc \
        --proto_path=${GOPATH}/src:. \
        --go_out=plugins=${GOPATH}/src/ \
    {} \;

RUN curl ${PROTOC_WEB} -o /usr/local/bin/protoc-gen-grpc-web && \
    chmod +x /usr/local/bin/protoc-gen-grpc-web
## generate static folder
RUN cd web && npm install && \
    find ./api/ -type f -name *.proto -exec \
    protoc \
    --proto_path=${GOPATH}/src:. \
    --js_out=import_style=commonjs:./src/proto \
    --grpc-web_out=import_style=commonjs,mode=grpcweb:./src/proto \
    {} \;
## go-rice

## build binary

## copy to scratch