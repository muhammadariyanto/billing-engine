run-http:
	go run main.go serveHttp --config .config.yaml

test:
	go test -count=1 -race -coverprofile=coverage.out ./...