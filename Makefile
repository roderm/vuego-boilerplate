dev:
	cd ./web && npm run dev
	go run main.go

build:
	cd ./web && npm run build;
	cd ricebox; rice embed-go;
	go build