package main

import (
	"log"
	"orisrightpaw/roblox-set-archive/internal/dal/query"
	"orisrightpaw/roblox-set-archive/routes"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type structValidator struct {
	validate *validator.Validate
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is required")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	query.SetDefault(db)

	app := fiber.New(fiber.Config{
		StructValidator: &structValidator{validate: validator.New()},
	})

	app.Use(logger.New())

	routes.CreateAssetsRoutes(app)
	routes.CreateSetsRoutes(app)
	routes.CreateUserRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
