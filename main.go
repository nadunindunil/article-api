package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/nadunindunil/article-api/article"
	_ "github.com/nadunindunil/article-api/docs"
	"github.com/nadunindunil/article-api/tag"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DSN = "host=localhost user=postgres password=mysecretpassword dbname=articledb port=5432 sslmode=disable"

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
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Couldn't connect to the DB!")
	}

	app := fiber.New()

	article.Init(DB, app)
	tag.Init(DB, app)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Listen(":3000")
}
