package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
) 

func main() {
	engine := html.New("./views", ".html")


  app := fiber.New(fiber.Config{
		Views: engine,
  })

  app.Get("/", func(c *fiber.Ctx) error {
    return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
	  })
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

	app.Get("/*", func(c *fiber.Ctx) error {
		return fiber.NewError(400, "Not found")
	})

  app.Listen(":3000")
}
