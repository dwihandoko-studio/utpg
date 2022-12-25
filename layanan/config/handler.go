package config

import (
	"net/http"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ReturnJsonSuccess(c *gin.Context, status int, message, redirect string, data interface{}) {
	location := url.URL{Path: "/" + redirect}
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  status,
			"message": message,
			"url":     location.RequestURI(),
			"data":    data,
		},
	)
}

func ReturnJsonError(c *gin.Context, status int, err error, redirect string) {
	location := url.URL{Path: "/" + redirect}
	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  status,
			"message": err.Error(),
			"url":     location.RequestURI(),
		},
	)
}

func ReturnSuccess(c *gin.Context, page, message string) {
	session := sessions.Default(c)
	session.Set("messsge", message)
	session.Set("page", page)
	session.Save()
	location := url.URL{Path: "/"}
	c.Redirect(http.StatusFound, location.RequestURI())
}

func ReturnError(c *gin.Context, status int, err error, page string) {
	// var data PageData
	session := sessions.Default(c)
	session.Set("message", err.Error())
	session.Set("page", page)
	session.Save()
	// if page != "" {
	// 	data.Init(c, page, err.Error())
	// } else {
	// 	data.Init(c, "Networks", err.Error())
	// }
	c.HTML(status, "layout", err.Error())
}
