package services

import (
	"errors"
	"github.com/grading-system-golang/internal/models"
)

func (s *ServiceV1) AddStudent(student models.Student) (int, error) {
	if student.Name == "" || student.Surname == "" {
		return 0, errors.New("student name and surname cannot be empty")
	}

	id, err := s.repository.AddStudent(student)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ServiceV1) DeleteStudent(id int) error {
	err := s.repository.DeleteStudent(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceV1) UpdateStudent(student models.Student) error {
	if student.Name == "" || student.Surname == "" {
		return errors.New("student name and surname cannot be empty")
	}

	err := s.repository.UpdateStudent(student)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceV1) GetAllStudents() ([]models.Student, error) {
	students, err := s.repository.AllStudents()
	if err != nil {
		return []models.Student{}, err
	}

	return students, nil
}

func (s *ServiceV1) GetStudentByID(id int) (models.Student, error) {
	student, err := s.repository.GetStudentByID(id)
	if err != nil {
		return models.Student{}, err
	}

	return student, nil
}
