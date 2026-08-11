package routes

import (
	"errors"
	"orisrightpaw/roblox-set-archive/internal/dal/query"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type SetsQuery struct {
	ID int64 `json:"id" validate:"required,gt=0"`
}

type SetsSearchQuery struct {
	Keyword string `query:"keyword" json:"keyword" validate:"max=64"`
	Page    int32  `query:"page" json:"page" validate:"gt=-1"`
}

func CreateSetsRoutes(app *fiber.App) {
	sets := app.Group("/api/sets")

	sets.Get("/search", func(c fiber.Ctx) error {
		search := new(SetsSearchQuery)
		if err := c.Bind().Query(search); err != nil {
			return fiber.ErrBadRequest
		}

		q := query.AssetSet.WithContext(c.Context())

		if search.Keyword != "" {
			q = q.Where(query.AssetSet.Name.Like("%" + search.Keyword + "%"))
		}

		count, err := q.Count()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		sets, err := q.
			Order(query.AssetSet.ID.Asc()).
			Limit(24).
			Offset(int(search.Page * 24)).
			Find()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		return c.JSON(fiber.Map{
			"pages":   count / 24,
			"total":   count,
			"results": sets,
		})
	})

	sets.Get("/:id", func(c fiber.Ctx) error {
		setQuery := new(SetsQuery)
		if err := c.Bind().URI(setQuery); err != nil {
			return fiber.ErrBadRequest
		}

		set, err := query.AssetSet.WithContext(c.Context()).Where(query.AssetSet.ID.Eq(setQuery.ID)).First()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.ErrNotFound
		} else if err != nil {
			return fiber.ErrInternalServerError
		}

		return c.JSON(set)
	})
}
