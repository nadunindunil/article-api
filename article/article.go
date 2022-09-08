package article

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

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

func Init(gormdb *gorm.DB, app *fiber.App) {
	DB = gormdb

	DB.AutoMigrate(&Article{})

	app.Get("/articles", getAll)
	app.Post("/articles", create)
	app.Get("/articles/:id", getOne)
}

// @Summary get all articles
// @ID get-all-articles
// @Produce json
// @Success 200 {object} Article
// @Router /articles [get]
func getAll(c *fiber.Ctx) error {
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
func create(c *fiber.Ctx) error {
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
func getOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var article Article
	DB.Find(&article, id)
	return c.JSON(&article)
}
