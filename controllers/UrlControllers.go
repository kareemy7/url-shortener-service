package controllers

import (
	"math/rand/v2"
	"net/http"
	"os"
	"url-shortener3/initializers"
	"url-shortener3/models"

	"github.com/gin-gonic/gin"
)

func Shorten(c *gin.Context) {
	var input struct {
		OriginalURL string
	}

	if c.Bind(&input) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read input",
		})
		return
	}

	user, _ := c.Get("user")
	userID := user.(models.User).ID

	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Shuffle(len(letters), func(i, j int) {
		letters[i], letters[j] = letters[j], letters[i]
	})

	u := make([]rune, 6)
	for i := range u {
		u[i] = letters[i]
	}
	shortCode := string(u)

	url := models.URL{UserID: userID, OriginalURL: input.OriginalURL, ShortCode: shortCode}
	urlDB := initializers.DB.Create(&url)

	if urlDB.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to add URL to DB",
		})
		return
	}

	c.Set("shortCode", shortCode)
	c.JSON(http.StatusOK, gin.H{
		"short_url": "http://localhost:" + os.Getenv("Port") + "/" + shortCode,
	})

}

func Retrieve(c *gin.Context) {
	shortCode := c.Param("shortCode")

	var url models.URL
	if err := initializers.DB.First(&url, "short_code = ?", shortCode).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "URL not found",
		})
		return
	}

	c.Redirect(http.StatusFound, url.OriginalURL)
}
