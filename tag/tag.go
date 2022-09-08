package tag

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Tag struct {
	ID        uint      `json:"id" gorm:"primaryKey" `
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name" gorm:"unique" `
}

type TagCreateDto struct {
	Name string `json:"name"`
}

func Init(gormdb *gorm.DB, app *fiber.App) {
	DB = gormdb

	DB.AutoMigrate(&Tag{})

	app.Get("/tags", getAll)
	app.Post("/tags", create)
	// app.Get("/tags/:id", getOne)
}

// @Summary get all tags
// @ID get-all-tags
// @Produce json
// @Success 200 {object} Tag
// @Router /tags [get]
func getAll(c *fiber.Ctx) error {
	var tags []Tag
	DB.Find(&tags)
	return c.JSON(&tags)
}

// @Summary add a new tag
// @ID create-tag
// @Produce json
// @Param data body TagCreateDto true "tag data"
// @Success 200 {object} Tag
// @Failure 500 {string} message
// @Router /tags [post]
func create(c *fiber.Ctx) error {
	tag := new(Tag)

	if err := c.BodyParser(&tag); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if err := DB.Save(&tag); err.Error != nil {
		return c.Status(500).SendString(err.Error.Error())
	}
	return c.JSON(&tag)
}
