package server

import (
	"context"
	"fmt"
	"strconv"

	"grpcAPI/api"
	"grpcAPI/internal/contract"
	"grpcAPI/internal/models"
)

type GrpcHandler struct {
	api.UnimplementedCalculatorServer

	c contract.Calculator
}

func NewCalculatorGrpcHandler(c contract.Calculator) *GrpcHandler {
	return &GrpcHandler{
		c: c,
	}
}

func (h *GrpcHandler) Add(ctx context.Context, req *api.CalculatorRequest) (*api.CalculatorResponse, error) {
	var pm1, pm2 float64
	var err error
	pm1, err = strconv.ParseFloat(req.Parameter1, 64)
	if err != nil {
		return nil, err
	}

	pm2, err = strconv.ParseFloat(req.Parameter1, 64)
	if err != nil {
		return nil, err
	}

	reqM := &models.AddRequest{
		Parameter1: pm1,
		Parameter2: pm2,
	}

	var res *models.AddResponse
	if res, err = h.c.Add(reqM); err != nil {
		return nil, err
	}

	return &api.CalculatorResponse{
		Result: fmt.Sprintf("%f", res.Result),
	}, nil
}
