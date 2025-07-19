package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfiguration struct {
	DbUrl          string
	MaxConnection  int
	MinIddleConn   int
	IdleTime       int
	LifeTime       int
	ServerLocation string
}

func LoadEnv() error {
	err := godotenv.Load("configs/.env")
	if err != nil {
		return err
	}

	return nil
}

func EnvToInt(conf string) (int, error) {
	data, err := strconv.Atoi(conf)

	if err != nil {
		return 0, err
	}

	return data, nil
}

func NewConfig() *AppConfiguration {
	dbUrl := os.Getenv("DB_URL")
	maxConn, _ := EnvToInt(os.Getenv("MAX_CONN"))
	minConn, _ := EnvToInt(os.Getenv("MIN_CONN"))
	lifeTime, _ := EnvToInt(os.Getenv("LIFE_TIME"))
	iddleTime, _ := EnvToInt(os.Getenv("IDDLE_TIME"))
	serverLocation := os.Getenv("SERVER_LOCATION")

	return &AppConfiguration{
		DbUrl:          dbUrl,
		MaxConnection:  maxConn,
		MinIddleConn:   minConn,
		LifeTime:       lifeTime,
		IdleTime:       iddleTime,
		ServerLocation: serverLocation,
	}
}
