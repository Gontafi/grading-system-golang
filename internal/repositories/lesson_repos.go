package repositories

import (
	"github.com/grading-system-golang/internal/models"
)

func (r *RepositoryV1) AddLesson(lesson models.Lesson) (int, error) {
	var id int

	err := r.db.QueryRow(
		r.ctx,
		`INSERT INTO lessons(name) VALUES ($1) returning id`,
		lesson.Name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *RepositoryV1) DeleteLesson(id int) error {
	_, err := r.db.Exec(r.ctx, `DELETE FROM lessons WHERE UserID = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryV1) UpdateLesson(lesson models.Lesson) error {
	_, err := r.db.Exec(
		r.ctx,
		`UPDATE lessons set name = $2 WHERE id = $1`,
		lesson.ID, lesson.Name,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryV1) AllLessons() ([]models.Lesson, error) {
	rows, err := r.db.Query(r.ctx, `SELECT * FROM lessons`)
	if err != nil {
		return []models.Lesson{}, err
	}
	defer rows.Close()

	var lessons []models.Lesson

	for rows.Next() {
		var lesson models.Lesson
		err := rows.Scan(&lesson.ID, &lesson.Name)
		if err != nil {
			return []models.Lesson{}, err
		}

		lessons = append(lessons, lesson)
	}

	return lessons, nil
}

func (r *RepositoryV1) GetLessonByID(id int) (models.Lesson, error) {
	var lesson models.Lesson
	err := r.db.QueryRow(r.ctx, `SELECT * FROM lessons WHERE id = $1`, id).Scan(
		&lesson.ID, &lesson.Name,
	)
	if err != nil {
		return models.Lesson{}, err
	}

	return lesson, nil
}
