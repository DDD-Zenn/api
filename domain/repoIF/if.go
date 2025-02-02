package repoIF

import "github.com/DDD-Zenn/api/domain/model"

type (
	Chat interface {
		Create(progress model.Chat) error
	}

	User interface {
		FindByUID(uid string) (model.User, error)
		Create(user model.User) error
		Update(user model.User) error
		Delete(uid string) error
	}
)
