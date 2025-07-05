package userdb

import (
	"context"
	"database/sql"
	"fmt"

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
type UserData struct {
	ID       string 
	Email    string
	Password string
}

func (u User) CheckUserExist(Payload models.SigninPayload) (*UserData,error) {

	query := `SELECT id, email, password FROM users WHERE email=$1`

	var user UserData

	err := db.DB.QueryRow(context.Background(),query,Payload.Email).Scan(&user.ID,user.Email,user.Password);

	if err != nil {
		if err == sql.ErrNoRows{
			return nil,fmt.Errorf("user not found")
		}
      return  nil,err
	}
	return &user,nil
}
