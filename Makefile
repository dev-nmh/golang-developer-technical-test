# ==============================================================================
# Main

run:
	go run ./cmd/web/main.go

build:
	go build ./cmd/web/main.go


run-linter:
	echo "Starting linters"
	golangci-lint run ./...