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

// GetAnimal implements repository.AnimalRepository.
func (ar *AnimalRepo) GetAnimal(id int) (entity.Animal, error) {
	var animal entity.Animal
	err := ar.db.QueryRow("SELECT id, name, class, legs FROM tbl_animals WHERE id = ? AND is_deleted = 0", id).Scan(&animal.ID, &animal.Name, &animal.Class, &animal.Legs)
	if err != nil {
		return entity.Animal{}, err
	}
	return animal, nil
}

// CreateAnimal implements repository.AnimalRepository.
func (ar *AnimalRepo) CreateAnimal(animal *entity.Animal) (int64, error) {
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

// UpdateAnimal implements repository.AnimalRepository.
func (ar *AnimalRepo) UpdateAnimal(id int, animal *entity.Animal) (int64, error) {
	// Periksa apakah hewan dengan ID tersebut ada
	var existingAnimal entity.Animal
	err := ar.db.QueryRow("SELECT id FROM tbl_animals WHERE id = ? AND is_deleted = 0", id).Scan(&existingAnimal.ID)
	if err != nil && err != sql.ErrNoRows {
		return 0, err
	}

	if existingAnimal.ID == 0 {
		return ar.CreateAnimal(animal)
	}

	sql := "UPDATE tbl_animals SET name = ?, class = ?, legs = ? WHERE id = ? AND is_deleted = 0"
	result, err := ar.db.Exec(sql, animal.Name, animal.Class, animal.Legs, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	// Return the number of affected rows
	return rowsAffected, nil
}

// DeleteAnimal implements repository.AnimalRepository.
func (ar *AnimalRepo) DeleteAnimal(id int) (int64, error) {
	// Jika ditemukan, hapus hewan
	sql := "UPDATE tbl_animals SET is_deleted = 1 WHERE id = ?"
	result, err := ar.db.Exec(sql, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
