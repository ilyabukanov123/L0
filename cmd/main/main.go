package main

import (
	"github.com/ilyabukanov123/L0/cmd/consumer"
	"github.com/ilyabukanov123/L0/internal/cache"
)

func main() {
	cache.NewCache()
	consumer.Consumer()
}
