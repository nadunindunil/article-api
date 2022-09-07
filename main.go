package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nadunindunil/article-api/article"
)

func Routers(app *fiber.App) {
	app.Get("/articles", article.GetAll)
	app.Post("/articles", article.Create)
	// app.Get("/articles/:id", article.getOne)
}

func main() {
	article.InitialMigration()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	Routers(app)

	app.Listen(":3000")
}
