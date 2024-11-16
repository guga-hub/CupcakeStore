package services

import (
	"github.com/guga_hub/cupcakestore/models"
	"github.com/guga_hub/cupcakestore/repositories"
)

type DashboardService interface {
	GetInfo(lastNDays int) *models.Dashboard
}

type dashboardService struct {
	dashboardRepository repositories.DashboardRepository
}

func NewDashboardService(dashboardRepository repositories.DashboardRepository) DashboardService {
	return &dashboardService{
		dashboardRepository: dashboardRepository,
	}
}

func (s *dashboardService) GetInfo(lastNDays int) *models.Dashboard {
	return s.dashboardRepository.GetInfo(lastNDays)
}
