package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/srselivan/user-balance-microservice/internal/server"
)

func main() {
	config := server.NewConfig()
	if _, err := toml.DecodeFile("configs/server.toml", config); err != nil {
		log.Print(err)
	}

	s := server.New(config)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
