package handlers

import (
	"PastebinBack/app/db"
	"PastebinBack/app/models"
	"PastebinBack/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addLink(c *gin.Context) {
	var link models.AddLink
	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	urlHash := utils.GetURLHash(link.Url)
	db.AddUrl(link.Url, urlHash)
	c.JSON(200, gin.H{"url": urlHash})
}

func goToLink(c *gin.Context) {
	hash := c.Param("hash")
	url := db.GetUrlByHash(hash)

	c.Redirect(http.StatusFound, url)
}
