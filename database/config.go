package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config interface {
	Dsn() string
	DbName() string
}

type config struct {
	dbHost string
	dbPort int
	dbName string
	dsn    string
	dbUser string
	dbPass string
}

func (c *config) DbName() string {
	return c.dbName
}

func (c *config) Dsn() string {
	return c.dsn
}

func NewConfig() *config {
	var cfg config
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Fatalf("Erro na conversão de DATABASE_PORT: %v", err.Error())
	}
	cfg.dbName = os.Getenv("DATABASE_NAME")
	cfg.dbHost = os.Getenv("DATABASE_HOST")
	cfg.dbUser = os.Getenv("DATABASE_USER")
	cfg.dbPass = os.Getenv("DATABASE_PASS")
	cfg.dbPort = port
	cfg.dsn = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", cfg.dbUser, cfg.dbPass, cfg.dbHost, cfg.dbPort, cfg.dbName)

	return &cfg
}
func (c *config) Dns() string {
	return c.dsn
}
