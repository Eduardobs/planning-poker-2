package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"liposo/planing-poker/endpoints"
	"liposo/planing-poker/handlers"
	"log"
)

func main() {
	liposoClient, err := handlers.CreateClient()
	if err != nil {
		panic(err)
	}

	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	endpoints.SetupPortalEndpoints(app, liposoClient)
	endpoints.SetupApiEndpoints(app, liposoClient)

	log.Fatal(app.Listen(":8080"))
}
