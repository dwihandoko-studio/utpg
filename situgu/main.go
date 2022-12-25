package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dwihandoko-studio/utpg/situgu/routes"
	// "github.com/joho/godotenv"
)

func main() {
	log.Println("Starting server...")
	portApp := "8080"

	// if portApp == "" {
	// 	portApp = "8080"
	// 	err := godotenv.Load("local.env")
	// 	if err != nil {
	// 		panic("connectionStringGorm error..." + err.Error())
	// 	}
	// }

	router := routes.Init()
	// router.Run(":8080")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", portApp),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", srv.Addr)

	// Wait for Kill signal of channel
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// This blocks until a signal is passed into the quit channel
	<-quit

	//The context is used to inform the server it has 5 seconds to finish
	//the request is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//shutdown server
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}

// func WrapRelayNetwork(network string, data models.Node) map[string]interface{} {
// 	return map[string]interface{}{
// 		"NetworkToUse": network,
// 		"Data":         data,
// 	}
// }

// func AuthRequired(c *gin.Context) {
// 	session := sessions.Default(c)

// 	loggedIn := session.Get("loggedIn")
// 	fmt.Println("loggedIn status: ", loggedIn)
// 	if loggedIn != true {
// 		adminExists, err := controller.HasAdmin()
// 		fmt.Println("response from HasAdmin(): ", adminExists, err)
// 		if err != nil {
// 			fmt.Println("error checking for admin")
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			c.Abort()
// 		}
// 		if !adminExists {
// 			fmt.Println("no admin")
// 			c.HTML(http.StatusOK, "new", nil)
// 			c.Abort()
// 		} else {
// 			message := session.Get("error")
// 			fmt.Println("user exists --- message\n", message)
// 			c.HTML(http.StatusUnauthorized, "Login", gin.H{"messge": message})
// 			c.Abort()
// 		}
// 	}
// 	fmt.Println("authorized - good to go")
// }
