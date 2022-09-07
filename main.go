package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/nadunindunil/article-api/article"
	_ "github.com/nadunindunil/article-api/docs"
)

func Routers(app *fiber.App) {
	app.Get("/articles", article.GetAll)
	app.Post("/articles", article.Create)
	app.Get("/articles/:id", article.GetOne)
}

// @title          Article API
// @version        1.0
// @description    This is a sample api.
// @termsOfService http://swagger.io/terms/

// @contact.name  Nadun Indunil
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:3000
// @BasePath /
func main() {
	article.InitialMigration()
	app := fiber.New()

	Routers(app)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Listen(":3000")
}
