package client

import (
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc"
	pb "grpcAPI/api"
	"grpcAPI/internal/models"
)

type CalculatorClient struct {
	client pb.CalculatorClient
}

func NewCalculatorClient(conn *grpc.ClientConn) *CalculatorClient {
	return &CalculatorClient{
		client: pb.NewCalculatorClient(conn),
	}
}

func (c *CalculatorClient) Add(ctx context.Context, req *models.AddRequest) (*models.AddResponse, error) {
	pbReq := &pb.CalculatorRequest{
		Parameter1: fmt.Sprintf("%f", req.Parameter1),
		Parameter2: fmt.Sprintf("%f", req.Parameter2),
	}

	res, err := c.client.Add(ctx, pbReq)
	if err != nil {
		return nil, err
	}

	resF, err := strconv.ParseFloat(res.Result, 64)
	if err != nil {
		return nil, err
	}

	return &models.AddResponse{
		Result: resF,
	}, nil
}
