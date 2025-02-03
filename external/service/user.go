package service

import (
	"github.com/DDD-Zenn/api/domain/model"
	"github.com/DDD-Zenn/api/domain/repoIF"
	"github.com/DDD-Zenn/api/external/serviceIF"
)

type UserService struct {
	repository repoIF.UserRepoIF
}

func NewUserService(userRepo repoIF.UserRepoIF) serviceIF.User {
	return &UserService{
		repository: userRepo,
	}
}

func (service *UserService) Create(uid, name string) error {
	user := model.NewUser(uid, name)
	return service.repository.Create(user)
}

func (service *UserService) FindByUID(uid string) (model.User, error) {
	return service.repository.FindByUID(uid)
}

func (service *UserService) Update(uid, name string) error {
	user := model.NewUser(uid, name)
	return service.repository.Update(user)
}

func (service *UserService) Delete(uid string) error {
	return service.repository.Delete(uid)
}
