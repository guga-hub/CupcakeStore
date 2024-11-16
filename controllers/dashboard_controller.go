package controllers

import (
	"github.com/guga_hub/cupcakestore/services"
	"github.com/guga_hub/cupcakestore/views"
	"github.com/gofiber/fiber/v2"
)

type DashboardController interface {
	RenderDashboard(ctx *fiber.Ctx) error
}

type dashboardController struct {
	dashboardService services.DashboardService
}

func NewDashboardController(s services.DashboardService) DashboardController {
	return &dashboardController{
		dashboardService: s,
	}
}

func (c *dashboardController) RenderDashboard(ctx *fiber.Ctx) error {
	data := c.dashboardService.GetInfo(30)
	return views.Render(ctx, "dashboard/dashboard", data, views.BaseLayout)
}
