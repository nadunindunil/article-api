package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/nadunindunil/article-api/article"
	"github.com/nadunindunil/article-api/config"
	_ "github.com/nadunindunil/article-api/docs"
	"github.com/nadunindunil/article-api/tag"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// @title          Article API
// @version        1.0
// @description    This is a sample api.
// @termsOfService http://swagger.io/terms/

// @contact.name  Nadun Indunil
// @contact.url   https://www.linkedin.com/in/nadunindunil/
// @contact.email nadun1indunil@gmail.com

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:3000
// @BasePath /
func main() {
	c, error := config.LoadConfig()
	if error != nil {
		log.Fatalln("Failed at config", error)
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DBUser, c.DBPass, c.DBHost, c.DBPort, c.DBName)
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Couldn't connect to the DB!")
	}

	app := fiber.New()

	article.Init(DB, app)
	tag.Init(DB, app)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Listen(c.Port)
}
