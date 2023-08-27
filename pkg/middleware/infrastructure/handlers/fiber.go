package handlers

import (
	configurator "github.com/eli95/template/pkg/config/domain/ports"
	logger "github.com/eli95/template/pkg/logger/domain/ports"
	"github.com/eli95/template/pkg/middleware/domain/ports"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

type FiberMdlHandlers struct {
	service      ports.MiddlewareApplication
	logger       logger.LoggerApplication
	configurator configurator.ConfigApplication
}

var _ ports.MiddlewareHandlers = (*FiberMdlHandlers)(nil)

func NewFiberMdlHandlers(service ports.MiddlewareApplication, logger logger.LoggerApplication, config configurator.ConfigApplication) *FiberMdlHandlers {
	return &FiberMdlHandlers{
		service:      service,
		logger:       logger,
		configurator: config,
	}
}

func (f *FiberMdlHandlers) Authenticate() fiber.Handler {
	config, err := f.configurator.GetConfig()
	if err != nil {
		f.logger.Error("Error getting config", err)
		return func(ctx *fiber.Ctx) error {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.JWT.Secret),
	})
}
