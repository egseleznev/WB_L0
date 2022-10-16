package main

import (
	"github.com/patrickmn/go-cache"
	"l0/http_server"
	"l0/service"
	"log"
)

func main() {

	c := cache.New(cache.NoExpiration, cache.NoExpiration)

	serviceConfig := service.NewConfig()
	serverConfig := http_server.NewConfig()

	service := service.New(serviceConfig, c)
	server := http_server.New(serverConfig, c)

	err := service.Start()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		err = server.Start()
		if err != nil {
			log.Fatal(err)
		}
		for {
		}
	}()

	go func() {
		defer service.Close()
		err = service.Subscribe("channel1")
		if err != nil {
			log.Fatal(err)
		}
		for {
		}
		service.Unsubscribe()
	}()

	for {
	}
}
