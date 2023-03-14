package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/AdrianeRibeiro/GoIntensivo/internal/infra/database"
	"github.com/AdrianeRibeiro/GoIntensivo/internal/usecase"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/genproto/googleapis/spanner/admin/database/v1"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := database.NewOrderRepository(db)
	usecase := usecase.CalculateFinalPrice{OrderRepository: repository}

	msgChanKafka := make(chan *ckafka.Message)

	topics := []string{"orders"}
	servers := ""
}

func kafkaWorker(msgChan chan *ckafka.Message, uc usecase.CalculateFinalPrice) {
	for msg := range msgChan {
		var OrderInputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Value, &OrderInputDTO)

		if err != nil {
			panic(err)
		}

		outputDto, err := uc.Execute(OrderInputDTO)
		if err != nil {
			panic(err)
		}

		fmt.Printf("kafka has processed order %s \n", outputDto.ID)
	}
}
