package controllers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/dwihandoko-studio/utpg/layanan/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func DashboardSituguInit(c *gin.Context) {
	session := sessions.Default(c)

	loggedIn := session.Get("loggedIn")
	log.Println("loggedIn status: ", loggedIn)

	if loggedIn != true {
		location := url.URL{Path: "/auth"}
		c.Redirect(http.StatusFound, location.RequestURI())
	} else {
		user := session.Get("user")
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
	}
}

// func DashboardSituguSu(c *gin.Context) {
// 	session := sessions.Default(c)

// 	loggedIn := session.Get("loggedIn")
// 	log.Println("loggedIn status: ", loggedIn)

// 	if loggedIn != true {
// 		location := url.URL{Path: "/auth"}
// 		c.Redirect(http.StatusFound, location.RequestURI())
// 	} else {
// 		c.HTML(http.StatusOK,
// 			"Dashboard_su",
// 			gin.H{
// 				"message":  "ERROR",
// 				"base_url": config.BaseUrl(c),
// 				"User": gin.H{
// 					"fullname": "handoko",
// 				},
// 			})
// 	}
// }

// func DashboardSituguAdmin(c *gin.Context) {
// 	session := sessions.Default(c)

// 	loggedIn := session.Get("loggedIn")
// 	log.Println("loggedIn status: ", loggedIn)

// 	if loggedIn != true {
// 		location := url.URL{Path: "/auth"}
// 		c.Redirect(http.StatusFound, location.RequestURI())
// 	} else {
// 		c.HTML(http.StatusOK,
// 			"Dashboard_admin",
// 			gin.H{
// 				"message":  "ERROR",
// 				"base_url": config.BaseUrl(c),
// 				"User": gin.H{
// 					"fullname": "handoko",
// 				},
// 			})
// 	}
// }
