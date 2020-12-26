package configuration

import (
	"os"
	"strconv"

	"mmbt/constants"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	Host         string
	HTTPPort     int
	HTTPOnly     bool
	HTTPSPort    int
	CertFilePath string
	KeyFilePath  string

	Environment     constants.Constant
	DevFrontendPort int
}

// LoadConfig gets config from .env files or from environment variables
func LoadConfig() *Config {
	_ = godotenv.Load(".env", ".env.local")

	c := Config{}

	c.Host = os.Getenv("HOST")

	c.HTTPPort, _ = strconv.Atoi(os.Getenv("HTTP_PORT"))

	if os.Getenv("HTTP_ONLY") == "true" {
		c.HTTPOnly = true
	} else {
		c.HTTPOnly = false
	}

	c.HTTPSPort, _ = strconv.Atoi(os.Getenv("HTTPS_PORT"))

	c.CertFilePath = os.Getenv("CERT_FILE_PATH")
	c.KeyFilePath = os.Getenv("KEY_FILE_PATH")

	c.Environment, _ = constants.Environment.Validate(os.Getenv("ENVIRONMENT"))
	c.DevFrontendPort, _ = strconv.Atoi(os.Getenv("DEV_FRONTEND_PORT"))

	return &c
}
