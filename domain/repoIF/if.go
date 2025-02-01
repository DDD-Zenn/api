package repoIF

import "github.com/DDD-Zenn/api/domain/model"

type (
	Chat interface {
		Create(progress model.Chat) error
	}
)
