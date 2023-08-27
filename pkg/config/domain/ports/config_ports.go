package ports

import (
	"github.com/eli95/template/pkg/config/domain/models"
	"os"
)

type ConfigApplication interface {
	Config() error
	GetConfig() (*models.Config, error)
}

type ConfigRepository interface {
	GetConfigFile() (*os.File, error)
}
