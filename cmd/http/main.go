package main

import (
	"github.com/eli95/template/docs"
	"github.com/eli95/template/pkg/factories"
	"github.com/eli95/template/pkg/server"
)

func main() {
	factory := factories.NewFactory(
		"logs/log.csv",
		"./config/env.json",
	)

	// Create and Share the dependencies throw the factory [Only one instance]
	factory.InitializeValidator()
	configurator := factory.InitializeConfigurator()
	factory.InitializeLogger()
	factory.InitializeDatabase()

	middlewares := factory.BuildMiddlewaresHandlers()
	usersHandlers := factory.BuildUserHandlers()

	server := server.NewServer(middlewares, usersHandlers, configurator)

	config, err := configurator.GetConfig()
	if err != nil {
		panic(err)
	}

	docs.SwaggerInfo.Title = config.App.Name
	docs.SwaggerInfo.Description = config.App.Description
	docs.SwaggerInfo.Version = config.App.Version
	docs.SwaggerInfo.Host = config.App.Host
	docs.SwaggerInfo.BasePath = config.App.BaseURL

	err = server.Run(config.Server.Port)
	if err != nil {
		panic(err)
	}
}
