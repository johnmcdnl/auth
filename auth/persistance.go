package auth

import (
	"fmt"
	"gopkg.in/redis.v4"
	"os"
	"sync"
	"log"
)

var client *redis.Client
var once sync.Once


func Connection() *redis.Client {
	once.Do(func() {
		log.Println("Creating a *redis.Client")
		client = NewClient()
	})

	return client
}

func NewClient() *redis.Client{

	client = redis.NewClient(&redis.Options{
		Addr:     "http:/172.17.0.0:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})

	_, err := client.Ping().Result()

	if err != nil {
		fmt.Println("Couldn't find a REDIS server")
		os.Exit(1)
	}

	return client
}
