package testhelpers

import (
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

// CreateTestDB creates a database to use for unit tests
func CreateTestDB() *pg.DB {
	_ = godotenv.Load(".env", ".env.local")

	testDBURL := os.Getenv("TEST_DB_URL")

	opt, err := pg.ParseURL(testDBURL)
	if err != nil {
		log.Fatalf("Failed to establish connection to test database, address: %v", testDBURL)
	}

	return pg.Connect(opt)
}
