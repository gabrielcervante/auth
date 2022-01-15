package main

import (
	"fmt"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Email    string
	Password string
	Id       int
}

func signUp(email string, password string) string {
	if email == "" || password == "" {
		return "Sorry, no credentials provided"
	}

	dsn := "host=localhost user=cervante password=cervantepswd dbname=auth port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	//var users []User
	//	if err := db.Where("email = ?", email).First(&users).Error; err != nil {
	//		return "Sorry, Email Already in Use"
	//	}

	var users User
	db.Raw("SELECT email FROM users WHERE email = ?", email).Scan(&users.Email)
	if users.Email != "" {
		return "Email already exists"
	}
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
		return "Password Error"
	}

	db.Create(&User{
		Email:    email,
		Password: hashPswd,
		Id:       userId,
	})

	return "Successfull, user created!"
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
