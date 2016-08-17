package main

import (
	"gopkg.in/redis.v4"
	"os"
	"fmt"
)

func init() {
	NewClient()
}

var client *redis.Client

func Connection() *redis.Client {
	return client
}

func NewClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0, // use default DB
	})

	_, err := client.Ping().Result()

	if err != nil {
		fmt.Println("Couldn't find a REDIS server")
		os.Exit(1)
	}

}