package repositories

import (
	"github.com/eli95/template/internal/user/domain/models"
)

func (repo *UserSqlite) Save(user *models.User) error {
	sql := `insert into user
    	(email, password, created_at)
		VALUES (?, ?, ?)`

	res, err := repo.db.Exec(sql, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = id

	return nil
}
