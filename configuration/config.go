package configuration

import (
	"os"
	"strconv"

	"mmbt/constants"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	Host     string
	HTTPPort int

	Environment     constants.Constant
	FrontendFwdPort int

	DbURL string
}

// LoadConfig gets config from .env files or from environment variables
func LoadConfig() *Config {
	_ = godotenv.Load(".env", ".env.local")

	c := Config{}

	c.Host = os.Getenv("HOST")
	c.HTTPPort, _ = strconv.Atoi(os.Getenv("HTTP_PORT"))

	c.Environment, _ = constants.Environment.Validate(os.Getenv("ENVIRONMENT"))
	c.FrontendFwdPort, _ = strconv.Atoi(os.Getenv("FRONTEND_FWD_PORT"))

	c.DbURL = os.Getenv("DB_URL")

	return &c
}
