package main

import (
	"log"

	"github.com/VladRomanciuc/CSProject/auth/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)


func main(){

	app := fiber.New()
	app.Use(logger.New())
	router.Routes(app)
	log.Fatal(app.Listen(":8080"))

}

