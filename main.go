package main

import (
	"context"
	"log"
	"os"

	"github.com/DDD-Zenn/api/application/chat"
	"github.com/DDD-Zenn/api/domain/repoIF"
	"github.com/DDD-Zenn/api/external/service"
	"github.com/DDD-Zenn/api/infrastructure/repo"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"

	"github.com/DDD-Zenn/api/router"
	"github.com/DDD-Zenn/api/presentation"
	"github.com/DDD-Zenn/api/infrastructure/database"
)

var chatUsecase *chat.ChatUsecase

func main() {
	ctx := context.Background()

	// DB初期化
	database.InitDB()

	// APIKEY
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

	// Hello
	helloPresenter := presentation.NewHelloPresenter()

	// User
	var userRepo repoIF.UserRepoIF = repo.NewUserRepo(ctx)
	userService := service.NewUserService(userRepo)
	userPresenter := presentation.NewUserPresenter(userService)

	// Chat
	var chatRepo repoIF.Chat = repo.NewChatRepo(ctx)
	chatUsecase = chat.NewChatUsecase(chatRepo, geminiService)
	chatPresenter := presentation.NewChatPresenter(chatUsecase)

	engine := router.SetupRouter(
		helloPresenter,
		userPresenter,
		chatPresenter,
	)
	engine.Run(":8080")
}
