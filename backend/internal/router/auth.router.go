package router

import (
	"test-praweda-app-eng/internal/handler"
	"test-praweda-app-eng/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitAuthRouter(router *gin.Engine, db *pgxpool.Pool) {
	AuthRouter := router.Group("/auth")
	AuthRepository := repository.NewAuthRepository(db)
	AuthHandler := handler.NewAuthHandler(AuthRepository)

	AuthRouter.POST("/login", AuthHandler.Login)
	AuthRouter.GET("/me", AuthHandler.Me)
	AuthRouter.DELETE("/logout", AuthHandler.Logout)
}
