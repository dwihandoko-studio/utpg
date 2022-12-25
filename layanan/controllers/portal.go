package controllers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/dwihandoko-studio/utpg/layanan/config"
	"github.com/dwihandoko-studio/utpg/layanan/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PortalPage(c *gin.Context) {
	session := sessions.Default(c)

	loggedIn := session.Get("loggedIn")
	// log.Println("loggedIn status: ", loggedIn)

	// tokenDeskrip, err := helpers.GetCookie(c, "data")
	// if err != nil {
	// 	location := url.URL{Path: "/auth"}
	// 	c.Redirect(http.StatusFound, location.RequestURI())
	// } else {
	// 	log.Println("ISI TOKEN: ", tokenDeskrip)
	if loggedIn != true {
		location := url.URL{Path: "/auth"}
		c.Redirect(http.StatusFound, location.RequestURI())
		// } else {
		// _, err := helpers.GetInfoJwt(*tokenDeskrip)
		// if err != nil {
		// 	log.Println("Error get info token", err)
		// 	location := url.URL{Path: "/auth"}
		// 	c.Redirect(http.StatusFound, location.RequestURI())
	} else {
		token := session.Get("jwtAct")
		if token == nil || token == "" {
			log.Println("TIDAK ADA SESSION USER ", token)
			session.Clear()
			session.Save()
			location := url.URL{Path: "/auth"}
			c.Redirect(http.StatusFound, location.RequestURI())
			return
		}
		// log.Println("TOKEN JWT USER ", token)
		user, errs := models.GetUser(token.(string))
		if errs != nil {
			log.Println("Error mengambil User: ", errs.Error())
			session.Clear()
			session.Save()
			location := url.URL{Path: "/auth"}
			c.Redirect(http.StatusFound, location.RequestURI())
			return
		}
		// log.Println("USER LOGIN: ", user.Fullname)
		session.Set("user", user)
		session.Save()
		c.HTML(http.StatusOK,
			"Portal_loged",
			gin.H{
				"message":  "ERROR",
				"base_url": config.BaseUrl(c),
				"User":     user,
			})
	}
	// }
}

func TestPage(c *gin.Context) {
	session := sessions.Default(c)

	loggedIn := session.Get("loggedIn")
	if loggedIn != true {
		location := url.URL{Path: "/auth"}
		c.Redirect(http.StatusFound, location.RequestURI())
	} else {
		user := session.Get("user")
		if user == nil || user == "" {
			log.Println("TIDAK ADA SESSION USER ", user)
			session.Clear()
			session.Save()
			location := url.URL{Path: "/auth"}
			c.Redirect(http.StatusFound, location.RequestURI())
			return
		}
		log.Println(user.(models.User).Fullname)
		c.HTML(http.StatusOK,
			"Portal_loged",
			gin.H{
				"message":  "ERROR",
				"base_url": config.BaseUrl(c),
				"User":     user,
			})
	}
	// }
}
