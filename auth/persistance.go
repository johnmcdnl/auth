package auth

import (
	"fmt"
	"gopkg.in/redis.v4"
	"os"
	"time"
)

func init() {
	NewClient()
}

var client *redis.Client

func Connection() *redis.Client {
	return client
}

func NewClient() {

	time.Sleep(5*time.Second)
	client = redis.NewClient(&redis.Options{
		Addr:     "http://192.168.0.9:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()

	if err != nil {
		fmt.Println("Couldn't find a REDIS server")
		os.Exit(1)
	}

}
