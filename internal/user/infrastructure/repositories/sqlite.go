package repositories

import (
	"github.com/eli95/template/internal/user/domain/ports"
	config "github.com/eli95/template/pkg/config/domain/ports"
	logger "github.com/eli95/template/pkg/logger/domain/ports"
	"github.com/jmoiron/sqlx"
)

type UserSqlite struct {
	db           *sqlx.DB
	configurator config.ConfigApplication
	logger       logger.LoggerApplication
}

var _ ports.UserRepository = (*UserSqlite)(nil)

func NewUserSqlite(config config.ConfigApplication, logger logger.LoggerApplication, db *sqlx.DB) *UserSqlite {
	return &UserSqlite{
		db:           db,
		configurator: config,
		logger:       logger,
	}
}
