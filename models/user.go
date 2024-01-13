package models

import (
	"context"
	"errors"
	"issue-maintainer/db"
	"issue-maintainer/utils"
	"strings"
)

type User struct {
	Id       int64
	Username string
	Email    string `binding:"required"`
	Password string `binding:"required"`
	GroupId  int64
}

func (u User) Save() error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	username := strings.Split(u.Email, "@")[0]

	_, err = db.DB.Exec(context.Background(), query, username, u.Email, hashedPassword)

	return err
}

/* validate the email or password */
func (u *User) ValidateCredentials() (int64, error) {
	query := `SELECT id, password FROM users WHERE email = $1`

	row := db.DB.QueryRow(context.Background(), query, u.Email)

	var id int64
	var password string
	err := row.Scan(&id, &password)

	if err != nil {
		return 0, errors.New("User not found")
	}

	err = utils.ComparePassword(u.Password, password)

	return id, err
}
