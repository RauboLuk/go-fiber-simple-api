package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type Task struct {
	Name  string
}

type ViewData struct {
	Title  string
	Tasks []Task
}

func main() {
	engine := html.New("./views", ".html")
	tasks := []Task{{Name: "Example task"}}

	viewData := ViewData{
		Title: "Hello, World!",
		Tasks: tasks,
	}


  app := fiber.New(fiber.Config{
		Views: engine,
  })

  app.Get("/", func(c *fiber.Ctx) error {
    return c.Render("index", viewData)
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
