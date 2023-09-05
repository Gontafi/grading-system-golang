package repositories

import (
	"github.com/grading-system-golang/internal/models"
)

func (r *RepositoryV1) AddLesson(lesson models.Lesson) (int, error) {
	var id int

	err := r.db.QueryRow(
		r.ctx,
		`INSERT INTO lessons(name, teacher_id) VALUES ($1, $2) returning id`,
		lesson.Name, lesson.TeacherID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *RepositoryV1) DeleteLesson(id int) error {
	_, err := r.db.Exec(r.ctx, `DELETE FROM lessons WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryV1) UpdateLesson(lesson models.Lesson) error {
	_, err := r.db.Exec(
		r.ctx,
		`UPDATE lessons set name = $2, teacher_id = $3 WHERE id = $1`,
		lesson.ID, lesson.Name, lesson.TeacherID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryV1) AllLessons() ([]models.Lesson, error) {
	rows, err := r.db.Query(r.ctx, `SELECT id, name, teacher_id FROM lessons`)
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
	err := r.db.QueryRow(r.ctx, `SELECT id, name, teacher_id FROM lessons WHERE id = $1`, id).Scan(
		&lesson.ID, &lesson.Name,
	)
	if err != nil {
		return models.Lesson{}, err
	}

	return lesson, nil
}

func (r *RepositoryV1) GetLessonFromMark(markID int) (models.Lesson, error) {
	var lesson models.Lesson
	err := r.db.QueryRow(
		r.ctx,
		`SELECT l.id, l.name, l.teacher_id 
				FROM marks m 
			    JOIN homeworks h ON h.id = m.home_work_id
			    JOIN lessons l ON l.id = h.lesson_id
				WHERE m.id = $1
			    `, markID,
	).Scan(lesson.ID, lesson.Name, lesson.TeacherID)
	if err != nil {
		return models.Lesson{}, err
	}

	return lesson, nil
}
