package controllers

import (
	"net/http"
	"os"
	"time"
	"url-shortener3/initializers"
	"url-shortener3/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	// Get email/password off req body
	var input struct {
		Email    string
		Password string
	}

	if c.Bind(&input) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read input",
		})
		return
	}

	// Hash the password
	bytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Hash the password",
		})
		return
	}

	// Create the user
	user := models.User{Email: input.Email, Password: string(bytes)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	//Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}

func Login(c *gin.Context) {
	//Get the email/password off req body
	var input struct {
		Email    string
		Password string
	}

	if c.Bind(&input) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read input",
		})
		return
	}
	var user models.User
	//Find the email in DB
	if err := initializers.DB.First(&user, "email = ?", input.Email).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong Email or Password",
		})
		return
	}

	//compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong Email or Password",
		})
		return
	}

	// Generate a jwt token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create the token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", true, true)
	//send it back
	c.JSON(http.StatusOK, gin.H{
		"massage": "You are successfully logged in",
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message":   "I'm logged in",
		"user_info": user,
	})
}
