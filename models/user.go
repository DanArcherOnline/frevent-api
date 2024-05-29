package models

import (
	"errors"

	"github.com/frevent/db"
	"github.com/frevent/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	saveQuery := `
	INSERT INTO users 
	(email, password) 
	VALUES (?, ?)
	`
	preparedSaveQuery, err := db.DB.Prepare(saveQuery)
	if err != nil {
		return err
	}
	defer preparedSaveQuery.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := preparedSaveQuery.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func (u *User) ValidatePassword() error {
	emailQuery := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(emailQuery, u.Email)

	var hashedPassword string
	err := row.Scan(&u.ID, &hashedPassword)
	if err != nil {
		return err
	}

	isPasswordValid := utils.IsPasswordValid(u.Password, hashedPassword)

	if !isPasswordValid {
		return errors.New("invalid credentials")
	}

	return nil
}
