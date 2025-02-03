package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/DDD-Zenn/api/domain/model"
	"github.com/DDD-Zenn/api/domain/repoIF"
	"github.com/DDD-Zenn/api/infrastructure/database"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(ctx context.Context) repoIF.UserRepoIF {
	return &UserRepo{DB: database.DB}
}

func (r *UserRepo) Create(user model.User) error {
	if r.DB == nil {
        return fmt.Errorf("database connection is nil")
    }
    _, err := r.DB.Exec("INSERT INTO users (uid, name) VALUES (?, ?)", user.UID, user.Name)
	return err
}

func (r *UserRepo) FindByUID(uid string) (model.User, error) {
	var user model.User
	err := r.DB.QueryRow("SELECT uid, name FROM users WHERE uid = ?", uid).Scan(&user.UID, &user.Name)
	return user, err
}

func (r *UserRepo) Update(user model.User) error {
	_, err := r.DB.Exec("UPDATE users SET name = ? WHERE uid = ?", user.Name, user.UID)
	return err
}

func (r *UserRepo) Delete(uid string) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE uid = ?", uid)
	return err
}
