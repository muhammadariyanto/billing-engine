run-http:
	go run main.go serveHttp --config .config.yaml

test:
	go test -race ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

gen-mock-repo:
	mockery --all --dir ./internal/repository/ --output ./mocks/repository --outpkg mocks

gen-mock-service:
	mockery --all --dir ./internal/service/ --output ./mocks/service --outpkg mocks
