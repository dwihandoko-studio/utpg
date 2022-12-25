package config

import (
	"os"

	"github.com/gin-gonic/gin"
)

func BaseUrl(c *gin.Context) string {
	scheme := "https"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + c.Request.Host + os.Getenv("PATH_URL")
	// return os.Getenv("BASE_URL")
}
