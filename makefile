all: run test test_coverage

run:
	go run main.go

test:
	go test ./...

test_coverage:
	 go test ./... -coverprofile=coverage.out