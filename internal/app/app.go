package app

import (
	"github.com/banggibima/go-gin-restful-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

type App struct {
	UserHandler *handlers.UserHandler
}

func NewApp(userHandler *handlers.UserHandler) *App {
	return &App{UserHandler: userHandler}
}

func (a *App) SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("/", a.UserHandler.GetUsersHandler)
			users.GET("/:id", a.UserHandler.GetUserByIDHandler)
			users.POST("/", a.UserHandler.CreateUserHandler)
			users.PUT("/:id", a.UserHandler.UpdateUserHandler)
			users.DELETE("/:id", a.UserHandler.DeleteUserHandler)
		}
	}
}
