package userRepository

import (
	"coffeshop/database"
	"coffeshop/models"
	"database/sql"
)

func FindUserByEmail(email string) (models.User, error) {
	var retrievedUser models.User
	var err error
	db := database.Client

	row := db.QueryRow("SELECT * FROM USERS WHERE email=?", email)
	if err = row.Scan(&retrievedUser.Id, &retrievedUser.Email, &retrievedUser.Password); err != sql.ErrNoRows {
		return retrievedUser, err
	}

	return retrievedUser, nil

}

func InsertUser(userInput models.User) (models.User, error) {

	var err error
	db := database.Client
	tx, err := db.Begin()
	if err != nil {
		return userInput, err
	}

	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO users (id, email, password) VALUES (?, ?, ?)")
	stmt.Exec(nil, userInput.Email, userInput.Password)

	if err != nil {
		return userInput, err
	}

	err = tx.Commit()

	if err != nil {
		return userInput, err
	}

	defer stmt.Close()

	return userInput, nil
}
