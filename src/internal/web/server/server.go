package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"tic-tac-toe/internal/web/handlers"
	"tic-tac-toe/internal/web/middleware"
	"tic-tac-toe/public"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewHTTPServer(
	h *handlers.GameHandler,
	authH *handlers.AuthHandler,
	i *handlers.GameInfoHandler,
	authMW *middleware.UserAuthenticator,
	lifecycle fx.Lifecycle,
) *gin.Engine {
	r := gin.Default()

	setupMiddleware(r)

	server := setupRoutes(r, h, i, authH, authMW)

	setupLifecycle(server, lifecycle)

	return r
}

func setupMiddleware(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
}

func setupRoutes(r *gin.Engine, h *handlers.GameHandler, i *handlers.GameInfoHandler, authH *handlers.AuthHandler, authMW *middleware.UserAuthenticator) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.GET("/", func(c *gin.Context) {
		index, err := public.Root.ReadFile("index.html")
		if err != nil {
			c.String(http.StatusInternalServerError, "не удалось прочитать index.html")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", index)
	})

	r.POST("/signup", authH.SignUp)
	r.POST("/login", authH.Login)
	r.POST("/auth/refresh", authH.Refresh)
	r.POST("/auth/token/refresh", authH.RefreshAccessToken)

	authorized := r.Group("/")
	authorized.Use(authMW.Handler())
	{
		authorized.POST("/game/create", h.CreateGame)
		authorized.GET("/games/available", h.GetAvailableGames)
		authorized.POST("/game/:uuid/join", h.JoinGame)
		authorized.POST("/game/:uuid", h.HandleMove)
		authorized.GET("/game/:uuid", h.GetGame)

		authorized.GET("/user/me", authH.GetUserInfo)

		authorized.GET("/games/history", i.GetUserHistory)

		authorized.GET("/games/leaders", i.GetLeaderboard)
	}

	return &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
}

func setupLifecycle(srv *http.Server, lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					fmt.Printf("Ошибка запуска сервера: %v\n", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Остановка сервера...")
			return srv.Shutdown(ctx)
		},
	})
}
