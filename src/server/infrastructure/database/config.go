package database

import (
	"os"
)

type config struct {
	host     string
	database string
	port     string
	user     string
	password string
}

func newConfigMysql() *config {
	return &config{
		host:     os.Getenv("MYSQL_HOST"),
		database: os.Getenv("MYSQL_DATABASE"),
		port:     os.Getenv("MYSQL_PORT"),
		user:     os.Getenv("MYSQL_USER"),
		password: os.Getenv("MYSQL_PASSWORD"),
	}
}
