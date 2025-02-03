package router

import (
	"github.com/DDD-Zenn/api/presentation"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	helloPresenter *presentation.HelloPresenter,
	userPresenter *presentation.UserPresenter,
	chatPresenter *presentation.ChatPresenter,
	) *gin.Engine {
	engine := gin.Default()

	RegisterHelloRoutes(engine, helloPresenter)
	RegisterUserRoutes(engine,  userPresenter)
	RegisterChatRoutes(engine,  chatPresenter)

	return engine
}