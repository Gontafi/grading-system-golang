package repositories

import (
	"github.com/grading-system-golang/internal/models"
)

func (r *RepositoryV1) AddStudent(student models.Student) (int, error) {

	_, err := r.db.Exec(
		r.ctx,
		`INSERT INTO students(user_id, name, surname) VALUES ($1, $2, $3)`,
		student.UserID, student.Name, student.Surname)
	if err != nil {
		return 0, err
	}

	return student.UserID, nil
}

func (r *RepositoryV1) DeleteStudent(id int) error {
	_, err := r.db.Exec(r.ctx, `DELETE FROM students WHERE user_id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryV1) UpdateStudent(student models.Student) error {
	_, err := r.db.Exec(
		r.ctx,
		`UPDATE students set name = $2, surname = $3 WHERE user_id = $1`,
		student.UserID, student.Name, student.Surname,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryV1) AllStudents() ([]models.Student, error) {
	rows, err := r.db.Query(r.ctx, `SELECT user_id, name, surname FROM students`)
	if err != nil {
		return []models.Student{}, err
	}
	defer rows.Close()

	var students []models.Student

	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.UserID, &student.Name, &student.Surname)
		if err != nil {
			return []models.Student{}, err
		}

		students = append(students, student)
	}

	return students, nil
}

func (r *RepositoryV1) GetStudentByID(id int) (models.Student, error) {
	var student models.Student
	err := r.db.QueryRow(r.ctx, `SELECT user_id, name, surname FROM students WHERE user_id = $1`, id).Scan(
		&student.UserID, &student.Name, &student.Surname,
	)
	if err != nil {
		return models.Student{}, err
	}

	return student, nil
}
