package router

import (
	"test-praweda-app-eng/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitEcommerceRouter(router *gin.Engine) {
	EcommerceRouter := router.Group("/ecommerce")
	EcommerceHandler := handler.NewEcommerceHandler()

	EcommerceRouter.POST("/diskon", EcommerceHandler.DiskonCalculation)
}
