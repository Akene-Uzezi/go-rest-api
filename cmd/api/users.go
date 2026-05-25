package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) getUsers(c *gin.Context) {
	context := c.Request.Context()
	users, err := app.models.Users.GetAll(context)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (app *application) deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := app.models.Users.Delete(id, c.Request.Context()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}