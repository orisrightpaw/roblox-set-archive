package routes

import (
	"errors"
	"orisrightpaw/roblox-set-archive/internal/dal/query"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type UsersQuery struct {
	ID int64 `json:"id" validate:"required,gt=0"`
}

type UsersSearchQuery struct {
	Keyword string `query:"keyword" json:"keyword" validate:"max=64"`
	Page    int32  `query:"page" json:"page" validate:"gt=-1"`
}

func CreateUserRoutes(app *fiber.App) {
	users := app.Group("/api/users")

	users.Get("/search", func(c fiber.Ctx) error {
		search := new(UsersSearchQuery)
		if err := c.Bind().Query(search); err != nil {
			return fiber.ErrBadRequest
		}

		q := query.User.WithContext(c.Context())

		if search.Keyword != "" {
			q = q.Where(query.User.UserName.Like("%" + search.Keyword + "%"))
		}

		count, err := q.Count()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		users, err := q.
			Order(query.User.ID.Asc()).
			Limit(24).
			Offset(int(search.Page * 24)).
			Find()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		return c.JSON(fiber.Map{
			"pages":   count / 24,
			"total":   count,
			"results": users,
		})
	})

	users.Get("/:id", func(c fiber.Ctx) error {
		userQuery := new(UsersQuery)
		if err := c.Bind().URI(userQuery); err != nil {
			return fiber.ErrBadRequest
		}

		user, err := query.User.WithContext(c.Context()).Where(query.User.ID.Eq(userQuery.ID)).First()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.ErrNotFound
		} else if err != nil {
			return fiber.ErrInternalServerError
		}

		owned, err := query.AssetSet.WithContext(c.Context()).Where(query.AssetSet.CreatorID.Eq(user.ID)).Order(query.AssetSet.ID.Asc()).Find()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		subscriptions, err := query.Subscriber.WithContext(c.Context()).Where(query.Subscriber.UserID.Eq(user.ID)).Find()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		subscribedSets := make([]int64, len(subscriptions))
		for i := range len(subscriptions) {
			subscribedSets[i] = int64(subscriptions[i].AssetSetID)
		}

		subscribed, err := query.AssetSet.WithContext(c.Context()).Where(query.AssetSet.ID.In(subscribedSets...), query.AssetSet.CreatorID.Neq(user.ID)).Order(query.AssetSet.ID.Asc()).Find()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		return c.JSON(fiber.Map{
			"id":         user.ID,
			"user_name":  user.UserName,
			"owned":      owned,
			"subscribed": subscribed,
		})
	})

	users.Get("/:id/thumbnail", func(c fiber.Ctx) error {
		user := new(UsersQuery)
		if err := c.Bind().URI(user); err != nil {
			return fiber.ErrBadRequest
		}

		return c.JSON(user)
	})
}
