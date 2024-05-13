package main

import (
	"github.com/backEnGO/controllers"
	"github.com/backEnGO/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	user := route.Group("/")
	user.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "bienvenue",
		})
	})
	user.POST("/signUp", controllers.SignUp)
	user.POST("/login", controllers.Login)
	user.GET("/validate", middleware.RequireAuth, controllers.Validate)
	user.POST("/addUser", controllers.AddUser)
	user.GET("/allUsers", controllers.AllUsers)
	user.GET("/oneUser/:id", controllers.OneUser)
	user.PUT("/updateUser/:id", controllers.UpdateUser)
	user.DELETE("/deleteOneUser/:id", controllers.DeleteOneUser)

}
