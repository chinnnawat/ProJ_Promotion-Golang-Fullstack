package routes

import (
	"github.com/chinnnawat/ProJ_Promotion-Golang/backend/api/handlers"
	"github.com/chinnnawat/ProJ_Promotion-Golang/backend/infrastructure/identity"
	"github.com/chinnnawat/ProJ_Promotion-Golang/backend/models/userModels"
	"github.com/gofiber/fiber/v2"
)

func InitPublicRoutes(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Hello Chin"))
	})

	grp := app.Group("/api/user")

	identityManager := identity.NewIdentityManager()
	registerUseCase := userModels.NewRegisterUseCase(identityManager)

	grp.Post("/register", handlers.RegisterHandler(registerUseCase))
}
