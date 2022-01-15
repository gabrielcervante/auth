package main

import (
	"log"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string
	Password string
	Id       int
}

func signUp(email string, password string) {

	var users User

	userId := rand.Intn(1000)

	db.Raw("SELECT id FROM users WHERE id = ?", userId).Scan(&users.Id)
	for {
		if users.Id != 0 {
			userId = rand.Intn(1000)
		} else {
			break
		}
	}

	hashPswd, passwordErr := hashPassword(password)

	if passwordErr != nil {
		log.Fatal(passwordErr)
	}

	db.Create(&User{
		Email:    email,
		Password: hashPswd,
		Id:       userId,
	})

}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
