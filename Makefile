test:
	@go test -v -race -coverprofile=coverage.txt -covermode=atomic ./pkg/middleware
