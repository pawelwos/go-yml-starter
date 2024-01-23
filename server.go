package main

import (
	"fmt"
	"pawelwos/fiber-test/lib/loaders"
	"time"

	"html/template"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/template/html/v2"
)

func main() {

	// Create a new engine
	engine := html.New("./views", ".html")

	fm := map[string]interface{}{
		// load num of posts
		"posts": loaders.Posts,
		// raw unescaped HTML
		"raw": func(s string) template.HTML {
			return template.HTML(s)
		},
		// current year
		"date": func(f string) template.HTML {
			year := time.Now().Year()
			return template.HTML(fmt.Sprint(year))
		}}
	// add above functions to engine
	engine.AddFuncMap(fm)
	// Pass the engine to the Views
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		home, err := loaders.Yaml("home", "pages")
		if err != nil {
			return c.SendStatus(404)
		}
		return c.Render("index", fiber.Map{
			"data": home,
		}, "layouts/main")
	})
	app.Get("/blog/:slug", func(c *fiber.Ctx) error {
		page, err := loaders.Yaml(c.Params("slug"), "blog")
		if err != nil {
			return c.SendStatus(404)
		}
		return c.Render("post", fiber.Map{
			"data": page,
		}, "layouts/main")
	})
	app.Get("/:slug", func(c *fiber.Ctx) error {
		page, err := loaders.Yaml(c.Params("slug"), "pages")
		if err != nil {
			return c.SendStatus(404)
		}
		return c.Render("page", fiber.Map{
			"data": page,
		}, "layouts/main")
	})
	app.Listen(":3000")
}
