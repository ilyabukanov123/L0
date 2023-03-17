package db

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/ilyabukanov123/L0/internal/model"
	_ "github.com/lib/pq" //
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "L0"
)

func InsertOrder(orderJson model.Order) {

	// Cтрока подключения
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Подключение к серверу и БД
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println("Произошла ошибка подключения к бд")
		panic(err)
	}

	defer db.Close() // закрываем подключение к БД после завершения работы функции
	b, err := json.Marshal(orderJson)
	if err != nil {
		fmt.Println("Произошла ошибка конвертации из структуры в JSON", err)
		return
	}
	result, err := db.Exec("insert into JsonOrder (Uid, Json) values ($1, $2)", orderJson.OrderUID, b)
	if err != nil {
		fmt.Println("Произошла ошибка при добавлении данных в БД")
		return
	}
	fmt.Println(result.RowsAffected()) // количество добавленных строк
	// rows, err := db.Query("select * from JsonOrder")
	// if err != nil {
	// 	fmt.Println("Произошла ошибка при чтении данных из БД")
	// 	return
	// }
	// fmt.Println(rows)
	GetOrder()
}

func GetOrder() map[string]*model.Order {

	// Cтрока подключения
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Подключение к серверу и БД
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println("Произошла ошибка подключения к бд")
		panic(err)
	}

	defer db.Close() // закрываем подключение к БД после завершения работы функции

	rows, err := db.Query("select * from JsonOrder")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	orders := make(map[string]*model.Order) // создаем map где ключом будет строка, а значением структура order

	for rows.Next() {
		var uid string                     // переменная для хранения uid из БД
		var jsonOrder string               // переменная для хранения json из БД
		err := rows.Scan(&uid, &jsonOrder) // считываем строку таблицы
		if err != nil {
			fmt.Println("Произошлка ошибка при получении данных из БД")
			continue
		}
		//fmt.Printf("%T %v\n", uid, uid)
		//fmt.Printf("%T %v\n", jsonOrder, jsonOrder)
		//fmt.Println(jsonOrder)

		var order *model.Order                              // создаем тип данных структуры order
		order, err = model.UnpackingJson([]byte(jsonOrder)) // распаковываем json в структуру
		if err != nil {
			fmt.Println("Произошла ошибка при распаковки структуры в JSON")
			panic(err)
		}
		orders[uid] = order // по ключу uid вкладываем json заказа
	}
	return orders
}
