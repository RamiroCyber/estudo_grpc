package main

import (
	"flag"
	"github.com/joho/godotenv"
	"github.com/ramiroribeiro/estudo_grpc/database"
	"log"
)

var local bool

func init() {
	flag.BoolVar(&local, "local", true, "run service local")
	flag.Parse()
}

func main() {
	if local {
		if err := godotenv.Load(".env"); err != nil {
			log.Panic(err)
		}
	}

	cfg := database.NewConfig()
	cfg.GenerateDSN()

	conn, err := database.NewConnection(cfg)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

}
