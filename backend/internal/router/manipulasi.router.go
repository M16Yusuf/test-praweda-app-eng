package router

import (
	"test-praweda-app-eng/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitManipulasiRouter(router *gin.Engine) {
	ManipulasiRouter := router.Group("/manipulasi")
	ManipulasiHandler := handler.NewManipulasiHandler()

	ManipulasiRouter.GET("", ManipulasiHandler.ManipulasiString)
}
