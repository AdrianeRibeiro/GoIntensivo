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

sudo docker stop $(sudo docker ps -q)
sudo docker-compose up -d
sudo docker ps
sudo docker-compose exec goapp bash
go mod tidy
go run cmd/consumer/main.go
/go/goapp/cmd/consumer# go build -o consumer
http://localhost:15672/#/connections
http://localhost:9021/clusters