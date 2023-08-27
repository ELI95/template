package factories

import (
	userApp "github.com/eli95/template/internal/user/application"
	userHdl "github.com/eli95/template/internal/user/infrastructure/handlers"
	userRepo "github.com/eli95/template/internal/user/infrastructure/repositories"
	configApp "github.com/eli95/template/pkg/config/application"
	configRepo "github.com/eli95/template/pkg/config/infrastructure/repositories"
	loggerApp "github.com/eli95/template/pkg/logger/application"
	loggerRepo "github.com/eli95/template/pkg/logger/infrastructure/repositories"
	mdlApp "github.com/eli95/template/pkg/middleware/application"
	mdlHdl "github.com/eli95/template/pkg/middleware/infrastructure/handlers"
	mdlRepo "github.com/eli95/template/pkg/middleware/infrastructure/repositories"
	valApp "github.com/eli95/template/pkg/validator/application"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Factory struct {
	//Variables
	logFilePath    string
	configFilePath string
	//Packages
	validator    *valApp.Validator
	configurator *configApp.ConfigService
	logger       *loggerApp.Logger
	db           *sqlx.DB
}

func NewFactory(logFilePath string, configFilePath string) *Factory {
	return &Factory{
		logFilePath:    logFilePath,
		configFilePath: configFilePath,
	}
}

func (f *Factory) InitializeDatabase() *sqlx.DB {
	if f.db != nil {
		return f.db
	}

	configurator := f.InitializeConfigurator()
	cfg, err := configurator.GetConfig()
	if err != nil {
		panic(err)
	}
	f.db = sqlx.MustConnect(cfg.Database.DriverName, cfg.Database.DataSourceName)
	return f.db
}

func (f *Factory) InitializeValidator() *valApp.Validator {
	if f.validator == nil {
		app := valApp.NewValidator()
		f.validator = app
		return app
	}
	return f.validator
}

func (f *Factory) InitializeLogger() *loggerApp.Logger {
	if f.configurator == nil {
		validator := f.InitializeValidator()
		path := f.logFilePath

		repo := loggerRepo.NewCSVFile(path)
		app := loggerApp.NewLogger(repo, validator)
		f.logger = app
		return app
	}
	return f.logger
}

func (f *Factory) InitializeConfigurator() *configApp.ConfigService {
	if f.configurator == nil {
		validator := f.InitializeValidator()
		logger := f.InitializeLogger()
		path := f.configFilePath

		repo := configRepo.NewJSONRepository(path)
		app := configApp.NewConfigService(repo, validator, logger)
		err := app.Config()
		if err != nil {
			panic(err)
		}
		f.configurator = app
		return app
	}
	return f.configurator
}

func (f *Factory) BuildMiddlewaresHandlers() *mdlHdl.FiberMdlHandlers {
	configurator := f.InitializeConfigurator()
	logger := f.InitializeLogger()
	validator := f.InitializeValidator()

	repo := mdlRepo.NewSqliteMdlRepository(configurator, logger)
	app := mdlApp.NewFiberMiddlewares(repo, validator, logger)
	return mdlHdl.NewFiberMdlHandlers(app, logger, configurator)
}

func (f *Factory) BuildUserHandlers() *userHdl.UserHdl {
	configurator := f.InitializeConfigurator()
	logger := f.InitializeLogger()
	validator := f.InitializeValidator()
	dbClient := f.InitializeDatabase()

	repo := userRepo.NewUserSqlite(configurator, logger, dbClient)
	app := userApp.NewUserApp(repo, validator, logger, configurator)
	return userHdl.NewUserHdl(app, logger)
}
