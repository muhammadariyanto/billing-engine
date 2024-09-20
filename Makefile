run-http:
	go run main.go serveHttp --config .config.yaml

test:
	go test -race ./... -coverprofile=coverage.out && go tool cover -html=coverage.out