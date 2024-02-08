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
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Fatalf("Erro na conversão de DATABASE_PORT: %v", err.Error())
	}
	return &config{
		dbName: os.Getenv("DATABASE_NAME"),
		dbHost: os.Getenv("DATABASE_HOST"),
		dbUser: os.Getenv("DATABASE_USER"),
		dbPass: os.Getenv("DATABASE_PASS"),
		dbPort: port,
	}
}
func (c *config) GenerateDSN() {
	c.dsn = fmt.Sprintf("mongodb://%s:%s@%s/%d/%s", c.dbUser, c.dbPass, c.dbHost, c.dbPort, c.dbName)
}
