package consumer

import (
	"fmt"

	"github.com/ilyabukanov123/L0/internal/cache"
	"github.com/ilyabukanov123/L0/internal/db"
	"github.com/ilyabukanov123/L0/internal/model"
	"github.com/ilyabukanov123/L0/internal/server"

	"github.com/nats-io/stan.go"
)

const (
	port      = ":4223"        // Порт nast-streaming на котором работает клиент
	clusterID = "test-cluster" // Настройка по умолчанию
	clientID  = "consumer"     // ID клиента
)

// type OrderModel struct {
// 	Uid  string
// 	Json OrderJson
// }

func Consumer() {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("nats://localhost"+port))
	if err != nil {
		fmt.Println("Произошла ошибка подключения к каналу")
		return
	}
	_, err = sc.Subscribe("Model_Json",
		func(message *stan.Msg) {
			//fmt.Println(string(message.Data))
			order, error := model.UnpackingJson(message.Data) // Вызываем функцию распаковки json в структуру и валидацию
			if error != nil {
				fmt.Println("Повторите передачу сообщения в канал")
				return
			}
			cache.SetCache(order.OrderUID, order)
			db.InsertOrder(*order)
			//cache.GetCache(order.OrderUID)

		})
	if err != nil {
		fmt.Println("Произошла ошибка подписки на канал: ", err)
		return
	}
	// stan.StartWithLastReceived())
	// err = sub.Close()
	server.Run()
	// runtime.Goexit() // ожидание добавление информации в канал
}
