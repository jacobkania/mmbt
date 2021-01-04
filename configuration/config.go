package configuration

import (
	"os"
	"strconv"

	"mmbt/constants"

	"github.com/joho/godotenv"
)

// Config is the application env configuration
var Config config

// Config holds the application configuration
type config struct {
	Host     string
	HTTPPort int

	Environment     constants.Constant
	FrontendFwdPort int

	BCryptCost int

	DbURL string
}

// LoadConfig gets config from .env files or from environment variables
func LoadConfig() {
	_ = godotenv.Load(".env", ".env.local")

	c := config{}

	c.Host = os.Getenv("HOST")
	c.HTTPPort, _ = strconv.Atoi(os.Getenv("HTTP_PORT"))

	c.Environment, _ = constants.Environment.Validate(os.Getenv("ENVIRONMENT"))
	c.FrontendFwdPort, _ = strconv.Atoi(os.Getenv("FRONTEND_FWD_PORT"))

	c.BCryptCost, _ = strconv.Atoi(os.Getenv("BCRYPT_COST"))

	c.DbURL = os.Getenv("DB_URL")

	Config = c
}
