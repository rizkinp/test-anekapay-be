package handler

import (
	"strconv"
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

// @Summary Get an animal
// @Description Returns an animal by ID
// @Success 200 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /animals/{id} [get]
func (ah *AnimalHandler) GetAnimal(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.ErrorResponse(c, 400, "malformed request")
	}
	animal, err := ah.animalUsecase.GetAnimal(id)

	if animal.ID == 0 || err != nil {
		return utils.ErrorResponse(c, 404, err.Error())
	}

	animalDTO := dto.AnimalDTO{
		ID:    animal.ID,
		Name:  animal.Name,
		Class: animal.Class,
		Legs:  animal.Legs,
	}

	return utils.SuccessResponse(c, 200, "success", animalDTO)
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

	// Buat hewan baru
	id, err := ah.animalUsecase.CreateAnimal(&entity.Animal{
		Name:  animalDTO.Name,
		Class: animalDTO.Class,
		Legs:  animalDTO.Legs,
	})

	if err != nil {
		return utils.ErrorResponse(c, 500, err.Error())
	}

	// Kembalikan data hewan yang baru dibuat
	createdAnimal := dto.AnimalDTO{
		ID:    int(id), // id sudah bertipe int64
		Name:  animalDTO.Name,
		Class: animalDTO.Class,
		Legs:  animalDTO.Legs,
	}

	return utils.SuccessResponse(c, 200, "Created successfully", createdAnimal)
}

// @Summary Update an animal
// @Description Updates an animal in the database
// @Accept  json
// @Produce  json
// @Param id path int true "Animal ID"
// @Param animal body dto.AnimalDTO true "Animal"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /animals/{id} [put]
func (ah *AnimalHandler) UpdateAnimal(c echo.Context) error {
	// Ambil ID dari parameter URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.ErrorResponse(c, 400, "Invalid animal ID")
	}

	// Cek apakah ID ada di database
	existingAnimal, err := ah.animalUsecase.GetAnimal(id)
	if err != nil || existingAnimal.ID == 0 {
		return utils.ErrorResponse(c, 404, "Data not found")
	}

	// Bind data dari request body ke AnimalDTO
	animalDTO := new(dto.AnimalDTO)
	if err := c.Bind(animalDTO); err != nil {
		return utils.ErrorResponse(c, 400, "Malformed request")
	}

	// Panggil usecase untuk update data
	rowsAffected, err := ah.animalUsecase.UpdateAnimal(id, &entity.Animal{
		Name:  animalDTO.Name,
		Class: animalDTO.Class,
		Legs:  animalDTO.Legs,
	})

	// Cek jika tidak ada baris yang terpengaruh
	if rowsAffected == 0 || err != nil {
		return utils.ErrorResponse(c, 404, "Update failed or no changes made")
	}

	// Kembalikan data hewan yang di-update
	updatedAnimal := dto.AnimalDTO{
		ID:    id,
		Name:  animalDTO.Name,
		Class: animalDTO.Class,
		Legs:  animalDTO.Legs,
	}

	return utils.SuccessResponse(c, 200, " Data updated successfully", updatedAnimal)
}

// @Summary Delete an animal
// @Description Deletes an animal from the database
// @Accept  json
// @Produce  json
// @Param id path int true "Animal ID"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 404 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Router /animals/{id} [delete]
func (ah *AnimalHandler) DeleteAnimal(c echo.Context) error {
	// Ambil ID dari parameter URL
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return utils.ErrorResponse(c, 400, "Invalid animal ID")
	}

	// Cek apakah ID ada di database
	existingAnimal, err := ah.animalUsecase.GetAnimal(id)
	if err != nil || existingAnimal.ID == 0 {
		return utils.ErrorResponse(c, 404, "Data not found")
	}

	// Panggil usecase untuk delete data
	rowsAffected, err := ah.animalUsecase.DeleteAnimal(id)

	// Cek jika tidak ada baris yang terpengaruh
	if rowsAffected == 0 || err != nil {
		return utils.ErrorResponse(c, 404, "Delete failed or no changes made")
	}

	return utils.SuccessResponse(c, 200, "Data deleted successfully", nil)
}
