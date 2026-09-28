package repository

import (
	"errors"
	"log"

	"employeejwt/internal/constants"
	"employeejwt/internal/models"

	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{db: db}
}

func (s *UserStore) CreateUser(user *models.User) error {
	var existing models.User
	err := s.db.Where("LOWER(email) = LOWER(?)", user.Email).First(&existing).Error
	if err == nil {
		return errors.New(constants.ErrUserAlreadyExists)
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("userStore.CreateUser: duplicate check failed for email %s: %v", user.Email, err)
		return err
	}

	err = s.db.Create(user).Error
	if err != nil {
		log.Printf("userStore.CreateUser: failed for user %+v: %v", user, err)
		return err
	}
	return nil
}

func (s *UserStore) GetByEmail(email string) (models.User, error) {
	var user models.User
	err := s.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("userStore.GetByEmail: user not found for email %s", email)
			return models.User{}, errors.New("user not found")
		}
		log.Printf("userStore.GetByEmail: query failed for email %s: %v", email, err)
		return models.User{}, err
	}
	return user, nil
}

func (s *UserStore) GetByID(id int) (models.User, error) {
	var user models.User
	err := s.db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("userStore.GetByID: user not found for id %d", id)
			return models.User{}, errors.New("user not found")
		}
		log.Printf("userStore.GetByID: query failed for id %d: %v", id, err)
		return models.User{}, err
	}
	return user, nil
}

func (s *UserStore) EnsureSchema() error {
	err := s.db.AutoMigrate(&models.User{})
	if err != nil {
		log.Printf("userStore.EnsureSchema: create table failed: %v", err)
		return err
	}
	return nil
}
