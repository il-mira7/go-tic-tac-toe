package di

import (
	app "tic-tac-toe/internal/app/services"
	domain "tic-tac-toe/internal/domain/services"
	"tic-tac-toe/internal/infrastructure/config"
	"tic-tac-toe/internal/infrastructure/database"
	"tic-tac-toe/internal/infrastructure/jwt"
	"tic-tac-toe/internal/infrastructure/repository"

	"tic-tac-toe/internal/web/handlers"
	"tic-tac-toe/internal/web/middleware"
	"tic-tac-toe/internal/web/server"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		config.NewConfig,
		func(cfg *config.Config) *config.DatabaseConfig { return &cfg.DB },

		database.NewPostgresPool,
		repository.NewPostgresGameRepository,
		repository.NewPostgresUserRepository,
		domain.NewUserService,
		domain.NewAuthService,
		handlers.NewAuthHandler,
		middleware.NewUserAuthenticator,
		app.NewRepositoryGameUseManager,
		domain.NewGameService,
		handlers.NewGameHandler,
		server.NewHTTPServer,

		func(cfg *config.Config) *config.JWTConfig { return &cfg.JWT },
		jwt.NewJwtProvider,
		handlers.NewGameInfoHandler,
		app.NewRepositoryGameInfoManager,
	),
)
