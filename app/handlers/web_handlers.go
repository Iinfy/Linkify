package handlers

import (
	"linkify/db"
	"linkify/models"
	"linkify/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func addLink(c *gin.Context) {
	var link models.AddLink
	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(400, gin.H{"error": "internal server error"})
		return
	}
	if len(link.Url) > 2048 {
		c.JSON(422, gin.H{"error": "URL too long"})
		return
	}
	if !strings.HasPrefix(link.Url, "http://") && !strings.HasPrefix(link.Url, "https://") {
		c.JSON(422, gin.H{"error": "only http/https allowed"})
		return
	}
	urlHash := utils.GetURLHash(link.Url)
	db.AddUrl(link.Url, urlHash)
	c.JSON(200, gin.H{"url": urlHash})
}

func goToLink(c *gin.Context) {
	hash := c.Param("hash")
	url := db.GetUrlByHash(hash)
	if url == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

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

func createQR(c *gin.Context) {
	url := c.Query("text")
	if url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing text"})
		return
	}
	png, err := utils.GenerateQR(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.Header("Content-Type", "image/png")
	c.Header("Content-Length", strconv.Itoa(len(png)))
	c.Data(http.StatusOK, "image/png", png)
}
