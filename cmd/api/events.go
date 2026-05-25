package main

import (
	"fmt"
	"net/http"
	"rest-api-in-gin/internal/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func(app *application) getHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "The api is running fine"})
}

func (app *application) createEvent(c *gin.Context) {
	var event database.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.models.Events.Insert(&event)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed To create error"})
		return
	}

	c.JSON(http.StatusCreated, event)
}


func (app *application) getAllEvents(c *gin.Context) {
	events, err := app.models.Events.GetAll()
	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed To retreive events"})
	}

	c.JSON(http.StatusOK, events)
}

func (app *application) getEvent(c *gin.Context) {
	context := c.Request.Context()
	 id, err := strconv.Atoi(c.Param("id")); 
	 if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	event, err := app.models.Events.Get(id, context)

	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event Not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed To retreive event"})
	}

	c.JSON(http.StatusOK, event)
}

func (app *application) updateEvent(c *gin.Context) {
	context := c.Request.Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	existingEvent, err := app.models.Events.Get(id, context)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
		return
	}

	if existingEvent == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	updatedEvent := &database.Event{}

	if err := c.ShouldBindJSON(updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 

	updatedEvent.Id = id

	if err := app.models.Events.Update(updatedEvent); err != nil {
		 c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
		 return
	}

	c.JSON(http.StatusOK, updatedEvent)
}

func (app *application) deleteEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Event Id"})
		return
	}

	if err := app.models.Events.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (app *application) deleteAllEvents(c *gin.Context) {
	context := c.Request.Context()
	if err := app.models.Events.DeleteAll(context); err != nil {
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting events"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}