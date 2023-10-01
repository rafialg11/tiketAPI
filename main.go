package main

import (
	// "log"
	"os"

	"tiketAPI/configs"
	"tiketAPI/routes"

	// "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	configs.InitDB()

	// Call eventRoutes and use that here
	e := echo.New()
	routes.Routes(e)

	PORT := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(PORT))
}
