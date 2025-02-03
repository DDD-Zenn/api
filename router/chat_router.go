package router

import (
	"github.com/DDD-Zenn/api/presentation"
	"github.com/gin-gonic/gin"
)

func RegisterChatRoutes(router *gin.Engine, ChatPresenter *presentation.ChatPresenter) {
	chatGroup := router.Group("/chat")
	{
		chatGroup.POST("", ChatPresenter.PostChat)
	}
}