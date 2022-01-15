package main

import (
	"github.com/gin-gonic/gin"
)

type createUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
	r.POST("/signUp", postSignUp)
	r.Run()

}

func postSignUp(c *gin.Context) {
	var newUser createUser

	c.BindJSON(&newUser)

	go signUp(newUser.Email, newUser.Password)

}
