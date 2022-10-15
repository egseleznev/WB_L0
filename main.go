package main

import (
	"github.com/patrickmn/go-cache"
	"l0/service"
	"log"
)

func main() {

	c := cache.New(cache.NoExpiration, cache.NoExpiration)

	config := service.NewConfig()

	s := service.New(config, c)
	defer s.Close()

	err := s.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = s.Subscribe("channel1")
	if err != nil {
		log.Fatal(err)
	}
	for {
	}

	s.Unsubscribe()

}
