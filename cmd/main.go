package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/srselivan/user-balance-microservice/internal/database"
	"github.com/srselivan/user-balance-microservice/internal/server"
)

func main() {
	config := server.NewConfig()
	if _, err := toml.DecodeFile("configs/microservice.toml", config); err != nil {
		log.Print(err)
	}

	bd := database.New(config)
	if err := bd.ConnectToDB(); err != nil {
		log.Fatal(err)
	}
	defer bd.CloseConnection()

	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
