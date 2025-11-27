package router

import (
	"net/http"
	"test-praweda-app-eng/internal/middleware"
	"test-praweda-app-eng/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitRouter(db *pgxpool.Pool) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware)

	// make directori public accesible
	router.Static("/img", "public")

	InitAuthRouter(router, db)
	InitEcommerceRouter(router)

	// handle not found route
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, model.Response{
			IsSuccess: false,
			Code:      http.StatusNotFound,
			Msg:       "Page not found!",
		})
	})

	return router
}
