package routers

import (
	"github.com/guga_hub/cupcakestore/controllers"
	"github.com/guga_hub/cupcakestore/database"
	"github.com/guga_hub/cupcakestore/middlewares"
	"github.com/guga_hub/cupcakestore/repositories"
	"github.com/guga_hub/cupcakestore/services"
	"github.com/gofiber/fiber/v2"
)

type UserRouter struct {
	userController controllers.UserController
}

func NewUserRouter() *UserRouter {
	// Initialize repositories
	userRepository := repositories.NewUserRepository(database.DB)

	// Initialize services with repositories
	userService := services.NewUserService(userRepository)

	// Initialize controllers with services
	userController := controllers.NewUserController(userService)

	return &UserRouter{
		userController: userController,
	}
}

func (r *UserRouter) InstallRouters(app *fiber.App) {
	userGroup := app.Group("/users").Use(middlewares.LoginRequired())
	userGroup.Get("/user/:id", r.userController.RenderUser)
	userGroup.Post("/user/update/:id", r.userController.Update)

	adminGroup := app.Group("/users").Use(middlewares.LoginAndStaffRequired())
	adminGroup.Get("/create", r.userController.RenderCreate)
	adminGroup.Post("/create", r.userController.Create)
	adminGroup.Get("/delete/:id", r.userController.RenderDelete)
	adminGroup.Post("/delete/:id", r.userController.Delete)
	adminGroup.Get("/", r.userController.RenderUsers)

}
