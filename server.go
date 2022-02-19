package main

import "github.com/gofiber/fiber/v2"

func main() {
  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })

	app.Get("/:value", func(c *fiber.Ctx) error {
		test := "test"
		return c.SendString(test + " " + c.Params("value"))
	})

  app.Listen(":3000")
}
