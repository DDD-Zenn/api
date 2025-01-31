package chat

import (
	"context"

	"github.com/DDD-Zenn/api/application/chat/command"
	"github.com/DDD-Zenn/api/domain/model"
	"github.com/DDD-Zenn/api/domain/repoIF"
	"github.com/DDD-Zenn/api/external/serviceIF"
	"github.com/DDD-Zenn/api/pkg/utils"
)

type ChatUsecase struct {
	repo      repoIF.Chat
	geminiSvc serviceIF.Gemini
}

func NewChatUsecase(repo repoIF.Chat, gemini serviceIF.Gemini) *ChatUsecase {
	return &ChatUsecase{
		repo:      repo,
		geminiSvc: gemini,
	}
}

func (uc *ChatUsecase) CreateChat(ctx context.Context, cmd command.CreateChatCommand) (*ChatDTO, error) {
	response, err := uc.geminiSvc.GenerateResponse(cmd.Prompt)
	if err != nil {
		return nil, err
	}

	//第3引数のidは本来はユーザーidを入れたいが，今はダミーでランダムなuuidを生成している
	chat, err := model.NewChat(ctx, utils.GenId(), utils.GenId(), response)
	if err != nil {
		return nil, err
	}

	//DBの操作が必要な場合はrepo層を実装
	// if err := uc.repo.Create(chat); err != nil {
	// 	return nil, err
	// }

	return &ChatDTO{
		ID:        chat.ID,
		Response:  chat.Response,
		CreatedBy: chat.CreatedBy,
		CreatedAt: chat.CreatedAt,
		UpdatedAt: chat.UpdatedAt,
	}, nil
}
