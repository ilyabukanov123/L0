package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ilyabukanov123/L0/internal/cache"
	"github.com/julienschmidt/httprouter"
)

/*
http.ResponseWriter - это ответ от сервера; w - это файл, в который заносится иформация от сервера
http.Request - это указатель на запрос от сервера
httprouter.Params - позволяет прочитать get, post параметры при запросе
*/

type OrderModel struct {
	Uid  string
	Json string
}

var tpl *template.Template

func Order(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	//fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
	file, error := filepath.Abs("../../internal/server/Order.html")
	if error != nil {
		fmt.Println("Произошла ошибка пути к файлу")
		panic(error)
	}
	tpl = template.Must(template.ParseFiles(file))
	id := param.ByName("id")
	data, error := json.Marshal(cache.GetCache(id))
	if error != nil {
		fmt.Println("Произошла ошибка при конвертации структуры в JSON")
	}
	structOrder := OrderModel{id, string(data)}
	tpl.Execute(w, structOrder)
}

func Run() {
	router := httprouter.New()
	router.GET("/order/:id", Order)
	log.Fatal(http.ListenAndServe(":8070", router))
}
