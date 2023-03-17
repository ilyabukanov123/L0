package cache

import (
	"fmt"
	"github.com/ilyabukanov123/L0/internal/db"
	"github.com/ilyabukanov123/L0/internal/model"
	"github.com/patrickmn/go-cache"
	"time"
)

var Cache *cache.Cache

// Функция по помещению данных в кэш
func SetCache(uid string, order *model.Order) {
	Cache.Set(uid, order, cache.DefaultExpiration)
}

// Функция по получению данных из кэша
func GetCache(key string) {
	value, found := Cache.Get(key)
	if found {
		fmt.Println(value)
	}
}

func RestoringTheCache() {
	orders := db.GetOrder()
	if len(orders) == 0 {
		fmt.Println("В БД нет информации о заказах")
		return
	}
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
