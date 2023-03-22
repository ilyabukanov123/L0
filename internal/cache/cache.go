package cache

import (
	"fmt"
	"time"

	"github.com/ilyabukanov123/L0/internal/db"
	"github.com/ilyabukanov123/L0/internal/model"
	"github.com/patrickmn/go-cache"
)

var Cache *cache.Cache

// Функция по помещению данных в кэш
func SetCache(uid string, order *model.Order) {
	Cache.Set(uid, order, cache.DefaultExpiration)
}

// Функция по получению данных из кэша
func GetCache(key string) *model.Order {
	if Cache != nil {
		value, found := Cache.Get(key) // извлекаем кэш под ключю
		if found {
			var order *model.Order = value.(*model.Order) // конвиртируем value в структуру Order
			return order
		}
	}
	return nil
}

func RestoringTheCache() {
	orders := db.GetOrder() // получаем из БД список заказов
	if len(orders) == 0 {
		fmt.Println("В БД нет информации о заказах")
		return
	}
	// перебираем все заказы и засовываем их в кэш
	for key, value := range orders {
		SetCache(key, value)
	}
}

func NewCache() {
	if Cache == nil {
		Cache = cache.New(5*time.Minute, 10*time.Minute)
		RestoringTheCache()
	}
}
