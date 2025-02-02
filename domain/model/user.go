package model

type User struct {
    UID  string
    Name string
}

func NewUser(uid, name string) User {
	return User{
		UID:  uid,
		Name: name,
	}
}