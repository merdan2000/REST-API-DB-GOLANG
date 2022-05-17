package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/merdan2000/internal/model"
	"github.com/merdan2000/internal/settings"
)

type Users interface {
	CreateUser(user model.Users, password string) (int, error)
	ReadByID(userID int) (model.Users, error)
	ReadUser(email, password string) (model.Users, error)
	UpdateUser(user model.Users) (model.Users, error)
	DeleteUserByID(userID int, password string) error
}

type Repository struct {
	Users
}

func NewRepository(settings *settings.Settings) *Repository {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		settings.Host, settings.Port, settings.DbName, settings.Password,
		settings.User, settings.SSLmode)

	db, err := sqlx.Open("pgx", sqlInfo)
	if err != nil {
		return &Repository{}
	}

	return &Repository{Users: NewRepoUser(db)}
}
