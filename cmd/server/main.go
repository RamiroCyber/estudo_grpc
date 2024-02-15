package main

import (
	"context"
	"github.com/ramiroribeiro/estudo_grpc/config"
	"github.com/ramiroribeiro/estudo_grpc/database"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	config.LoadEnvironment()
	database.ConnectDB()
	database.SetupCollections()
}

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	database.Client.Disconnect(ctx)

	log.Println("SERVER DOWN")

}
