package main

import (
	"tic-tac-toe/internal/di"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.Module,

		fx.Invoke(func(r *gin.Engine) {
		}),
	).Run()
}
