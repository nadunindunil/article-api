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
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"date"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
}

type ArticleCreateDto struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Couldn't connect to the DB!")
	}
	DB.AutoMigrate(&Article{})
}

// @Summary get all articles
// @ID get-all-articles
// @Produce json
// @Success 200 {object} Article
// @Router /articles [get]
func GetAll(c *fiber.Ctx) error {
	var articles []Article
	DB.Find(&articles)
	return c.JSON(&articles)
}

// @Summary add a new article
// @ID create-article
// @Produce json
// @Param data body ArticleCreateDto true "article data"
// @Success 200 {object} Article
// @Failure 500 {string} message
// @Router /articles [post]
func Create(c *fiber.Ctx) error {
	article := new(Article)

	if err := c.BodyParser(&article); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&article)
	return c.JSON(&article)
}

// @Summary get a article
// @ID get-article
// @Produce json
// @Param id path string true "article ID"
// @Success 200 {object} Article
// @Failure 404 {object} string
// @Router /articles/{id} [get]
func GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var article Article
	DB.Find(&article, id)
	return c.JSON(&article)
}
