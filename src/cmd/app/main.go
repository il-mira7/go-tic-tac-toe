package main

import (
	"log"
	"tic-tac-toe/internal/di"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	fx.New(
		di.Module,

		fx.Invoke(func(r *gin.Engine) {
		}),
	).Run()
}
