package models

import (
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
