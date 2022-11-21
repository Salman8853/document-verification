package config

import (
	"os"
)

type config struct {
	PORT      string
	MONGO_URI string
	DB_NAME   string
}

var APP_CONFIG config

func init() {
	port := os.Getenv("PORT")

	mongo_uri := os.Getenv("MONGO_URI")
	db_name := os.Getenv("DB_NAME")

	APP_CONFIG = config{PORT: port, MONGO_URI: mongo_uri, DB_NAME: db_name}

}
