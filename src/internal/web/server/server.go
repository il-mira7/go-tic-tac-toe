package server

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"os"
	"tic-tac-toe/internal/web/handlers"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewHTTPServer(h *handlers.GameHandler, lifecycle fx.Lifecycle, fs embed.FS) *gin.Engine {
	r := gin.Default()

	setupMiddleware(r)

	server := setupRoutes(r, h, fs)

	setupLifecycle(server, lifecycle)

	return r
}

func setupMiddleware(r *gin.Engine) {
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
}

func setupRoutes(r *gin.Engine, h *handlers.GameHandler, fs embed.FS) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//r.StaticFile("/", "./public/index.html")
	r.GET("/", func(c *gin.Context) {
		index, err := fs.ReadFile("index.html")
		if err != nil {
			fmt.Printf("Ошибка embed: %v\n", err)
			c.String(http.StatusNotFound, "File not found")
			return
		}
		c.Data(http.StatusOK, "text/html; charset=utf-8", index)
	})

	r.POST("/game/:uuid", h.HandleMove)

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
