package repositories

import (
	"github.com/grading-system-golang/internal/models"
	"time"
)

func (r *RepositoryV1) GetTopRating(period time.Duration, limit int) ([]models.Rating, error) {

	start := time.Now().Add(-period)
	end := time.Now()

	rows, err := r.db.Query(
		r.ctx,
		`
		SELECT
			s.id AS student_id,
			s.name AS student_name,
			COALESCE(SUM(m.home_work_grade + m.attendance_grade), 0) AS score
		FROM users s
		LEFT JOIN homeworks hw ON s.id = hw.student_id
		LEFT JOIN marks m ON hw.id = m.home_work_id
		WHERE m.date > $1 AND m.date < $2
		GROUP BY s.id, s.name
		ORDER BY score DESC
		LIMIT $3;
		`, start, end, limit,
	)
	if err != nil {
		return []models.Rating{}, err
	}

	var ratings []models.Rating

	for rows.Next() {
		var rating models.Rating
		err := rows.Scan(
			&rating.StudentID,
			&rating.StudentName,
			&rating.Score,
		)
		if err != nil {
			return []models.Rating{}, err
		}

		ratings = append(ratings, rating)
	}

	return ratings, nil
}

func (r *RepositoryV1) GetTopRatingByLesson(lessonID int, period time.Duration, limit int) ([]models.Rating, error) {

	start := time.Now().Add(-period)
	end := time.Now()

	rows, err := r.db.Query(
		r.ctx,
		`
		SELECT
			s.id AS student_id,
			s.name AS student_name,
			COALESCE(SUM(m.home_work_grade + m.attendance_grade), 0) AS score
		FROM users s
		LEFT JOIN homeworks hw ON s.id = hw.student_id
		LEFT JOIN marks m ON hw.id = m.home_work_id
		WHERE m.date > $2 AND m.date < $3 AND hw.lesson_id = $1
		GROUP BY s.id, s.name
		ORDER BY score DESC
		LIMIT $4;
		`,
		lessonID, start, end, limit,
	)
	if err != nil {
		return []models.Rating{}, err
	}

	var ratings []models.Rating

	for rows.Next() {
		var rating models.Rating
		err := rows.Scan(
			&rating.StudentID,
			&rating.StudentName,
			&rating.Score,
		)
		if err != nil {
			return nil, err
		}

		ratings = append(ratings, rating)
	}

	return ratings, nil
}
