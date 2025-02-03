package presentation

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type HelloPresenter struct {}

func NewHelloPresenter() *HelloPresenter {
	return &HelloPresenter{}
}

func (h *HelloPresenter) Hello(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}