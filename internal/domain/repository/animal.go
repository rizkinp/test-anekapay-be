package repository

import "test-anekapay-backend/internal/domain/entity"

type AnimalRepository interface {
	GetAllAnimals() (int, []entity.Animal, error)
	GetAnimal(id int) (entity.Animal, error)
	CreateAnimal(animal *entity.Animal) (int64, error)
	UpdateAnimal(id int, animal *entity.Animal) (int64, error)
	DeleteAnimal(id int) (int64, error)
}
