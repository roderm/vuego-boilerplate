
build: proto rice
	## TODO: go-binary

.PHONY: rice
rice:
	cd ./web && npm run build
	cd ricebox && rice embed-go;

.PHONY: proto
proto:
	find ./api/ -type f -name *.proto -exec \
        protoc \
        --proto_path=${GOPATH}/src:. \
        --go_out=plugins=grpc:${GOPATH}/src/ \
		--js_out=import_style=commonjs:./web/ \
    	--grpc-web_out=import_style=commonjs,mode=grpcweb:./web/ \
    {} \;

dev:
	go run cmd/main.go develop