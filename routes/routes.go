package routes

import (
	"tiketAPI/controllers"
	"tiketAPI/middleware"

	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	// Buat middleware CORS
    	corsHandler := middleware.CORSWithConfig(middleware.CORSConfig{				
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
    	})

   	e.Use(corsHandler)
	
	// Public routes
	e.POST("/login", controllers.Login)
	e.POST("/register", controllers.Register)

	g := e.Group("/api")
	g.Use(middleware.JwtMiddleware)
	g.GET("/events", controllers.GetEvents)
	g.GET("/events/:id", controllers.GetEventByID)
	g.POST("/events", controllers.CreateEvent)
	g.PUT("/events/:id", controllers.UpdateEvent)
	g.DELETE("/events/:id", controllers.DeleteEvent)
}
