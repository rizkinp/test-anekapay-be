package usecase

import (
	"test-anekapay-backend/internal/domain/entity"
	"test-anekapay-backend/internal/domain/repository"
)

type AnimalUsecase interface {
	GetAllAnimals() (int, []entity.Animal, error)
	GetAnimal(id int) (entity.Animal, error)
	CreateAnimal(animal *entity.Animal) (int64, error)
	UpdateAnimal(id int, animal *entity.Animal) (int64, error)
	DeleteAnimal(id int) (int64, error)
}

type animalUsecase struct {
	animalRepo repository.AnimalRepository
}

// CreateAnimal implements AnimalUsecase.
func (a *animalUsecase) CreateAnimal(animal *entity.Animal) (int64, error) {
	return a.animalRepo.CreateAnimal(animal)
}

// GetAnimal implements AnimalUsecase.
func (a *animalUsecase) GetAnimal(id int) (entity.Animal, error) {
	return a.animalRepo.GetAnimal(id)
}

// GetAllAnimals implements AnimalUsecase.
func (a *animalUsecase) GetAllAnimals() (int, []entity.Animal, error) {
	return a.animalRepo.GetAllAnimals()
}

// UpdateAnimal implements AnimalUsecase.
func (a *animalUsecase) UpdateAnimal(id int, animal *entity.Animal) (int64, error) {
	return a.animalRepo.UpdateAnimal(id, animal)
}

// DeleteAnimal implements AnimalUsecase.
func (a *animalUsecase) DeleteAnimal(id int) (int64, error) {
	return a.animalRepo.DeleteAnimal(id)
}
func NewAnimalUseCase(ar repository.AnimalRepository) AnimalUsecase {
	return &animalUsecase{ar}
}
