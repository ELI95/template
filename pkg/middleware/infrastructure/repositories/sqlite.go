package repositories

import (
	"github.com/eli95/template/pkg/config/domain/ports"
	logger "github.com/eli95/template/pkg/logger/domain/ports"
	mdl "github.com/eli95/template/pkg/middleware/domain/ports"
)

type SqliteMdlRepository struct {
	config ports.ConfigApplication
	logger logger.LoggerApplication
}

var _ mdl.MiddlewareRepository = (*SqliteMdlRepository)(nil)

func NewSqliteMdlRepository(config ports.ConfigApplication, logger logger.LoggerApplication) *SqliteMdlRepository {
	return &SqliteMdlRepository{
		config: config,
		logger: logger,
	}
}

func (m *SqliteMdlRepository) Authenticate(token string) (err error) {
	//TODO implement me
	panic("implement me")
}
