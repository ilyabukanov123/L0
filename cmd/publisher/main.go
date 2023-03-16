package main

import (
	"fmt"
	"os"

	"github.com/nats-io/stan.go"
)

const (
	port      = ":4223"        // Порт nast-streaming на котором работает клиент
	clusterID = "test-cluster" // Настройка по умолчанию
	clientID  = "publisher"    // ID клиента
)

func main() {
	fileName := os.Args[1]                 // go run publicher.go main.json. 0 аргумент go run publicher.go; 1 аргумент main.json
	contents, err := os.ReadFile(fileName) // Считываем данных в json
	// Проверяем возникла ли ошибка при чтении json файла
	if err != nil {
		fmt.Println("Невозможно прочитать файл. Возникла ошибка: ", err)
		return
	}
	// fmt.Printf("%T", contents)
	fmt.Println("Содержимое файла:\n", string(contents)) // Смотрим содержимое файла путем конвертации последовательности байтов в string
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL("nats://localhost"+port))
	// Проверяем наличие ошибок при подключении к nats-streaming'у через интерфей
	if err != nil {
		panic(err)
	}
	err = sc.Publish("Model_Json", []byte(contents))
	if err != nil {
		panic(err)
	}
	// if err != nil {
	// 	fmt.Println("Возникла ошибка при подключении к nats-streaming", err)
	// 	return
	// } else {
	// 	sc.Publish("Model_Json", []byte(contents)) // Отправляем массив байтов в канал nats-streaming
	// }
}
