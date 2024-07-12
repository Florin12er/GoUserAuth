package main

import (
	"userAuth/controllers"
	"userAuth/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
    initializers.LoadEnv()
    initializers.ConnectToDb()
    initializers.SyncDatabase()
}

func main() {
    gin.SetMode(gin.ReleaseMode) // Set release mode for production
    r := gin.Default()

    // Set trusted proxies
    // Example: Trust only localhost
    if err := r.SetTrustedProxies([]string{"127.0.0.1"}); err != nil {
        panic(err)
    }

    r.POST("/signup",controllers.Signup) 
    r.POST("/login", controllers.Login)
    r.Run() // By default it serves on :8080 unless a PORT environment variable was defined.
}

