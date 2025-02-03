package serviceIF

import (
	"github.com/DDD-Zenn/api/domain/model"
)

type (
	Gemini interface {
		GenerateResponse(prompt string) (string, error)
	}

	User interface {
		FindByUID(uid string) (model.User, error)
		Create(uid, name string) error
		Update(uid, name string) error
		Delete(uid string) error
	}
)
