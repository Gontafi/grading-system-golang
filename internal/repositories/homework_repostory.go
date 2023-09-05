package repositories

import "github.com/grading-system-golang/internal/models"

func (r *RepositoryV1) AddHomeWork(work models.HomeWork) (int, error) {
	var id int

	err := r.db.QueryRow(
		r.ctx,
		`INSERT INTO homeworks(student_id, lesson_id, status_id, home_task_response)
			VALUES ($1, $2, $3, $4)`,
		work.StudentID, work.LessonID, work.StatusID, work.HomeTaskResponse,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *RepositoryV1) GetHomeWorkById(id int) (models.HomeWork, error) {
	var homeWork models.HomeWork
	err := r.db.QueryRow(
		r.ctx,
		`SELECT id, student_id, lesson_id, status_id, home_task_response FROM homeworks WHERE id = $1`, id).Scan(
		&homeWork.ID, &homeWork.StudentID, &homeWork.LessonID, &homeWork.StatusID, &homeWork.HomeTaskResponse,
	)
	if err != nil {
		return models.HomeWork{}, err
	}

	return homeWork, nil
}

func (r *RepositoryV1) GetAllHomeWorks() ([]models.HomeWork, error) {
	var homeWorks []models.HomeWork

	rows, err := r.db.Query(
		r.ctx,
		`SELECT id, student_id, lesson_id, status_id, home_task_response FROM homeworks`)
	if err != nil {
		return []models.HomeWork{}, err
	}

	for rows.Next() {
		var homeWork models.HomeWork
		err := rows.Scan(
			&homeWork.ID, &homeWork.StudentID, &homeWork.LessonID, &homeWork.StatusID, &homeWork.HomeTaskResponse,
		)
		if err != nil {
			return []models.HomeWork{}, err
		}

		homeWorks = append(homeWorks, homeWork)
	}

	return homeWorks, nil
}

func (r *RepositoryV1) UpdateHomeWork(work models.HomeWork) error {
	_, err := r.db.Exec(
		r.ctx,
		`UPDATE homeworks SET student_id = $2, lesson_id = $3, status_id = $4, home_task_response = $5
                 WHERE id = $1`, work.ID, work.StudentID, work.LessonID, work.StatusID, work.HomeTaskResponse,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryV1) DeleteHomeWork(id int) error {
	_, err := r.db.Exec(
		r.ctx,
		`DELETE FROM homeworks WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
