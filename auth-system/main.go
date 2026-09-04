package main

import(
	 "auth-system/config"
	 "github.com/gofiber/fiber/v2"
	 "auth-system/routes"

)
func main() {

	config.ConnectDB()
	app := fiber.New()
	routes.SetupRoutes(app)


   
	
	app.Listen(":8080")
}