dev:
	cd ./web && npm run dev
	go run main.go

build:
	cd ./web && npm run build;
	rice embed-go
	go build