package main

import (
	"example.com/restapi/db"
	"example.com/restapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
	// password := "test"
	// hashed, _ := utils.HashingPassword(password)
	// fmt.Println("Hashed:", hashed)

	// match := utils.CompareHashPassword(password, hashed)
	// fmt.Println("Match:", match)
}
