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

func LoginPage(c *gin.Context) {
	session := sessions.Default(c)

	loggedIn := session.Get("loggedIn")
	log.Println("loggedIn status: ", loggedIn)
	// tokenDeskrip, _ := helpers.GetCookie(c, "data")
	// if tokenDeskrip != nil {
	// 	_, err := helpers.GetInfoJwt(*tokenDeskrip)
	// 	if err != nil {
	// 		c.HTML(http.StatusOK,
	// 			"Login",
	// 			gin.H{
	// 				"message":  "ERROR",
	// 				"base_url": config.BaseUrl(c),
	// 			})
	// 	} else {
	// 		location := url.URL{Path: "/portal"}
	// 		c.Redirect(http.StatusFound, location.RequestURI())
	// 	}
	// } else {
	// 	c.HTML(http.StatusOK,
	// 		"Login",
	// 		gin.H{
	// 			"message":  "ERROR",
	// 			"base_url": config.BaseUrl(c),
	// 		})
	// }

	if loggedIn == true {
		location := url.URL{Path: "/portal"}
		c.Redirect(http.StatusFound, location.RequestURI())
	} else {
		c.HTML(http.StatusOK,
			"Login",
			gin.H{
				"message":  "ERROR",
				"base_url": config.BaseUrl(c),
			})
	}

	// fmt.Println("Processing Login")
	// username := c.PostForm("username")
	// password := c.PostForm("password")
	// session := sessions.Default(c)

	// if username == "" || password == "" {
	// 	session.Set("message", "belum login")
	// 	session.Set("loggedIn", false)
	// 	c.HTML(http.StatusUnauthorized, "Login", gin.H{"message": "belum login"})
	// }

	// //don't need the jwt
	// res, err := models.GetTokenApp(username, password)

	// fmt.Println(res.AccessToken)
	// if err != nil {
	// 	fmt.Println("error verifying AuthRequest: ", err)
	// 	session.Set("message", err.Error())
	// 	session.Set("loggedIn", false)
	// 	c.HTML(http.StatusUnauthorized, "Login", gin.H{"message": err})
	// } else {
	// 	log.Printf("TOKENNYA %v\n", res.AccessToken)
	// 	if res.AccessToken == "" {
	// 		session.Set("loggedIn", true)
	// 		//init message
	// 		session.Set("message", "")
	// 		session.Options(sessions.Options{MaxAge: 28800})
	// 		session.Set("jwt", res.AccessToken)
	// 		session.Save()
	// 		location := url.URL{Path: "/portal"}
	// 		c.Redirect(http.StatusFound, location.RequestURI())
	// 	} else {
	// 		session.Set("message", "belum login")
	// 		session.Set("loggedIn", false)
	// 		c.HTML(http.StatusUnauthorized, "Login", gin.H{"message": "belum login"})
	// 	}
	// }
}

func RequestLogin(c *gin.Context) {

	var username = c.PostForm("username")
	var password = c.PostForm("password")

	res, err := models.GetTokenApp(username, password)
	if err != nil {
		config.ReturnJsonError(c, http.StatusBadRequest, err, "")
		return
	}

	// errs := helpers.SetCookie(c, "data", res.AccessToken)
	// err = helpers.SetCookie(c, "data", fmt.Sprintf("Bearer %s", res.AccessToken))
	// if errs != nil {
	// 	config.ReturnJsonError(c, http.StatusBadRequest, err, "")
	// 	return
	// } else {
	// 	config.ReturnJsonSuccess(c, 200, "Login berhasil", "portal", nil)
	// }

	// jwt, errs := helpers.GetInfoJwt(res.AccessToken)
	// if errs != nil {
	// 	config.ReturnJsonError(c, http.StatusBadRequest, errs, "")
	// 	return
	// }
	// fmt.Println("INI TOKENNYA: ", jwt)
	// config.ReturnJsonError(c, http.StatusBadRequest, errs, "")
	// 	return

	log.Println("INI TOKEN JWTNYA: ", res.AccessToken)

	session := sessions.Default(c)
	session.Set("loggedIn", true)
	// session.Set("level", jwt.Level)
	//init message
	session.Set("message", "")
	session.Set("jwtAct", res.AccessToken)
	session.Set("jwtRft", res.RefreshToken)
	session.Options(sessions.Options{MaxAge: 3600})
	session.Save()

	config.ReturnJsonSuccess(c, 200, "Login berhasil", "portal", nil)
}

func RequestLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	config.ReturnJsonSuccess(c, 200, "Logout berhasil", "auth", nil)
}
