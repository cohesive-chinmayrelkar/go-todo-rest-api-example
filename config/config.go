package config

import (
	"os"
	"strconv"

)
type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func emptyCheck(val string, name string) {
	if val == "" {
		panic(name + " is empty")
	}
}

func GetConfig() *Config {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err.Error()+ " could get DB_PORT")
	}
	host := os.Getenv("DB_HOST")
	Username := os.Getenv("DB_USER")
	Password := os.Getenv("DB_PASS")
	Name := os.Getenv("DB_NAME")

	emptyCheck(host, "DB_HOST")
	emptyCheck(Username, "DB_USER")
	emptyCheck(Password, "DB_PASS")
	emptyCheck(Name, "DB_NAME")

	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     host,
			Port:     port,
			Username: Username,
			Password: Password,
			Name:     Name,
			Charset:  "utf8",
		},
	}
}
