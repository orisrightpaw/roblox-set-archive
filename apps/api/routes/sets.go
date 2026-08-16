package routes

import (
	"errors"
	"fmt"
	"log"
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

type LuaWebServiceQuery struct {
	SetID   int32  `query:"sid" validate:"excluded_with=Type,required_without=Type"`
	MaxSets int    `query:"nsets" validate:"required_if=Type user"`
	Type    string `query:"type" validate:"omitempty,oneof=user base"`
	UserID  int64  `query:"userid" validate:"required_if=Type user"`
}

type List struct {
	Values []Value `xml:"Value"`
}

type Value struct {
	Table Table `xml:"Table"`
}

type Table struct {
	Entries []Entry `xml:"Entry"`
}

type Entry struct {
	Key   string `xml:"Key"`
	Value string `xml:"Value"`
}

func Set(name string, categoryId int64, description string, assetSetId int64, creatorName string, imageAssetId int64, setType string) *Value {
	return &Value{
		Table: Table{
			Entries: []Entry{
				{
					Key:   "Name",
					Value: name,
				},
				{
					Key:   "CategoryId",
					Value: fmt.Sprint(categoryId),
				},
				{
					Key:   "Description",
					Value: description,
				},
				{
					Key:   "AssetSetId",
					Value: fmt.Sprint(assetSetId),
				},
				{
					Key:   "CreatorName",
					Value: creatorName,
				},
				{
					Key:   "ImageAssetId",
					Value: fmt.Sprint(imageAssetId),
				},
				{
					Key:   "SetType",
					Value: setType,
				},
			},
		},
	}
}

func Asset(name string, assetId int64, assetSetId int64, assetVersionId int64, creatorName string) *Value {
	return &Value{
		Table: Table{
			Entries: []Entry{
				{
					Key:   "Name",
					Value: name,
				},
				{
					Key:   "AssetId",
					Value: fmt.Sprint(assetId),
				},
				{
					Key:   "AssetSetId",
					Value: fmt.Sprint(assetSetId),
				},
				{
					Key:   "AssetVersionId",
					Value: fmt.Sprint(assetVersionId),
				},
				{
					Key:   "CreatorName",
					Value: creatorName,
				},
				{
					Key:   "IsTrusted",
					Value: "true",
				},
			},
		},
	}
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

	sets.Get("/luawebservice", func(c fiber.Ctx) error {
		lwsQuery := new(LuaWebServiceQuery)
		if err := c.Bind().Query(lwsQuery); err != nil {
			log.Println(err.Error())
			return fiber.ErrBadRequest
		}

		list := new(List)

		switch lwsQuery.Type {
		case "user":
			if lwsQuery.UserID <= 0 || lwsQuery.MaxSets <= 0 {
				return fiber.ErrBadRequest
			}

			user, err := query.User.WithContext(c.Context()).Where(query.User.ID.Eq(lwsQuery.UserID)).First()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.XML(list)
			} else if err != nil {
				return fiber.ErrInternalServerError
			}

			baseId := user.ID * -8
			if lwsQuery.MaxSets >= 1 {
				list.Values = append(list.Values, *Set("My Models", baseId, "A set of my models.", baseId, user.UserName, 21267705, "private"))
			}
			if lwsQuery.MaxSets >= 2 {
				list.Values = append(list.Values, *Set("My Decals", baseId+1, "A set of my decals.", baseId+1, user.UserName, 21002577, "private"))
			}
			if lwsQuery.MaxSets <= 2 {
				return c.XML(list)
			}

			subscriptions, err := query.Subscriber.WithContext(c.Context()).Where(query.Subscriber.UserID.Eq(user.ID)).Order(query.Subscriber.AssetSetID.Asc()).Limit(lwsQuery.MaxSets - 2).Find()
			if err != nil {
				return fiber.ErrInternalServerError
			}

			subscribedSets := make([]int64, len(subscriptions))
			for i := range len(subscriptions) {
				subscribedSets[i] = int64(subscriptions[i].AssetSetID)
			}

			sets, err := query.AssetSet.WithContext(c.Context()).Where(query.AssetSet.ID.In(subscribedSets...)).Order(query.AssetSet.ID.Asc()).Find()
			if err != nil {
				return fiber.ErrInternalServerError
			}

			for i := range len(sets) {
				list.Values = append(list.Values, *Set(sets[i].Name, sets[i].ID, sets[i].Description, sets[i].ID, sets[i].CreatorName, sets[i].ImageAssetID, "user"))
			}

			return c.XML(list)
		case "base":
			base := []int64{2, 3, 4}

			sets, err := query.AssetSet.WithContext(c.Context()).Where(query.AssetSet.ID.In(base...)).Order(query.AssetSet.ID.Asc()).Find()
			if err != nil {
				return fiber.ErrInternalServerError
			}

			for i := range len(sets) {
				list.Values = append(list.Values, *Set(sets[i].Name, sets[i].ID, sets[i].Description, sets[i].ID, sets[i].CreatorName, sets[i].ImageAssetID, "base"))
			}

			return c.XML(list)
		default:
			if lwsQuery.SetID < 0 {
				return c.XML(list)
			}

			set, err := query.AssetSet.WithContext(c.Context()).Where(query.AssetSet.ID.Eq(int64(lwsQuery.SetID))).First()
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fiber.ErrNotFound
			} else if err != nil {
				return fiber.ErrInternalServerError
			}

			assets, err := query.Asset.WithContext(c.Context()).Where(query.Asset.AssetSetID.Eq(set.ID)).Find()
			if err != nil {
				return fiber.ErrInternalServerError
			}

			for i := range len(assets) {
				list.Values = append(list.Values, *Asset(assets[i].AssetName, assets[i].AssetID, assets[i].AssetSetID, assets[i].AssetVersionID, assets[i].CreatorName))
			}

			return c.XML(list)
		}
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
