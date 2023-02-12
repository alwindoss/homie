package main

import (
	"embed"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

//go:embed views
var viewFS embed.FS

func main() {
	// engine := html.New("./templates", ".html")
	engine := html.NewFileSystem(http.FS(viewFS), ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "views/layout/base",
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("views/index", fiber.Map{
			"hello": "world",
		})
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("views/about", fiber.Map{
			"Title": "About Page",
		}, "views/layout/aboutlayout")
	})
	app.Get("/ping", pingHandler)
	app.Listen(":8080")
}

func pingHandler(c *fiber.Ctx) error {
	return c.JSON(fmt.Sprintf("Hello the time is %s", time.Now().String()))
}
