
build: proto rice
	## TODO: go-binary

.PHONY: rice
rice:
	cd ./web && npm run build
	cd ricebox && rice embed-go;

.PHONY: proto
proto:
	go get -u github.com/golang/protobuf/protoc-gen-go
	## TODO: js proto
	find ./api/ -type f -name *.proto -exec \
        protoc \
        --proto_path=${GOPATH}/src:. \
        --go_out=plugins=${GOPATH}/src/ \
    {} \;

dev:
	go run cmd/main.go develop