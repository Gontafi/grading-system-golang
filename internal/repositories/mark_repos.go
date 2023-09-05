package repositories

import (
	"github.com/grading-system-golang/internal/models"
)

func (r *RepositoryV1) CreateMark(mark models.Mark) (int, error) {
	var markID int
	err := r.db.QueryRow(
		r.ctx,
		`
		INSERT INTO marks (home_work_id, home_work_grade, attendance_grade, date)
		VALUES ($1, $2, $3, $4)
		RETURNING id
		`,
		mark.HomeWorkID, mark.HomeWorkGrade, mark.AttendanceGrade, mark.Date).Scan(&markID)

	if err != nil {
		return 0, err
	}

	return markID, nil
}

func (r *RepositoryV1) GetMarkByID(markID int) (models.Mark, error) {
	var mark models.Mark
	err := r.db.QueryRow(
		r.ctx,
		`
		SELECT id, home_work_id, home_work_grade, attendance_grade, date
		FROM marks WHERE id = $1`,
		markID,
	).Scan(&mark.ID, &mark.HomeWorkID, &mark.HomeWorkGrade, &mark.AttendanceGrade, &mark.Date)

	if err != nil {
		return models.Mark{}, err
	}

	return mark, nil
}

func (r *RepositoryV1) GetAllMarks() ([]models.Mark, error) {
	var marks []models.Mark
	rows, err := r.db.Query(
		r.ctx,
		`
		SELECT id, home_work_id, home_work_grade, attendance_grade, date
		FROM marks`,
	)
	if err != nil {
		return []models.Mark{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var mark models.Mark
		err := rows.Scan(&mark.ID, &mark.HomeWorkID, &mark.HomeWorkGrade, &mark.AttendanceGrade, &mark.Date)
		if err != nil {
			return []models.Mark{}, err
		}
		marks = append(marks, mark)
	}

	return marks, nil
}

func (r *RepositoryV1) DeleteMark(markID int) error {
	_, err := r.db.Exec(r.ctx, `DELETE FROM marks WHERE id = $1`, markID)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryV1) UpdateMark(mark models.Mark) error {
	_, err := r.db.Exec(
		r.ctx,
		`UPDATE marks set home_work_id = $2, home_work_grade = $3, attendance_grade = $4, date = $5 WHERE id = $1`,
		mark.ID, mark.HomeWorkID, mark.HomeWorkGrade, mark.AttendanceGrade, mark.Date,
	)
	if err != nil {
		return err
	}

	return nil
}
