package handler

import (
	"context"
	"encoding/json"
	"github.com/Ascemme/microservice.users.status/pkg/repository"
	"github.com/Ascemme/microservice.users.status/protos/ascemme/grpcProto"
	"log"
)

type GRPCServer struct {
	grpcProto.StatusGrpcServer
	repo *repository.Repository
}

func NewGRPCServer(repo *repository.Repository) *GRPCServer {
	return &GRPCServer{repo: repo}
}

func (s *GRPCServer) Status(ctx context.Context, req *grpcProto.RequestStatus) (*grpcProto.ResponseStatus, error) {
	uid := int(req.Uid)
	status, err := s.repo.GetStatusByUid(uid)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	result, err := json.Marshal(status)
	if err != nil {
		log.Println(err)
	}
	return &grpcProto.ResponseStatus{
		Result: result,
	}, err
}
