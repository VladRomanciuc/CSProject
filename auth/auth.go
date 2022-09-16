package main

import (
	"log"
	"github.com/VladRomanciuc/CSProject/auth/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)


func main(){

	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())
	router.Routes(app)
	log.Fatal(app.ListenTLS(":8080", "./127.0.0.1.pem", "./127.0.0.1-key.pem"))
}

