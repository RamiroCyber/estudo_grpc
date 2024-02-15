package config

import (
	"fmt"
	"github.com/joho/godotenv"
	util "github.com/ramiroribeiro/estudo_grpc/utils"
)

func LoadEnvironment() {
	err := godotenv.Load(".env")
	if err != nil {
		util.Logger("ERROR", fmt.Sprintf(".env: %v", err))
	}
}
