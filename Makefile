# ==============================================================================
# Main

run:
	go run ./cmd/web/main.go

build:
	go build ./cmd/web/main.go

run-test:
	go test -v ./test/ 