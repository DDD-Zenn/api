package user

type UserIF interface {
	FindByUID(uid string) (*UserDTO, error)
	Create(uid, name string) error
	Update(uid, name string) error
	Delete(uid string) error
	GetPost() string
}
