test:
	@go test -v -race -coverprofile=profile.out -covermode=atomic ./pkg/middleware
	