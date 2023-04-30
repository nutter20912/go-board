package bootstrap

import (
	"board/middlewares"
	"board/providers"
	"board/routers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type App struct {
	router *gin.Engine
}

func NewServer() *http.Server {
	app := App{router: gin.New()}
	app.
		registerProvider().
		registerMiddleware().
		registerRouter()

	return &http.Server{
		Addr:    os.Getenv("APP_HOST") + os.Getenv("APP_PORT"),
		Handler: app.router,
	}
}

// 註冊服務
func (app *App) registerProvider() *App {
	app.router.Use(gin.Recovery())
	app.router.Use(providers.Logger)

	return app
}

// 註冊中間件
func (app *App) registerMiddleware() *App {
	app.router.Use(middlewares.ErrorHandler)
	app.router.Use(middlewares.RequestLogger)

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
