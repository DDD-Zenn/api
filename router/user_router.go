package router

import (
	"github.com/DDD-Zenn/api/presentation"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, UserPresenter *presentation.UserPresenter) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("", UserPresenter.Create)
		userGroup.GET("/:uid", UserPresenter.FindByUID)
		userGroup.PUT("/:uid", UserPresenter.Update)
		userGroup.DELETE("/:uid", UserPresenter.Delete)
		userGroup.GET("/post", UserPresenter.GetPost)
	}
}
