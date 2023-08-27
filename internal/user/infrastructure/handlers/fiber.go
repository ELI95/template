package handlers

import (
	"github.com/eli95/template/internal/user/domain/ports"
	logger "github.com/eli95/template/pkg/logger/domain/ports"
)

type UserHdl struct {
	app    ports.UserApplication
	logger logger.LoggerApplication
}

var _ ports.UserHandlers = (*UserHdl)(nil)

func NewUserHdl(app ports.UserApplication, logger logger.LoggerApplication) *UserHdl {
	return &UserHdl{app: app, logger: logger}
}
