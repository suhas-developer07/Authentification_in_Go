package userdb

import (
	"context"

	"github.com/suhas-developer07/Authentification_in_Go/db"
	"github.com/suhas-developer07/Authentification_in_Go/models"
)

type User struct{}

var UserRepository = User{}

func (u User) CreateUserQuery(Payload models.SignupPayload) error {
	query := `INSERT into users (username,email,password) VALUES($1, $2, $3)`

	_, err := db.DB.Exec(context.Background(), query, Payload.Username, Payload.Email, Payload.Password)

	if err != nil {
		return err
	}
	return nil
}

func (u User) CheckUserExist(Payload models.SigninPayload) error {

	return nil
}
