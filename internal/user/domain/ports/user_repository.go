package ports

import "github.com/eli95/template/internal/user/domain/models"

type UserRepository interface {
	Save(user *models.User) error
	FindByCredentials(credentials *models.AuthRequest) (*models.User, error)
}
