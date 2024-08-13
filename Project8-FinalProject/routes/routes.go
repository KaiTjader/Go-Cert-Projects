package routes

import (
	"example.com/project/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticatedGroup := server.Group("/")
	authenticatedGroup.Use()
	authenticatedGroup.POST("/events", createEvent)
	authenticatedGroup.PUT("/events:id", updateEvent)
	authenticatedGroup.DELETE("/events:id", deleteEvent)
	authenticatedGroup.POST("/events/:id/register", registerForEvent)
	authenticatedGroup.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
