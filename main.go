package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/DDD-Zenn/api/application/chat"
	chatCmd "github.com/DDD-Zenn/api/application/chat/command"
	"github.com/DDD-Zenn/api/domain/repoIF"
	"github.com/DDD-Zenn/api/external/service"
	"github.com/DDD-Zenn/api/infrastructure/repo"
	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var chatUsecase *chat.ChatUsecase

func main() {

	ctx := context.Background()
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY is not set")
	}

	// genai.Client を初期化
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}
	defer client.Close()

	// Gemini サービスを初期化
	geminiService := service.NewGeminiService(client)
	var cRepo repoIF.Chat = repo.NewChatRepo(ctx)
	chatUsecase = chat.NewChatUsecase(cRepo, geminiService)
	engine := gin.Default()

	engine.POST("/chat", func(c *gin.Context) {
		var cmd chatCmd.CreateChatCommand

		if err := c.ShouldBindJSON(&cmd); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validCmd, err := chatCmd.NewChatCreate(cmd)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		output, err := chatUsecase.CreateChat(ctx, validCmd)
		if err != nil {
			log.Printf("Failed to list tasks: %v", err)
		}

		respBytes, err := json.Marshal(output)
		if err != nil {
			log.Printf("Failed to marshal tasks: %v", err)
		}

		c.Data(http.StatusOK, "application/json; charset=utf-8", respBytes)
	})

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	engine.Run(":8080")
}
