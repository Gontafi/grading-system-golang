package services

import "github.com/grading-system-golang/internal/models"

func (s *ServiceV1) AddHomeWork(work models.HomeWork) (int, error) {
	studentLesson, err := s.repository.GetStudentLesson(work.StudentID, work.LessonID)
	if err != nil || studentLesson.ID == 0 {
		return 0, err
	}

	id, err := s.repository.AddHomeWork(work)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *ServiceV1) GetHomeWorkById(id int) (models.HomeWork, error) {
	work, err := s.repository.GetHomeWorkById(id)
	if err != nil {
		return models.HomeWork{}, err
	}

	return work, nil
}

func (s *ServiceV1) GetAllHomeWorks() ([]models.HomeWork, error) {
	works, err := s.repository.GetAllHomeWorks()
	if err != nil {
		return []models.HomeWork{}, err
	}

	return works, nil
}

func (s *ServiceV1) UpdateHomeWork(work models.HomeWork) error {
	err := s.repository.UpdateHomeWork(work)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceV1) DeleteHomeWork(id int) error {
	err := s.repository.DeleteHomeWork(id)
	if err != nil {
		return err
	}

	return nil
}
