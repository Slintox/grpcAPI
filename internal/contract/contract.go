package contract

import "grpcAPI/internal/models"

type Calculator interface {
	Add(request *models.AddRequest) (*models.AddResponse, error)
}
