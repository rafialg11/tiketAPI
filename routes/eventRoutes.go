package routes

import (
	"tiketAPI/controllers"

	"github.com/labstack/echo/v4"
)

func EventRoutes(e *echo.Echo) {
	e.GET("/events", controllers.GetEvents)
	e.GET("/events/:id", controllers.GetEventByID)
	e.POST("/events", controllers.CreateEvent)
	e.PUT("/events/:id", controllers.UpdateEvent)
	e.DELETE("/events/:id", controllers.DeleteEvent)
}
