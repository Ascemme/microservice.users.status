package handler

import (
	"context"
	"github.com/Ascemme/microservice.users.status/protos/ascemme/grpcProto"
)

type GRPCServer struct {
	grpcProto.StatusGrpcServer
}

func (s *GRPCServer) Status(ctx context.Context, req *grpcProto.RequestStatus) (*grpcProto.ResponseStatus, error) {
	return &grpcProto.ResponseStatus{Uid: 12}, nil
}
