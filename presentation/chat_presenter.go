package presentation

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/DDD-Zenn/api/application/chat"
	chatCmd "github.com/DDD-Zenn/api/application/chat/command"
)

type ChatPresenter struct {
	chatUsecase *chat.ChatUsecase
}

func NewChatPresenter(chatUsecase *chat.ChatUsecase) *ChatPresenter {
	return &ChatPresenter{chatUsecase: chatUsecase}
}

func (h *ChatPresenter) PostChat(context *gin.Context) {
	var cmd chatCmd.CreateChatCommand

	if err := context.ShouldBindJSON(&cmd); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validCmd, err := chatCmd.NewChatCreate(cmd)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := h.chatUsecase.CreateChat(context, validCmd)
	if err != nil {
		log.Printf("Failed to list tasks: %v", err)
	}

	respBytes, err := json.Marshal(output)
	if err != nil {
		log.Printf("Failed to marshal tasks: %v", err)
	}

	context.Data(http.StatusOK, "application/json; charset=utf-8", respBytes)
}