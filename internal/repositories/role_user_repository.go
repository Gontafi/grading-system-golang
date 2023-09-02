package repositories

import (
	"github.com/grading-system-golang/internal/models"
)

func (r *RepositoryV1) AddRoleUser(roleID int, userID int) (int, error) {
	var id int
	err :=
		r.db.QueryRow(r.ctx,
			`INSERT INTO role_user(role_id, user_id) VALUES($1, $2) RETURNING id`,
			roleID, userID).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *RepositoryV1) RemoveRoleUser(roleID int, userID int) error {
	_, err := r.db.Exec(r.ctx,
		`DELETE FROM role_user WHERE role_id = $1 AND user_id = $2`,
		roleID, userID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryV1) GetUsersForRole(roleID int) ([]models.User, error) {
	rows, err := r.db.Query(
		r.ctx,
		`
			SELECT u.id, u.username, u.name, u.created_at, u.updated_at 
			FROM users u JOIN role_user ru ON u.id = ru.user_id WHERE ru.role_id = $1
		`, roleID,
	)
	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Name, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return []models.User{}, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *RepositoryV1) GetRolesForUser(userID int) ([]models.Role, error) {
	rows, err := r.db.Query(
		r.ctx,
		`
			SELECT r.id, r.name, r.status_id 
			FROM roles r JOIN role_user ru ON r.id = ru.role_id WHERE ru.user_id = $1
		`, userID,
	)
	if err != nil {
		return []models.Role{}, err
	}
	defer rows.Close()

	var roles []models.Role
	for rows.Next() {
		var role models.Role
		err := rows.Scan(&role.ID, &role.Name, &role.StatusID)
		if err != nil {
			return []models.Role{}, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}

func (r *RepositoryV1) GetUserRole(userID int) (models.Role, error) {
	var role models.Role
	err := r.db.QueryRow(
		r.ctx,
		`
			SELECT r.id, r.name, r.status_id 
			FROM roles r 
			JOIN role_user ru ON r.id = ru.role_id 
			WHERE user_id = $1
			`, userID,
	).Scan(&role.ID, &role.Name, &role.StatusID)

	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}
