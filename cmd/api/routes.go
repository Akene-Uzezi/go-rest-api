package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	g := gin.Default()
	g.RedirectTrailingSlash = true

	g.GET("/ping", func (c *gin.Context) {
		c.JSON(200, gin.H{"message": "PONG"})
	})

	v1 := g.Group("/api/v1")
	{
		v1.GET("/", app.getHome)
		v1.POST("/events", app.createEvent)
		v1.GET("/events", app.getAllEvents)
		v1.GET("/events/:id", app.getEvent)
		v1.PUT("/events/:id", app.updateEvent)
		v1.DELETE("/events/:id", app.deleteEvent)

		v1.POST("/auth/register", app.registerUser)
		v1.DELETE("/delete/:id", app.deleteUser)
		v1.DELETE("/events", app.deleteAllEvents)
	}
	return  g
}