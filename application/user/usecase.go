package user

import (
	"github.com/DDD-Zenn/api/domain/model"
	"github.com/DDD-Zenn/api/domain/repoIF"
	"github.com/DDD-Zenn/api/external/serviceIF"
)

type UserUsecase struct {
	repository repoIF.UserRepoIF
	xSvc       serviceIF.X
}

func NewUserUsecase(userRepo repoIF.UserRepoIF, x serviceIF.X) UserIF {
	return &UserUsecase{
		repository: userRepo,
		xSvc:       x,
	}
}

func (uc *UserUsecase) Create(uid, name string) error {
	user := model.NewUser(uid, name)
	return uc.repository.Create(user)
}

func (uc *UserUsecase) FindByUID(uid string) (*UserDTO, error) {
	user, err := uc.repository.FindByUID(uid)
	if err != nil {
		return nil, err
	}

	return &UserDTO{
		UID:  user.UID,
		Name: user.Name,
	}, nil
}

func (uc *UserUsecase) Update(uid, name string) error {
	user := model.NewUser(uid, name)
	return uc.repository.Update(user)
}

func (uc *UserUsecase) Delete(uid string) error {
	return uc.repository.Delete(uid)
}

func (uc *UserUsecase) GetPost() string {
	post := uc.xSvc.GetMyPost()
	return post
}
