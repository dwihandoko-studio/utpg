package routes

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/dwihandoko-studio/utpg/situgu/controllers"
	"github.com/dwihandoko-studio/utpg/situgu/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var (
	RedisHost = os.Getenv("REDISHOST")
	RedisPort = os.Getenv("REDISPORT")
)

func Init() *gin.Engine {
	gob.Register(models.User{})
	funcMap := template.FuncMap{
		"printTimestamp": PrintTimestamp,
		// "wrapRelayNetwork": WrapRelayNetwork,
	}
	// fmt.Println("using database: ", servercfg.GetDB())
	// if err := database.InitializeDatabase(); err != nil {
	// 	log.Fatal("Error connecting to Database:\n", err)
	// }
	router := gin.Default()
	store, _ := redis.NewStore(10, "tcp", RedisHost+":"+RedisPort, "", []byte("secret"))

	router.Use(sessions.Sessions("lyndisdik", store))
	//router.LoadHTMLGlob("html/*")
	templates := template.Must(template.New("").Funcs(funcMap).ParseGlob("html/*"))
	router.SetHTMLTemplate(templates)
	// router.Static("theme", "./assets")
	// router.Static("favicon", "./favicon")
	// router.POST("/authenticated", controllers.RequestLogin)
	router.GET("/situgu", controllers.DashboardInit)
	router.GET("/situgu/s", controllers.DashboardSu)
	router.GET("/situgu/a", controllers.DashboardAdmin)
	// router.POST("/login", ProcessLogin)
	//use  authorization middleware
	private := router.Group("/situgu", AuthRequired)
	{
		//router.Use(AuthRequired)
		private.GET("/situgu", controllers.DashboardInit)
		private.GET("/situgu/s", controllers.DashboardSu)
		private.GET("/situgu/a", controllers.DashboardAdmin)
		// //network handlers
		// private.POST("/create_network", CreateNetwork)
		// private.POST("/delete_network", DeleteNetwork)
		// private.POST("/edit_network", EditNetwork)
		// private.POST("/update_network", UpdateNetwork)
		// private.GET("/refreshkeys/:net", RefreshKeys)
		// //key handlers
		// private.POST("/create_key", NewKey)
		// private.POST("/delete_key", DeleteKey)
		// //user handlers
		// private.POST("/create_user", CreateUser)
		// private.POST("/delete_user", DeleteUser)
		// private.GET("/edit_user", EditUser)
		// private.POST("/update_user/:user", UpdateUser)
		// //node handlers
		// private.POST("/edit_node", EditNode)
		// private.POST("/delete_node", DeleteNode)
		// private.POST("/update_node/:net/:mac", UpdateNode)
		// private.GET("/node_health", NodeHealth)
		// //gateway handlers
		// private.POST("/create_egress/:net/:mac", CreateEgress)
		// private.POST("/process_egress/:net/:mac", ProcessEgress)
		// private.POST("/delete_egress/:net/:mac", DeleteEgress)
		// private.POST("/create_ingress/:net/:mac", CreateIngress)
		// private.POST("/delete_ingress/:net/:mac", DeleteIngress)
		// private.POST("/create_relay/:net/:mac", CreateRelay)
		// private.POST("/delete_relay/:net/:mac", DeleteRelay)
		// private.POST("/process_relay/:net/:mac", ProcessRelayCreation)
		// //ext client handlers
		// private.POST("/create_ingress_client/:net/:mac", CreateIngressClient)
		// private.POST("/delete_ingress_client/:net/:id", DeleteIngressClient)
		// private.POST("/edit_ingress_client/:net/:id", EditIngressClient)
		// private.POST("/get_qr/:net/:id", GetQR)
		// private.POST("/get_client_config/:net/:id", GetClientConfig)
		// private.POST("/update_client/:net/:id", UpdateClient)
		// //dns handlers
		// private.POST("/create_dns", CreateDNS)
		// private.POST("/delete_dns/:net/:name/:address", DeleteDNS)
		// //logout
		// private.GET("/logout", LogOut)
	}
	// files := router.Group("/file", FileAuth)
	// {
	// 	files.StaticFS("", http.Dir("file"))
	// 	files.POST(":file", FileUpload)
	// }
	return router
}

func PrintTimestamp(t int64) string {
	time := time.Unix(t, 0)
	return time.String()
}

func AuthRequired(c *gin.Context) {
	// tokenDeskrip, err := helpers.GetCookie(c, "data")
	// if err != nil {
	// 	fmt.Println("loggedIn status: ", false)
	// 	c.HTML(http.StatusUnauthorized, "Login", gin.H{"messge": err.Error()})
	// 	c.Abort()
	// }
	// fmt.Println(tokenDeskrip["data"].(string))
	// fmt.Println("authorized - good to go")
	// c.Next()

	session := sessions.Default(c)

	loggedIn := session.Get("loggedIn")
	fmt.Println("loggedIn status: ", loggedIn)
	if loggedIn != true {
		// adminExists, err := controller.HasAdmin()
		// fmt.Println("response from HasAdmin(): ", adminExists, err)
		// if err != nil {
		// 	fmt.Println("error checking for admin")
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 	c.Abort()
		// }
		// if !adminExists {
		// 	fmt.Println("no admin")
		// 	c.HTML(http.StatusOK, "new", nil)
		// 	c.Abort()
		// } else {
		message := session.Get("error")
		fmt.Println("user exists --- message\n", message)
		c.HTML(http.StatusUnauthorized, "Login", gin.H{"messge": message})
		c.Abort()
		// }
	}
	fmt.Println("authorized - good to go")
}
