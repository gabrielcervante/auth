package main

import (
	"net/http"
	"regexp"

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

	if newUser.Email == "" || newUser.Password == "" {
		c.IndentedJSON(http.StatusOK, gin.H{
			"Error": "Sorry, no credentials provided",
		})
		return
	}

	if !isEmailValid(newUser.Email) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"Error": "Email format not valid",
		})
		return
	}

	var user createUser
	db.Raw("SELECT email FROM users WHERE email = ?", newUser.Email).Scan(&user.Email)
	if user.Email != "" {
		c.IndentedJSON(http.StatusOK, gin.H{
			"Error": "Sorry, email already exists",
		})
		return
	}

	go signUp(newUser.Email, newUser.Password)

	c.IndentedJSON(http.StatusOK, gin.H{
		"Success": "Account created",
	})

}

func isEmailValid(email string) bool {

	if len(email) < 3 && len(email) > 254 {
		return false
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	return emailRegex.MatchString(email)

}
