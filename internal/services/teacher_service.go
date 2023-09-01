package services

import (
	"errors"
	"github.com/grading-system-golang/internal/models"
)

func (s *ServiceV1) AddTeacher(teacher models.Teacher) (int, error) {
	if teacher.Name == "" || teacher.Surname == "" {
		return 0, errors.New("teacher name and surname cannot be empty")
	}

	id, err := s.repository.AddTeacher(teacher)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ServiceV1) DeleteTeacher(id int) error {
	err := s.repository.DeleteTeacher(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceV1) UpdateTeacher(teacher models.Teacher) error {
	if teacher.Name == "" || teacher.Surname == "" {
		return errors.New("teacher name and surname cannot be empty")
	}

	err := s.repository.UpdateTeacher(teacher)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceV1) GetAllTeachers() ([]models.Teacher, error) {
	teachers, err := s.repository.AllTeachers()
	if err != nil {
		return []models.Teacher{}, err
	}

	return teachers, nil
}

func (s *ServiceV1) GetTeacherByID(id int) (models.Teacher, error) {
	teacher, err := s.repository.GetTeacherByID(id)
	if err != nil {
		return models.Teacher{}, err
	}

	return teacher, nil
}
