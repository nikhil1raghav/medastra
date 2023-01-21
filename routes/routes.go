package routes

import (
	"github.com/gin-gonic/gin"
	"medastra/controller"
	"net/http"
)

func NewRouter(controller *controller.Controller) *gin.Engine {
	service := gin.Default()
	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "medastra")
	})
	service.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Live")
	})
	service.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, "Not found")
	})
	router := service.Group("/api/v1/")
	userRouter := router.Group("/user/")
	userRouter.POST("/userprofile/", controller.CreateUserProfile)
	return service
}