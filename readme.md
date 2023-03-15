go run main.go
go build main.go -> gera o bin
go mod tidy -> baixa o pacote, equivalente ao npm i
go test ./...

criar banco de dados:
sqlite3 orders.db

subir docker
entrar no bash
http://localhost:9021
go run cmd/consumer/main.go