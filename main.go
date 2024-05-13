package main

import (
	"fmt"

	"github.com/backEnGO/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDb()

}
func main() {
	r := gin.Default()

	fmt.Println("Server online")
	Routes(r)
	r.Run() // listen and serve on 0.0.0.0:8080

}
