package main

import (
	"fmt"
	"log"

	"github.com/d3215k/go-embed-reactjs/web"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func main() {
    app := fiber.New()

    app.Use(cors.New())

    app.Use(cors.New(cors.Config{
        AllowOrigins: "http://localhost:3000, http://localhost:5173",
        AllowHeaders:  "Origin, Content-Type, Accept",
    }))

    app.Use(func(c *fiber.Ctx) error {
        fmt.Printf("Request: %s %s\n", c.Method(), c.Path())
        return c.Next()
	})

    app.Get("/hello", handleHello)

    app.Use("/", filesystem.New(filesystem.Config{
        Root: web.BuildHTTPFS(),
        NotFoundFile: "index.html",
    }))

    log.Fatal(app.Listen(":3000"))
}

func handleHello(c *fiber.Ctx) error {
    return c.JSON(fiber.Map{"message": "hello from the fiber server"})
}