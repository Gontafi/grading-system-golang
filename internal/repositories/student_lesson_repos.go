package repositories

import (
	"github.com/grading-system-golang/internal/models"
)

func (r *RepositoryV1) AddStudentToLesson(studentID int, lessonID int) (int, error) {
	var id int
	err :=
		r.db.QueryRow(r.ctx,
			`INSERT INTO student_lesson(student_id, lesson_id) VALUES($1, $2)`,
			studentID, lessonID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *RepositoryV1) RemoveStudentFromLesson(studentID int, lessonID int) error {
	_, err := r.db.Exec(r.ctx,
		`DELETE FROM student_lesson WHERE student_id = $1 AND lesson_id = $2`,
		studentID, lessonID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryV1) GetStudentsForLesson(lessonID int) ([]models.User, error) {
	rows, err := r.db.Query(
		r.ctx,
		`SELECT s.id, s.name, s.username FROM users s JOIN lessons l ON s.id = l.id WHERE l.id = $1`,
		lessonID)
	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Username)
		if err != nil {
			return []models.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *RepositoryV1) GetLessonsForStudent(studentID int) ([]models.Lesson, error) {
	rows, err := r.db.Query(
		r.ctx,
		`SELECT l.id, l.name FROM lessons l JOIN users s ON l.id = s.id WHERE s.id = $1`,
		studentID)
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

func (r *RepositoryV1) GetStudentLesson(studentID int, lessonID int) (models.StudentLesson, error) {
	var studentLesson models.StudentLesson
	err := r.db.QueryRow(
		r.ctx,
		`SELECT id, student_id, lesson_id FROM student_lesson WHERE student_id = $1 AND lesson_id = $2`,
		studentID, lessonID,
	).Scan(&studentLesson.ID, &studentLesson.StudentID, &studentLesson.LessonID)

	if err != nil {
		return models.StudentLesson{}, err
	}

	return studentLesson, nil
}
