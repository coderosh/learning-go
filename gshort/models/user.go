package models

import (
	"errors"
	"gshort/db"
	"gshort/utils"
)

type User struct {
	ID       int64
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES(?, ?)`

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	id, err := db.PrepareAndExec(query, u.Email, hashedPassword)
	if err != nil {
		return err
	}

	u.ID = id

	return nil
}

func (u *User) VerifyEmailAndPassword() error {
	query := `SELECT id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)

	var hashedPassword string
	err := row.Scan(&u.ID, &hashedPassword)

	if err != nil {
		return errors.New("invalid credentials")
	}

	isValid := utils.ComparePasswordWithHashed(u.Password, hashedPassword)
	if !isValid {
		return errors.New("invalid credentials")
	}

	return nil
}
