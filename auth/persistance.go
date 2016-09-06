package auth

import (
	"fmt"
	"gopkg.in/redis.v4"
	"os"
	"sync"
)

var client *redis.Client
var once sync.Once


func Connection() *redis.Client {
	once.Do(func() {
		NewClient()
	})
	return client
}

func NewClient() {

	client = redis.NewClient(&redis.Options{
		Addr:     "http://192.168.0.9:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})

	_, err := client.Ping().Result()

	if err != nil {
		fmt.Println("Couldn't find a REDIS server")
		os.Exit(1)
	}

}
