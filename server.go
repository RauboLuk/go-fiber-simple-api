package main

import "fmt"
import "github.com/gofiber/fiber/v2"

func main() {
  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })

	app.Get("params/optional/:param", func(c *fiber.Ctx) error {
		test := "test"
		fmt.Println(test)

		if c.Params("param") != "" {
			return c.SendString("Param " + c.Params("param"))
		}
		return c.SendString("Param not provided " + c.Params("param"))
	})

	app.Get("params/:param?", func(c *fiber.Ctx) error {
		return c.SendString("Param " + c.Params("param"))
	})

	app.Get("/wildcard/*", func(c *fiber.Ctx) error {
		return c.SendString("Wildcard: " + c.Params("*"))
	})

  app.Listen(":3000")
}
