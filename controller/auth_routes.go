package controller

import (
	"net/http"
	"regexp"

	"github.com/GatorsTigers/ConcurrentBookingSystem/database"
	"github.com/GatorsTigers/ConcurrentBookingSystem/models"
	"github.com/gin-gonic/gin"
)

// AddUser wraps the POST User method
func AddUser(context *gin.Context) {
	var user models.User
	// fmt.Println(json.NewDecoder(r.Body))
	var screenJson []models.Screen
	if err := context.BindJSON(&screenJson); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse screen response",
		})
		return
	}
	// Check if email is valid
	if !isEmailValid(user.EmailId) {
		context.JSON(http.StatusBadRequest, "Error occured while trying to add user - email is not valid")
	} else {
		// Extract company domain from email
		// Add user
		newUser, er2 := database.AddUser(&user)
		if er2 != nil {
			context.JSON(http.StatusBadRequest, er2.Error())
		} else {
			context.JSON(http.StatusOK, newUser)
		}
	}
}

// isEmailValid is for Email validation
func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}
