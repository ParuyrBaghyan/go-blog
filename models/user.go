package models

import (
	"errors"
	"go-blog/db"
	"go-blog/utils"
)

type User struct {
	Id       int64
	UserName string 
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
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

func (u *User) ValidateCredentials() error {
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
