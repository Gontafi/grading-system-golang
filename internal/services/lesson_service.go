package services

import (
	"errors"
	"github.com/grading-system-golang/internal/models"
)

func (s *ServiceV1) AddLesson(lesson models.Lesson) (int, error) {
	if lesson.Name == "" {
		return 0, errors.New("lesson name cannot be empty")
	}

	id, err := s.repository.AddLesson(lesson)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ServiceV1) DeleteLesson(id int) error {
	err := s.repository.DeleteLesson(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceV1) UpdateLesson(lesson models.Lesson) error {
	if lesson.Name == "" {
		return errors.New("lesson name cannot be empty")
	}

	err := s.repository.UpdateLesson(lesson)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceV1) GetAllLessons() ([]models.Lesson, error) {
	lessons, err := s.repository.AllLessons()
	if err != nil {
		return []models.Lesson{}, err
	}

	return lessons, nil
}

func (s *ServiceV1) GetLessonByID(id int) (models.Lesson, error) {
	lesson, err := s.repository.GetLessonByID(id)
	if err != nil {
		return models.Lesson{}, err
	}

	return lesson, nil
}
