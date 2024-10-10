package usecase

import (
	"test-anekapay-backend/internal/domain/entity"
	"test-anekapay-backend/internal/domain/repository"
)

type AnimalUsecase interface {
	GetAllAnimals() (int, []entity.Animal, error)
	CreateAnimal(animal *entity.Animal) (int64, error)
}

type animalUsecase struct {
	animalRepo repository.AnimalRepository
}

// CreateAnimal implements AnimalUsecase.
func (a *animalUsecase) CreateAnimal(animal *entity.Animal) (int64, error) {
	return a.animalRepo.CreateAnimal(animal)
}

// GetAllAnimals implements AnimalUsecase.
func (a *animalUsecase) GetAllAnimals() (int, []entity.Animal, error) {
	return a.animalRepo.GetAllAnimals()
}

func NewAnimalUseCase(ar repository.AnimalRepository) AnimalUsecase {
	return &animalUsecase{ar}
}
