package services

import (
	"errors"
	"github.com/grading-system-golang/internal/models"
)

func (s *ServiceV1) CreateMark(mark models.Mark) (int, error) {
	if mark.HomeWorkGrade < 1 || mark.HomeWorkGrade > 5 || (mark.AttendanceGrade != 0 && mark.AttendanceGrade != 1) {
		return 0, errors.New("invalid mark data")
	}

	markID, err := s.repository.CreateMark(mark)
	if err != nil {
		return 0, err
	}

	return markID, nil
}

func (s *ServiceV1) GetMarkByID(markID int) (models.Mark, error) {
	mark, err := s.repository.GetMarkByID(markID)
	if err != nil {
		return models.Mark{}, err
	}

	return mark, nil
}

func (s *ServiceV1) GetAllMarks() ([]models.Mark, error) {
	marks, err := s.repository.GetAllMarks()
	if err != nil {
		return []models.Mark{}, err
	}

	return marks, err
}

func (s *ServiceV1) DeleteMark(markID int) error {
	err := s.repository.DeleteMark(markID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ServiceV1) UpdateMark(mark models.Mark) error {
	err := s.repository.UpdateMark(mark)
	if err != nil {
		return err
	}

	return nil
}
