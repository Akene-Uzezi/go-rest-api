package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (app *application) routes() http.Handler {
	g := gin.Default()
	g.RedirectTrailingSlash = true

	g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "PONG"})
	})

	v1 := g.Group("/api/v1")
	{
		v1.GET("/", app.getHome)
		// event routes
		v1.GET("/events", app.getAllEvents)
		v1.GET("/events/:id", app.getEvent)
		v1.GET("events/:id/attendees", app.getAttendeesForEvent)
		v1.DELETE("/events", app.deleteAllEvents)

		//attendee routes
		v1.GET("/attendees/:id/events", app.getEventsByAttendee)

		// auth routes
		v1.POST("/auth/register", app.registerUser)
		v1.POST("/auth/login", app.login)

		//user routes
		v1.DELETE("/delete/:id", app.deleteUser)
		v1.GET("/users", app.getUsers)
	}

	authGroup := v1.Group("/")
	authGroup.Use(app.AuthMiddleware())
	{
		authGroup.POST("/events", app.createEvent)
		authGroup.PUT("/events/:id", app.updateEvent)
		authGroup.DELETE("/events/:id", app.deleteEvent)
		authGroup.POST("/events/:id/attendees/:userId", app.addAttendeeToEvent)
		authGroup.DELETE("/events/:id/attendees/:userId", app.deleteAttendeeFromEvent)
	}

	g.GET("/swagger/*any", func(c *gin.Context) {
		if c.Request.RequestURI == "/swagger/" {
			c.Redirect(301, "/swagger/index.html") 
		}
		ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8080/swagger/doc.json"))(c)
	})

	return g
}
