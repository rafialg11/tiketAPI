package main

import (
	"log"
	"net/http"
	"os"

	"tiketAPI/configs"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	configs.InitDB()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	PORT := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(PORT))
}
