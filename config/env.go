package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Host                   string
	Port                   int
	DBUser                 string
	DBPassword             string
	DBName                 string
	SSLMode                string
	JWTExpirationInSeconds int
	JWT_SECRET             string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		Host:                   getEnv("HOST", "localhost"),
		Port:                   getEnvInt("PORT", 8080),
		DBUser:                 getEnv("DBUSER", "root"),
		DBPassword:             getEnv("DBPASSWORD", "root"),
		DBName:                 "ecom-1",
		SSLMode:                "disable",
		JWTExpirationInSeconds: getEnvInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
		JWT_SECRET:             getEnv("JWT_SECRET", "aushdnkjzxncjalrhwlajsnfjdnlawj"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		num, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return num
	}
	return fallback
}
