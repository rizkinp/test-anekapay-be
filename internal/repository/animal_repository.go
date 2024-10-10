// internal/repo/animal_repo.go
package repository

import (
	"database/sql"
	"errors"
	"test-anekapay-backend/internal/domain/entity"
	"test-anekapay-backend/internal/domain/repository"
)

type AnimalRepo struct {
	db *sql.DB
}

func NewAnimalRepo(db *sql.DB) repository.AnimalRepository {
	return &AnimalRepo{db}
}

func (ar *AnimalRepo) GetAllAnimals() (int, []entity.Animal, error) {
	var total int
	err := ar.db.QueryRow("SELECT COUNT(*) FROM tbl_animals WHERE is_deleted = 0").Scan(&total)
	if err != nil {
		return 0, nil, err
	}
	sql := "SELECT id, name, class, legs FROM tbl_animals WHERE is_deleted = 0 ORDER BY created_at DESC"
	rows, err := ar.db.Query(sql)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()
	var animals []entity.Animal
	for rows.Next() {
		var animal entity.Animal
		err := rows.Scan(&animal.ID, &animal.Name, &animal.Class, &animal.Legs)
		if err != nil {
			return 0, nil, err
		}
		animals = append(animals, animal)
	}
	return total, animals, nil
}

// CreateAnimal implements repository.AnimalRepository.
func (ar *AnimalRepo) CreateAnimal(animal *entity.Animal) (int64, error) {
	// Check if an animal with the same name already exists
	var existingAnimal entity.Animal
	err := ar.db.QueryRow("SELECT id, name, class, legs FROM tbl_animals WHERE name = ? AND is_deleted = 0", animal.Name).Scan(&existingAnimal.ID, &existingAnimal.Name, &existingAnimal.Class, &existingAnimal.Legs)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}
	if existingAnimal.ID != 0 {
		return 0, errors.New("animal with the same name already exists")
	}

	sql := "INSERT INTO tbl_animals (name, class, legs) VALUES (?, ?, ?)"
	result, err := ar.db.Exec(sql, animal.Name, animal.Class, animal.Legs)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
