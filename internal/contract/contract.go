package contract

import "github.com/Slintox/grpcAPI/internal/models"

type Calculator interface {
	Add(request *models.AddRequest) (*models.AddResponse, error)
}
