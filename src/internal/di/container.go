package di

import (
	"tic-tac-toe/internal/domain/interfaces"
	domain "tic-tac-toe/internal/domain/services"
	"tic-tac-toe/internal/infrastructure/repository"
	infrastructure "tic-tac-toe/internal/infrastructure/services"
	"tic-tac-toe/internal/infrastructure/storage"
	"tic-tac-toe/internal/web/handlers"
	"tic-tac-toe/internal/web/server"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		fx.Annotate(
			domain.NewGameService,
			fx.As(new(interfaces.GameService)),
			fx.ResultTags(`name:"logic"`),
		),

		storage.NewGameStorage,
		repository.NewGameRepository,

		fx.Annotate(
			infrastructure.NewRepositoryGameService,
			fx.ParamTags(``, `name:"logic"`),
			fx.As(new(interfaces.GameService)),
		),

		handlers.NewGameHandler,
		server.NewHTTPServer,
	),
)
