package contract

import (
	"github.com/Slintox/grpcAPI/pkg/models"
)

type Calculator interface {
	Add(request *models.AddRequest) (*models.AddResponse, error)
}
