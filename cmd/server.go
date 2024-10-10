package server

import (
	"log"
	"test-anekapay-backend/config"

	"github.com/labstack/echo/v4"
)

func Run() error {
	e := echo.New()
	// Setup database configuration
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	// Start server
	return e.Start(":8000")
}
