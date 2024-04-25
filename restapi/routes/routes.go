package routes

import (
	"example.com/restapi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", GetEvents)
	server.POST("/events", middlewares.Authenticate, CreateEvents)
	server.GET("/events/:id", GetEvent)
	server.PUT("/events/:id", UpdateEvent)
	server.DELETE("/events/:id", DeleteEvent)
	server.POST("/signup", SignUp)
	server.POST("/login", login)
	server.GET("/users", GetUsers)

}
