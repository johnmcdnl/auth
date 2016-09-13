package auth

import (
	"gopkg.in/redis.v4"
	"log"
)

var client *redis.Client

func Connection() (*redis.Client, error) {

	var err error

	if client ==nil{
		log.Println("Creating a *redis.Client")
		client, err = NewClient()
	}



	return client, err
}

func NewClient() (*redis.Client, error) {

	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})

	_, err := client.Ping().Result()

	return client, err
}
