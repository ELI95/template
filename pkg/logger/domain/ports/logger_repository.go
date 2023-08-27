package ports

import "github.com/eli95/template/pkg/logger/domain/models"

type LoggerRepository interface {
	Save(log *models.Log) error
}
