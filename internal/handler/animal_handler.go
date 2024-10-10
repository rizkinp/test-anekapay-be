package handler

import (
	"test-anekapay-backend/internal/domain/dto"
	"test-anekapay-backend/internal/domain/entity"
	"test-anekapay-backend/internal/shared/utils"
	"test-anekapay-backend/internal/usecase"

	"github.com/labstack/echo/v4"
)

// AnimalHandler handles animal-related requests
type AnimalHandler struct {
	animalUsecase usecase.AnimalUsecase
}

// NewAnimalHandler creates a new AnimalHandler instance
func NewAnimalHandler(animalUsecase usecase.AnimalUsecase) *AnimalHandler {
	return &AnimalHandler{animalUsecase: animalUsecase}
}

// @Summary Get all animals
// @Description Returns a list of all animals in the database
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /animals [get]
// @Summary Get all animals
// @Description Returns a list of all animals in the database
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /animals [get]
func (ah *AnimalHandler) GetAllAnimals(c echo.Context) error {
	total, animals, err := ah.animalUsecase.GetAllAnimals()
	if err != nil {
		return utils.ErrorResponse(c, 500, err.Error())
	}

	if len(animals) == 0 {
		return utils.ErrorResponse(c, 404, "data not found")
	}

	animalDTOs := make([]dto.AnimalDTO, len(animals))
	for i, animal := range animals {
		animalDTOs[i] = dto.AnimalDTO{
			ID:    animal.ID,
			Name:  animal.Name,
			Class: animal.Class,
			Legs:  animal.Legs,
		}
	}

	return utils.SuccessResponse(c, 200, "success", struct {
		Total int             `json:"total"`
		Data  []dto.AnimalDTO `json:"data"`
	}{
		Total: total,
		Data:  animalDTOs,
	})
}

// @Summary Create an animal
// @Description Creates a new animal in the database
// @Accept  json
// @Produce  json
// @Param animal body dto.AnimalDTO true "Animal"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /animals [post]
func (ah *AnimalHandler) CreateAnimal(c echo.Context) error {
	animalDTO := new(dto.AnimalDTO)
	if err := c.Bind(animalDTO); err != nil {
		return utils.ErrorResponse(c, 400, "malformed request")
	}

	id, err := ah.animalUsecase.CreateAnimal(&entity.Animal{
		Name:  animalDTO.Name,
		Class: animalDTO.Class,
		Legs:  animalDTO.Legs,
	})

	if err != nil {
		return utils.ErrorResponse(c, 500, err.Error())
	}

	return utils.SuccessResponse(c, 200, "success", id)
}
