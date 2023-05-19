package bootstrap

import (
	"board/config"
	"board/middlewares"
	"board/providers"
	"board/routers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
	config config.Config
}

// 建立 server 實例
func NewServer(config config.Config) *http.Server {
	app := App{
		router: gin.New(),
		config: config,
	}
	app.
		registerProvider().
		registerMiddleware().
		registerRouter()

	return &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.App.Host, config.App.Port),
		Handler: app.router,
	}
}

// 註冊服務
func (app *App) registerProvider() *App {
	app.router.Use(gin.Recovery())

	logger := providers.Logger()
	db := providers.DB(app.config, logger)
	hub := providers.Hub()
	channelManager := providers.ChannelManager(app.config)

	app.router.Use(func(c *gin.Context) {
		c.Set("config", app.config)
		c.Set("log", logger)
		c.Set("db", db)
		c.Set("hub", hub)
		c.Set("channelManager", channelManager)
		c.Next()
	})

	return app
}

// 註冊中間件
func (app *App) registerMiddleware() *App {
	app.router.Use(middlewares.ErrorHandler)
	//app.router.Use(middlewares.RequestLogger)

	return app
}

// 註冊路由
func (app *App) registerRouter() *App {
	for _, controller := range routers.Apis {
		app.router.Handle(
			controller.Method,
			controller.Path,
			controller.Action,
		)
	}

	return app
}
