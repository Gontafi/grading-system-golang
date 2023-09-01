package repositories

import (
	"github.com/grading-system-golang/internal/models"
)

func (r *RepositoryV1) AddRole(role models.Role) (int, error) {

	_, err := r.db.Exec(r.ctx, `INSERT INTO roles(role_name) VALUES ($1)`, role.RoleName)
	if err != nil {
		return 0, err
	}

	return role.ID, nil
}

func (r *RepositoryV1) DeleteRole(id int) error {
	_, err := r.db.Exec(r.ctx, `DELETE FROM roles WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryV1) UpdateRole(role models.Role) error {
	_, err := r.db.Exec(r.ctx, `UPDATE roles set role_name = $2 WHERE id = $1`, role.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryV1) AllRoles() ([]models.Role, error) {
	rows, err := r.db.Query(r.ctx, `SELECT id, role_name FROM roles`)
	if err != nil {
		return []models.Role{}, err
	}
	defer rows.Close()

	var roles []models.Role

	for rows.Next() {
		var role models.Role
		err := rows.Scan(&role.ID, &role.RoleName)
		if err != nil {
			return []models.Role{}, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}

func (r *RepositoryV1) GetRoleByID(id int) (models.Role, error) {
	var role models.Role
	err := r.db.QueryRow(r.ctx, `SELECT id, role_name FROM roles WHERE id = $1`, id).Scan(
		&role.ID, &role.RoleName,
	)
	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}
