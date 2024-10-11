package router

import (
	"test-anekapay-backend/internal/handler"

	"github.com/labstack/echo/v4"
)

func NewAnimalRouter(g *echo.Group, ah *handler.AnimalHandler) {
	g.GET("/animals", ah.GetAllAnimals)
	g.GET("/animals/:id", ah.GetAnimal)
	g.POST("/animals", ah.CreateAnimal)
	g.PUT("/animals/:id", ah.UpdateAnimal)
	g.DELETE("/animals/:id", ah.DeleteAnimal)
}
