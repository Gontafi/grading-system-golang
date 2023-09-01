package services

import (
	"github.com/grading-system-golang/internal/models"
)

func (s *ServiceV1) AddStudentToLesson(studentID int, lessonID int) (int, error) {
	id, err := s.repository.AddStudentToLesson(studentID, lessonID)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ServiceV1) RemoveStudentFromLesson(studentID int, lessonID int) error {
	err := s.repository.RemoveStudentFromLesson(studentID, lessonID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceV1) GetStudentsForLesson(lessonID int) ([]models.Student, error) {
	students, err := s.repository.GetStudentsForLesson(lessonID)
	if err != nil {
		return []models.Student{}, err
	}

	return students, nil
}

func (s *ServiceV1) GetLessonsForStudent(studentID int) ([]models.Lesson, error) {
	lessons, err := s.repository.GetLessonsForStudent(studentID)
	if err != nil {
		return []models.Lesson{}, err
	}

	return lessons, nil
}

func (s *ServiceV1) GetStudentLesson(studentID int, lessonID int) (models.StudentLesson, error) {
	studentLesson, err := s.repository.GetStudentLesson(studentID, lessonID)
	if err != nil {
		return models.StudentLesson{}, err
	}

	return studentLesson, nil
}
