package repository

import "test-anekapay-backend/internal/domain/entity"

type AnimalRepository interface {
	GetAllAnimals() (int, []entity.Animal, error)
	CreateAnimal(animal *entity.Animal) (int64, error)
}
