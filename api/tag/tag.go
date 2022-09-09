package tag

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Tag struct {
	ID        int       `json:"id" gorm:"primaryKey" `
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name" gorm:"not null;unique" `
}

type TagCreateDto struct {
	Name string `json:"name"`
}

type TagByNameDateResponseDto struct {
	Tag         string   `json:"tag"`
	Count       int      `json:"count"`
	Articles    []int    `json:"articles"`
	RelatedTags []string `json:"related_tags"`
}

type TagsByNameDateDao struct {
	ArticleID int
	TagID     int
	Name      string
	FullCount int
}

func Init(gormdb *gorm.DB, app *fiber.App) {
	DB = gormdb

	DB.AutoMigrate(&Tag{})

	app.Get("/tags", getAll)
	app.Post("/tags", create)
	app.Get("/tags/:name/:date", getTagsByNameAndDate)
}

// @Summary get all tags
// @ID get-all-tags
// @Tags tags
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
// @Tags tags
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

// @Summary get tags by name and date
// @ID get-tags-by-name-and-date
// @Tags tags
// @Produce json
// @Param name path string true "tag name"
// @Param date path string true "date"
// @Success 200 {object} Tag
// @Failure 500 {string} message
// @Router /tags/{name}/{date} [get]
func getTagsByNameAndDate(c *fiber.Ctx) error {
	dateString := "20060102"

	name := c.Params("name")
	articleDate := c.Params("date")

	date, error := time.Parse(dateString, articleDate)

	if error != nil {
		return c.Status(500).SendString(error.Error())
	}

	var results []TagsByNameDateDao

	// https://stackoverflow.com/questions/28888375/run-a-query-with-a-limit-offset-and-also-get-the-total-number-of-rows
	subQuery := DB.Select("ata.article_id, ata.tag_id, t.name").
		Table("article_tag_association as ata").
		Joins("inner join tags as t on ata.tag_id = t.id inner join articles a on a.id = ata.article_id").
		Where("t.name = ? and a.date = ?", name, date)

	if err := DB.Select("q.article_id, q.tag_id, q.name, count(q.article_id) OVER() as full_count").
		Table("(?) as q", subQuery).
		Limit(10).
		Find(&results); err.Error != nil {
		return c.Status(500).SendString(err.Error.Error())
	}

	var articles []int

	for _, val := range results {
		articles = append(articles, val.ArticleID)
	}

	var tags []string

	subQueryLvl3 := DB.Select("ata.article_id").
		Table("article_tag_association as ata").
		Joins("inner join tags as t on t.id=ata.tag_id inner join articles a on a.id = ata.article_id").
		Where("t.name = ? and a.date = ?", name, date)
	subQueryLvl2 := DB.Select("*").
		Table("article_tag_association").
		Where("article_id in (?)", subQueryLvl3)
	subQueryLvl1 := DB.Select("id").
		Table("tags").
		Where("name = ?", name)

	if err := DB.Distinct("t2.name").
		Table("tags as t2").
		Joins("inner join (?) as p on p.tag_id=t2.id", subQueryLvl2).
		Where("p.tag_id <> (?)", subQueryLvl1).
		Find(&tags); err.Error != nil {
		return c.Status(500).SendString(err.Error.Error())
	}

	var count int
	if count = 0; len(results) > 0 {
		count = int(results[0].FullCount)
	}

	response := TagByNameDateResponseDto{
		Tag:         name,
		Count:       count,
		Articles:    articles,
		RelatedTags: tags,
	}

	return c.JSON(&response)
}
