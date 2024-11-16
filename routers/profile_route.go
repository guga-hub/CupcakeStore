package routers

import (
	"github.com/guga_hub/cupcakestore/controllers"
	"github.com/guga_hub/cupcakestore/database"
	"github.com/guga_hub/cupcakestore/middlewares"
	"github.com/guga_hub/cupcakestore/repositories"
	"github.com/guga_hub/cupcakestore/services"
	"github.com/gofiber/fiber/v2"
)

type ProfileRouter struct {
	profileController controllers.ProfileController
}

func NewProfileRouter() *ProfileRouter {
	// Initialize repositories
	profileRepository := repositories.NewProfileRepository(database.DB)

	// Initialize services with repositories
	profileService := services.NewProfileService(profileRepository)

	// Initialize controllers with services
	profileController := controllers.NewProfileController(profileService)

	return &ProfileRouter{
		profileController: profileController,
	}
}

func (r *ProfileRouter) InstallRouters(app *fiber.App) {
	profile := app.Group("/profile").Use(middlewares.LoginRequired())

	profile.Get("/:id", r.profileController.RenderProfile)
	profile.Post("/update/:id", r.profileController.Update)
}
