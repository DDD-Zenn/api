package router

import (
	"github.com/DDD-Zenn/api/presentation"
	"github.com/gin-gonic/gin"
)

func RegisterHelloRoutes(router *gin.Engine, UserPresenter *presentation.HelloPresenter) {
	router.GET("/", UserPresenter.Hello)
}