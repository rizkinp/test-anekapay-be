package server

import (
	"log"
	"test-anekapay-backend/config"
	"test-anekapay-backend/internal/handler"
	"test-anekapay-backend/internal/repository"
	"test-anekapay-backend/internal/router"
	"test-anekapay-backend/internal/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Run() error {
	e := echo.New()
	// Setup database configuration
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	// Add logging middleware
	e.Use(middleware.Logger())

	// Add CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Create a group for the API routes
	api := e.Group("/api")

	//Setup Routes
	animalRepo := repository.NewAnimalRepo(db)
	animalUsecase := usecase.NewAnimalUseCase(animalRepo)
	animalHandler := handler.NewAnimalHandler(animalUsecase)
	router.NewAnimalRouter(api, animalHandler)

	// Start server
	return e.Start(":8080")
}
