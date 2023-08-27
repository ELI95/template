package ports

import "github.com/eli95/template/internal/user/domain/models"

type UserApplication interface {
	Create(registerRequest *models.RegisterRequest) error
	Login(credentials *models.AuthRequest) (authToken *models.AuthResponse, err error)
}
