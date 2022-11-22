package server

import (
	"google.golang.org/grpc"
	"grpcAPI/api"
	"grpcAPI/internal/contract"
)

type GrpcServer struct {
	handler *GrpcHandler

	srv *grpc.Server
}

func NewCalculatorGrpcServer(c contract.Calculator) *GrpcServer {
	srv := grpc.NewServer()

	return &GrpcServer{
		srv:     srv,
		handler: NewCalculatorGrpcHandler(c),
	}
}

func (s *GrpcServer) Start(conn int32) {
	api.RegisterCalculatorServer(s.srv, s.handler)

	//start
}

func Stop() {

}