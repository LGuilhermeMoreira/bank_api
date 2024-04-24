package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	Port int
}

func NewConfig() *config {
	return &config{
		Port: getPort(),
	}
}

func getPort() int {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	result, err := strconv.ParseInt(port, 10, 32)

	if err != nil {
		panic(err)
	}

	return int(result)
}
