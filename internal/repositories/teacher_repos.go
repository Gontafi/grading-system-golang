package repositories

import (
	"github.com/grading-system-golang/internal/models"
)

func (r *RepositoryV1) AddTeacher(teacher models.Teacher) (int, error) {

	_, err := r.db.Exec(
		r.ctx,
		`INSERT INTO teachers(user_id, name, surname) VALUES ($1, $2, $3)`,
		teacher.UserID, teacher.Name, teacher.Surname)
	if err != nil {
		return 0, err
	}

	return teacher.UserID, nil
}

func (r *RepositoryV1) DeleteTeacher(id int) error {
	_, err := r.db.Exec(r.ctx, `DELETE FROM teachers WHERE user_id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryV1) UpdateTeacher(teacher models.Teacher) error {
	_, err := r.db.Exec(
		r.ctx,
		`UPDATE teachers set name = $2, surname = $3 WHERE user_id = $1`,
		teacher.UserID, teacher.Name, teacher.Surname,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryV1) AllTeachers() ([]models.Teacher, error) {
	rows, err := r.db.Query(r.ctx, `SELECT user_id, name, surname FROM teachers`)
	if err != nil {
		return []models.Teacher{}, err
	}
	defer rows.Close()

	var teachers []models.Teacher

	for rows.Next() {
		var teacher models.Teacher
		err := rows.Scan(&teacher.UserID, &teacher.Name, &teacher.Surname)
		if err != nil {
			return []models.Teacher{}, err
		}

		teachers = append(teachers, teacher)
	}

	return teachers, nil
}

func (r *RepositoryV1) GetTeacherByID(id int) (models.Teacher, error) {
	var teacher models.Teacher
	err := r.db.QueryRow(r.ctx, `SELECT user_id, name, surname FROM teachers WHERE id = $1`, id).Scan(
		&teacher.UserID, &teacher.Name, &teacher.Surname,
	)
	if err != nil {
		return models.Teacher{}, err
	}

	return teacher, nil
}
