package controllers

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/dwihandoko-studio/utpg/layanan/models"
	"github.com/dwihandoko-studio/utpg/situgu/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func DashboardInit(c *gin.Context) {
	session := sessions.Default(c)

	loggedIn := session.Get("loggedIn")
	fmt.Println("loggedIn status: ", loggedIn)

	if loggedIn != true {
		location := url.URL{Path: "/auth"}
		c.Redirect(http.StatusFound, location.RequestURI())
	} else {

		// tokenDeskrip, err := helpers.GetCookie(c, "data")
		// if err != nil {
		// 	log.Println("Error get cookie", err)
		// 	location := url.URL{Path: "/auth"}
		// 	c.Redirect(http.StatusFound, location.RequestURI())
		// } else {
		user := session.Get("user")
		// log.Println("ISI TOKEN: ", tokenDeskrip["data"].(string))
		// infoJwt, err := helpers.GetInfoJwt(tokenDeskrip["data"].(string))
		// if err != nil {
		// 	log.Println("Error get info token", err)
		// 	location := url.URL{Path: "/auth"}
		// 	c.Redirect(http.StatusFound, location.RequestURI())
		// } else {
		if user == nil || user == "" {
			log.Println("Error get info token", user)
			location := url.URL{Path: "/auth"}
			c.Redirect(http.StatusFound, location.RequestURI())
		}

		if user.(models.User).RoleUser == 1 {
			location := url.URL{Path: "/situgu/s"}
			c.Redirect(http.StatusFound, location.RequestURI())
		} else {
			location := url.URL{Path: "/situgu/a"}
			c.Redirect(http.StatusFound, location.RequestURI())
		}
		// }
	}
}

func DashboardSu(c *gin.Context) {
	session := sessions.Default(c)

	loggedIn := session.Get("loggedIn")
	fmt.Println("loggedIn status: ", loggedIn)

	if loggedIn != true {
		location := url.URL{Path: "/auth"}
		c.Redirect(http.StatusFound, location.RequestURI())
	} else {
		// _, err := helpers.GetCookie(c, "data")
		// if err != nil {
		// 	location := url.URL{Path: "/auth"}
		// 	c.Redirect(http.StatusFound, location.RequestURI())
		// } else {
		c.HTML(http.StatusOK,
			"Dashboard_su",
			gin.H{
				"message":  "ERROR",
				"base_url": config.BaseUrl(c),
				"User": gin.H{
					"fullname": "handoko",
				},
			})
	}
}

func DashboardAdmin(c *gin.Context) {
	session := sessions.Default(c)

	loggedIn := session.Get("loggedIn")
	fmt.Println("loggedIn status: ", loggedIn)

	if loggedIn != true {
		location := url.URL{Path: "/auth"}
		c.Redirect(http.StatusFound, location.RequestURI())
	} else {
		// _, err := helpers.GetCookie(c, "data")
		// if err != nil {
		// 	location := url.URL{Path: "/auth"}
		// 	c.Redirect(http.StatusFound, location.RequestURI())
		// } else {
		c.HTML(http.StatusOK,
			"Dashboard_admin",
			gin.H{
				"message":  "ERROR",
				"base_url": config.BaseUrl(c),
				"User": gin.H{
					"fullname": "handoko",
				},
			})
	}
}
