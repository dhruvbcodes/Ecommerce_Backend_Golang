package config 

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Config struct{
	PublicHost string
	Port string
	DBUser string
	DBPasswd string
	DBAddr string
	DBName string
}

// creating a singleton instance of the Config struct to not initialse everytime (good for performance)
var Envs = initConfig()

func initConfig() Config {

	godotenv.Load() // LOADS THE ENV VARIBLES IN RUNTIME

	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http.localhost"),
		Port: getEnv("PORT", "8080"),
		DBUser: getEnv("DB_USER", "root"),
		DBPasswd: getEnv("DB_PASSWD", "myPassword"),
		DBAddr: fmt.Sprintf("%s %s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName: getEnv("DB_NAME", "ecommerce"),

	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}