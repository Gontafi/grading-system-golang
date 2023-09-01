package repositories

import (
	"github.com/grading-system-golang/internal/models"
)

func (r *RepositoryV1) AddUser(user models.User) (int, error) {
	var id int

	err := r.db.QueryRow(
		r.ctx,
		`INSERT INTO users(username, password_hash, role, created_at, updated_at) 
			VALUES ($1, $2, $3, $4, $5) returning id`,
		user.Username, user.PasswordHash, user.Role, user.CreatedAt, user.UpdatedAt,
	).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *RepositoryV1) DeleteUser(id int) error {
	_, err := r.db.Exec(r.ctx, `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryV1) UpdateUser(user models.User) error {
	_, err := r.db.Exec(
		r.ctx,
		`UPDATE users set username = $2, password_hash = $3, role = $4, updated_at = $5 WHERE id = $1`,
		user.ID, user.Username, user.PasswordHash, user.Role, user.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryV1) AllUsers() ([]models.User, error) {
	rows, err := r.db.Query(r.ctx, `SELECT * FROM users`)
	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return []models.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *RepositoryV1) GetUserByID(id int) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(r.ctx, `SELECT * FROM users WHERE id = $1`, id).Scan(
		&user.ID, &user.Username, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *RepositoryV1) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(r.ctx, `SELECT * FROM users WHERE username = $1`, username).Scan(
		&user.ID, &user.Username, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
