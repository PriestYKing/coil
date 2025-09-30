package service

import (
	"coil/model"
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func LoginService(db *sql.DB, email string, password string) (*model.User, error) {
	var user model.User

	err := db.QueryRow("SELECT username,email,password_hash FROM coil.users where email=$1", email).Scan(&user.Username, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("User not found")
	}
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("Invalid Email or Password")
	}

	return &user, nil
}

func RegisterService(db *sql.DB, username string, email string, password string) (*model.User, error) {
	var user model.User

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("Invalid Email or Password")
	}

	er := db.QueryRow("INSERT INTO coil.users(username,email,password_hash) VALUES($1,$2,$3) RETURNING username,email,password_hash", username, email, hasedPassword).Scan(&user.Username, &user.Email, &user.Password)
	if er == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
