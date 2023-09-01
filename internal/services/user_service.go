package services

import (
	"errors"
	"github.com/grading-system-golang/internal/models"
	"time"
)

func (s *ServiceV1) AddUser(user models.User) (int, error) {
	if user.Username == "" || user.RoleID == 0 || user.Name == "" || user.Surname == "" {
		return 0, errors.New("user data cannot be empty")
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	id, err := s.repository.AddUser(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ServiceV1) DeleteUser(id int) error {
	err := s.repository.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceV1) UpdateUser(user models.User) error {
	if user.Username == "" || user.PasswordHash == "" || user.Name == "" || user.Surname == "" {
		return errors.New("user data cannot be empty")
	}
	err := s.repository.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceV1) GetAllUsers() ([]models.User, error) {
	users, err := s.repository.AllUsers()
	if err != nil {
		return []models.User{}, err
	}

	return users, nil
}

func (s *ServiceV1) GetUserByID(id int) (models.User, error) {
	user, err := s.repository.GetUserByID(id)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *ServiceV1) GetUserByUsername(username string) (models.User, error) {
	user, err := s.repository.GetUserByUsername(username)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
