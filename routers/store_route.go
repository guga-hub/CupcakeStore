package routers

import (
	"github.com/guga_hub/cupcakestore/controllers"
	"github.com/guga_hub/cupcakestore/database"
	"github.com/guga_hub/cupcakestore/repositories"
	"github.com/guga_hub/cupcakestore/services"
	"github.com/gofiber/fiber/v2"
)

type StoreRouter struct {
	storeController controllers.StoreController
}

func NewStoreRouter() *StoreRouter {
	// Initialize repositories
	productRepository := repositories.NewProductRepository(database.DB)

	// Initialize services with repositories
	productService := services.NewProductService(productRepository)

	// Initialize controllers with services
	storeController := controllers.NewStoreController(productService)
	return &StoreRouter{
		storeController: storeController,
	}
}

func (r *StoreRouter) InstallRouters(app *fiber.App) {
	store := app.Group("/store")
	store.Get("/", r.storeController.RenderStore)
}
