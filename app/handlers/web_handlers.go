package handlers

import (
	"linkify/db"
	"linkify/models"
	"linkify/utils"
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

	userAgent := c.GetHeader("User-Agent")
	db.RecordClick(hash, userAgent)
	c.Redirect(http.StatusFound, url)
}

func getLinkStats(c *gin.Context) {
	hash := c.Param("hash")
	click_count := db.GetClicksByHash(hash)
	url := db.GetUrlByHash(hash)

	c.JSON(200, gin.H{"click_count": click_count, "url": url})
}
