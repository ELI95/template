package repositories

import (
	"github.com/eli95/template/internal/user/domain/models"
)

func (repo *UserSqlite) FindByCredentials(credentials *models.AuthRequest) (*models.User, error) {
	req := models.User{}
	req.Email = credentials.Email
	req.Password = req.Hash256Password(credentials.Password)

	user := models.User{}
	sql := `select id, email from user where email=? and password = ? limit 1`
	err := repo.db.QueryRow(sql, req.Email, req.Password).Scan(&user.ID, &user.Email)
	if err != nil {
		repo.logger.Error("Error finding user", err.Error())
		return nil, err
	}
	return &user, nil
}
