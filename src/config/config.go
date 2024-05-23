package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	Port         int
	DatabaseUri  string
	DatabaseName string
	JwtTime      int
	JwtPassword  string
}

func NewConfig() *config {
	return &config{
		Port:         getPort(),
		DatabaseUri:  getDatabaseURI(),
		DatabaseName: getDatabaseName(),
		JwtTime:      getJwtTime(),
		JwtPassword:  getJWTPassword(),
	}
}

func getPort() int {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	result, err := strconv.ParseInt(port, 10, 32)

	if err != nil {
		panic(err)
	}

	return int(result)
}

func getDatabaseURI() string {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}

	return os.Getenv("DATABASE_URI")
}

func getDatabaseName() string {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}

	return os.Getenv("DATABASE_NAME")
}

func getJwtTime() int {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}

	time := os.Getenv("JWT_TIME")

	result, err := strconv.ParseInt(time, 10, 16)

	if err != nil {
		panic(err)
	}

	return int(result)
}

func getJWTPassword() string {
	if err := godotenv.Load("../../.env"); err != nil {
		panic(err)
	}

	return os.Getenv("JWT_PASSWORD")
}
