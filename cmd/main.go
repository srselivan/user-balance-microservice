package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/srselivan/user-balance-microservice/internal/database"
	"github.com/srselivan/user-balance-microservice/internal/server"
)

func main() {
	srvConfig := server.NewConfig()
	if _, err := toml.DecodeFile("configs/microservice.toml", srvConfig); err != nil {
		log.Print(err)
	}

	dbConfig := database.NewConfig()
	if _, err := toml.DecodeFile("configs/microservice.toml", dbConfig); err != nil {
		log.Print(err)
	}

	bd := database.New(dbConfig)
	if err := bd.ConnectToDB(); err != nil {
		log.Fatal(err)
	}
	defer bd.CloseConnection()

	s := server.New(srvConfig)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
