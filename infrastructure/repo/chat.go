package repo

import (
	"context"

	"github.com/DDD-Zenn/api/domain/model"
	"github.com/DDD-Zenn/api/domain/repoIF"
)

//DBとの連携・処理は以下の部分で行う

type chatRepo struct {
}

func NewChatRepo(ctx context.Context) repoIF.Chat {
	return &chatRepo{}
}

func (r *chatRepo) Create(chat model.Chat) error {
	return nil
}
