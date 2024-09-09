# ==============================================================================
# Main

run:
	go run ./cmd/web/main.go

build:
	go build ./cmd/web/main.go

run-test:
	go test -v ./test/ 

#make run-migration user=root password=root host=localhost port=3306 dbname=golang_developer_technical_test
run-migration:
	migrate -path db/migration/ -database 'mysql://${user}:${password}@tcp(${host}:${port})/${dbname}?multiStatements=true' -verbose up