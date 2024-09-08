# ==============================================================================
# Main

run:
	go run ./cmd/web/main.go

build:
	go build ./cmd/web/main.go

run-test:
	go test -v ./test/ 

#user:password@tcp(host:port)/dbname?multiStatements=true
run-migration:
	migrate -path db/migration/ -database 'mysql://${user}:${password}@tcp(${host}:${port})/${dbname}?multiStatements=true' -verbose up