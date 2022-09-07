package article

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DSN = "host=localhost user=postgres password=mysecretpassword dbname=articledb port=5432 sslmode=disable"

type Article struct {
	gorm.Model
	Title string    `json:"title"`
	Date  time.Time `json:"date"`
	Body  string    `json:"body"`
}

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Couldn't connect to the DB!")
	}
	DB.AutoMigrate(&Article{})
}

func GetAll(c *fiber.Ctx) error {
	var articles []Article
	DB.Find(&articles)
	return c.JSON(&articles)
}

func Create(c *fiber.Ctx) error {
	article := new(Article)

	if err := c.BodyParser(article); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&article)
	return c.JSON(&article)
}
