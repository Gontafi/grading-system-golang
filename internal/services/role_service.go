package services

import (
	"errors"
	"github.com/grading-system-golang/internal/models"
)

func (s *ServiceV1) AddRole(role models.Role) (int, error) {
	if role.RoleName == "" {
		return 0, errors.New("role name cannot be empty")
	}

	id, err := s.repository.AddRole(role)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ServiceV1) DeleteRole(id int) error {
	err := s.repository.DeleteRole(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceV1) UpdateRole(role models.Role) error {
	if role.RoleName == "" {
		return errors.New("role name cannot be empty")
	}

	err := s.repository.UpdateRole(role)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceV1) GetAllRoles() ([]models.Role, error) {
	Roles, err := s.repository.AllRoles()
	if err != nil {
		return []models.Role{}, err
	}

	return Roles, nil
}

func (s *ServiceV1) GetRoleByID(id int) (models.Role, error) {
	role, err := s.repository.GetRoleByID(id)
	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}
