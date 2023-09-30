package controllers

import (
	"fmt"
	"net/http"

	"tiketAPI/models"
	"tiketAPI/services"

	"github.com/labstack/echo/v4"
)

// GetEvents is a function to get all events
func GetEvents(c echo.Context) error {
	events, err := services.GetEvents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, events)
}

// GetEventByID is a function to get event by ID
func GetEventByID(c echo.Context) error {
	id := c.Param("id")
	event, err := services.GetEventByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, event)
}

// CreateEvent is a function to create event
func CreateEvent(c echo.Context) error {
	var event models.Event
	c.Bind(&event)
	createdEvent, err := services.CreateEvent(event)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]string{
		"message": fmt.Sprintf("Event with ID %d has been created", createdEvent.(models.Event).ID),
		"event":   createdEvent.(models.Event).Name,
	})
}

// UpdateEvent is a function to update event
func UpdateEvent(c echo.Context) error {
	id := c.Param("id")
	var event models.Event
	c.Bind(&event)
	updatedEvent, err := services.UpdateEvent(id, event)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedEvent)
}

// DeleteEvent is a function to delete event
func DeleteEvent(c echo.Context) error {
	id := c.Param("id")
	deletedEvent, err := services.DeleteEvent(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("Event with ID %s has been deleted", id),
		"event":   deletedEvent.(models.Event).Name,
	})
}
