package article

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nadunindunil/article-api/tag"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Article struct {
	ID        int            `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	Date      datatypes.Date `json:"date"`
	Title     string         `json:"title"`
	Body      string         `json:"body"`
	Tags      []tag.Tag      `json:"tags" gorm:"many2many:article_tag_association;"`
}

type TagDto struct {
	ID int `json:"id"`
}

type ArticleCreateDto struct {
	Title string   `json:"title" example:"Article Title"`
	Body  string   `json:"body" example:"Article Body"`
	Tags  []TagDto `json:"tags"`
	Date  string   `json:"date" example:"2022-04-23T00:00:00Z"`
}

type ArticleResponseDto struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Date      time.Time `json:"date"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Tags      []tag.Tag `json:"tags"`
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
// @Tags articles
// @Produce json
// @Success 200 {object} ArticleResponseDto
// @Router /articles [get]
func getAll(c *fiber.Ctx) error {
	var articles []Article
	DB.Preload("Tags").Find(&articles)
	return c.JSON(&articles)
}

// @Summary add a new article
// @ID create-article
// @Tags articles
// @Produce json
// @Param data body ArticleCreateDto true "article data"
// @Success 201 {object} ArticleResponseDto
// @Failure 400 {string} message
// @Failure 500 {string} message
// @Router /articles [post]
func create(c *fiber.Ctx) error {
	article := new(Article)

	if err := c.BodyParser(&article); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	if err := DB.Save(&article); err.Error != nil {
		return c.Status(500).SendString(err.Error.Error())
	}
	return c.SendStatus(201)
}

// @Summary get a article
// @ID get-article
// @Tags articles
// @Produce json
// @Param id path string true "article ID"
// @Success 200 {object} ArticleResponseDto
// @Failure 404 {object} string
// @Router /articles/{id} [get]
func getOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var article Article
	DB.Preload("Tags").Find(&article, id)
	return c.JSON(&article)
}
