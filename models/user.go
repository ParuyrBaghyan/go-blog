package models

import (
	"errors"
	"go-blog/db"
	"go-blog/utils"
)

type SignupUser struct {
	Id       int64
	UserName string `binding:"required"`
	Email    string `binding:"required"`
	Password string `binding:"required,min=8"`
}

type SigninUser struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required,min=8"`
}

type UserProfile struct {
	Id       int64
	UserName string `binding:"required"`
	Email    string `binding:"required"`
}

func (u *SignupUser) Save() error {
	query := `INSERT INTO users(userName , email , password) VALUES(?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return errors.New("could not hash password ❌")
	}

	_, err = stmt.Exec(u.UserName, u.Email, hashedPassword)
	if err != nil {
		return err
	}

	return err
}

func (u *SigninUser) ValidateCredentials() error {
	query := "SELECT id , password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.Id, &retrivedPassword)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrivedPassword)
	if !passwordIsValid {
		return errors.New("credentials invalid! ❌")
	}

	return err
}

func GetUserById(userId int64) (*SignupUser, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.DB.QueryRow(query, userId)

	var user SignupUser
	err := row.Scan(&user.Id, &user.UserName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
