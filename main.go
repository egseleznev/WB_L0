package main

import (
	"github.com/patrickmn/go-cache"
	"github.com/xlab/closer"
	"l0/http_server"
	"l0/service"
	"log"
)

func main() {

	c := cache.New(cache.NoExpiration, cache.NoExpiration) // create cache with no expire

	service := service.New(service.NewConfig(), c)
	server := http_server.New(http_server.NewConfig(), c) // create service and server with created cache and default config

	if err := service.Start(); err != nil {
		log.Fatal(err)
	} // starting service

	if err := service.Subscribe("channel1"); err != nil {
		log.Fatal(err)
	} // subscribe on receiving messages

	closer.Bind(func() {
		service.Unsubscribe()
		service.Close()
	}) // graceful shutdown

	go func() {
		if err := server.Start(); err != nil {
			log.Fatal(err)
		}
		closer.Hold()
	}() // starts server in another goroutine

	closer.Hold() // waiting for app closing
}
