package routes

import (
	"starterkit-go-auth/controllers"
	"starterkit-go-auth/middleware"

	"github.com/gin-gonic/gin"
)

func Route()  {
	r := gin.Default()

	////// AUTH \\\\\\

	r.POST("/register", controllers.Register) // Route Register
	r.POST("/login", controllers.Login) // Route Login
	r.GET("/user", middleware.MiddlewareAuth, controllers.User) // Route User

	r.Run()
}