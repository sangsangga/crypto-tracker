package userRepository

import (
	"coffeshop/database"
	"coffeshop/models"
	"database/sql"
)

type UserRegistrationRequestDTO struct {
	Email                string
	Password             string
	PasswordConfirmation string
}

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

func InsertUser(request UserRegistrationRequestDTO) (models.User, error) {

	var err error
	db := database.Client
	tx, err := db.Begin()
	if err != nil {
		return models.User{}, err
	}

	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO users (id, email, password) VALUES (?, ?, ?)")

	if err != nil {
		return models.User{}, err
	}

	res, err := stmt.Exec(nil, request.Email, request.Password)

	if err != nil {
		return models.User{}, err
	}

	err = tx.Commit()

	if err != nil {
		return models.User{}, err
	}

	defer stmt.Close()

	stmt.QueryRow().Scan()

	id, err := res.LastInsertId()

	if err != nil {
		return models.User{}, err
	}

	return models.User{Id: id}, nil
}
