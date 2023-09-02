package services

import (
	"encoding/json"
	"github.com/grading-system-golang/internal/models"
	"github.com/redis/go-redis/v9"
)

func (s *ServiceV1) GetUserRole(userID int) (models.Role, error) {
	cacheKey := s.getCacheKey("userRole", userID)

	data, err := s.rdb.Get(s.ctx, cacheKey).Bytes()
	if err == nil {
		var role models.Role
		if err := json.Unmarshal(data, &role); err != nil {
			return models.Role{}, err
		}
		return role, nil
	} else if err != redis.Nil {
		return models.Role{}, err
	}

	role, err := s.repository.GetUserRole(userID)
	if err != nil {
		return models.Role{}, err
	}

	roleJSON, err := json.Marshal(role)
	if err != nil {
		return models.Role{}, err
	}
	if err := s.rdb.Set(s.ctx, cacheKey, roleJSON, s.expiry).Err(); err != nil {
		return models.Role{}, err
	}

	return role, nil
}

func (s *ServiceV1) AddRoleUser(roleID int, userID int) (int, error) {
	id, err := s.repository.AddRoleUser(roleID, userID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ServiceV1) RemoveRoleUser(roleID int, userID int) error {
	err := s.repository.RemoveRoleUser(roleID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceV1) GetUsersForRole(roleID int) ([]models.User, error) {
	users, err := s.repository.GetUsersForRole(roleID)
	if err != nil {
		return []models.User{}, err
	}

	return users, nil
}

func (s *ServiceV1) GetRolesForUser(userID int) ([]models.Role, error) {
	roles, err := s.repository.GetRolesForUser(userID)
	if err != nil {
		return []models.Role{}, err
	}

	return roles, nil
}
