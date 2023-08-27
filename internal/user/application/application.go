package application

import (
	"github.com/eli95/template/internal/user/domain/ports"
	configurator "github.com/eli95/template/pkg/config/domain/ports"
	logger "github.com/eli95/template/pkg/logger/domain/ports"
	val "github.com/eli95/template/pkg/validator/domain/ports"
)

type UserApp struct {
	repo         ports.UserRepository
	validator    val.ValidatorApplication
	logger       logger.LoggerApplication
	configurator configurator.ConfigApplication
}

var _ ports.UserApplication = (*UserApp)(nil)

func NewUserApp(
	repo ports.UserRepository,
	validator val.ValidatorApplication,
	logger logger.LoggerApplication,
	configurator configurator.ConfigApplication,
) *UserApp {
	return &UserApp{
		repo:         repo,
		validator:    validator,
		logger:       logger,
		configurator: configurator,
	}
}
