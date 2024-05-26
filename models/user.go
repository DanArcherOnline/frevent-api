package models

import "github.com/frevent/db"

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
	result, err := preparedSaveQuery.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err
}
