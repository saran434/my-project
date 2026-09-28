package repository

import (
	"errors"
	"log"
	"strings"

	"employeejwt/internal/constants"
	"employeejwt/internal/database"
	"employeejwt/internal/models"

	"gorm.io/gorm"
)

type PostgresStore struct {
	db *gorm.DB
}

func NewPostgresStore(db *gorm.DB) *PostgresStore {
	if db == nil {
		db = database.DB
	}
	return &PostgresStore{db: db}
}

func (s *PostgresStore) Create(emp *models.Employee) error {
	if email := strings.TrimSpace(emp.Email); email != "" {
		var existing models.Employee
		err := s.db.Where("LOWER(email) = LOWER(?)", email).First(&existing).Error
		if err == nil {
			return errors.New(constants.ErrEmployeeAlreadyExists)
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("postgresStore.Create: duplicate check failed for email %s: %v", email, err)
			return err
		}
	}

	err := s.db.Create(emp).Error
	if err != nil {
		log.Printf("postgresStore.Create: failed to insert employee %+v: %v", emp, err)
		return err
	}
	return nil
}

func (s *PostgresStore) List() ([]models.Employee, error) {
	var emps []models.Employee
	if err := s.db.Find(&emps).Error; err != nil {
		log.Printf("postgresStore.List: query failed: %v", err)
		return nil, err
	}
	return emps, nil
}

func (s *PostgresStore) Get(id int) (models.Employee, error) {
	var e models.Employee
	err := s.db.First(&e, id).Error
	if err != nil {
		log.Printf("postgresStore.Get: failed for employee id %d: %v", id, err)
		return e, err
	}
	return e, nil
}

func (s *PostgresStore) Update(id int, emp *models.Employee) error {
	if email := strings.TrimSpace(emp.Email); email != "" {
		var existing models.Employee
		err := s.db.Where("LOWER(email) = LOWER(?) AND id != ?", email, id).First(&existing).Error
		if err == nil {
			return errors.New(constants.ErrEmployeeAlreadyExists)
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("postgresStore.Update: duplicate check failed for employee id %d, email %s: %v", id, email, err)
			return err
		}
	}

	err := s.db.Model(&models.Employee{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":       emp.Name,
		"email":      emp.Email,
		"department": emp.Department,
	}).Error
	if err != nil {
		log.Printf("postgresStore.Update: failed for employee id %d, payload %+v: %v", id, emp, err)
		return err
	}
	return nil
}

func (s *PostgresStore) Delete(id int) error {
	err := s.db.Delete(&models.Employee{}, id).Error
	if err != nil {
		log.Printf("postgresStore.Delete: failed for employee id %d: %v", id, err)
		return err
	}
	return nil
}
