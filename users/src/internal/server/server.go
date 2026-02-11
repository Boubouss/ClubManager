package server

import (
	"ClubManager/users/internal/service"
	"ClubManager/users/internal/types"
	"ClubManager/shared/proto"
	"context"
	"net"

	"google.golang.org/grpc"
)

func MakeServerAndRun(addr string, svc service.UserService) error {
  userServer := NewUserServiceServer(svc)

  ln, err := net.Listen("tcp", addr)
  if err != nil {
    return err
  }

  opts := []grpc.ServerOption{}
  server := grpc.NewServer(opts...)

  proto.RegisterUserServiceServer(server, userServer)
  return server.Serve(ln)
}

type UserServiceServer struct {
  svc service.UserService
  proto.UnimplementedUserServiceServer
}

func NewUserServiceServer(svc service.UserService) *UserServiceServer {
  return &UserServiceServer{
    svc: svc,
  }
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *proto.CreateUserRequest) (*proto.CreateUserResponse, error) {
  user := &types.UserForm{
    Username: req.Username,
    Email: req.Email,
    Phonenumber: req.Phonenumber,
    Password: req.Password,
  }

  id, err := s.svc.CreateUser(ctx, user)
  if err != nil {
    return nil, err
  }

  res := &proto.CreateUserResponse{
    Id: id.String(),
  }

  return res, nil
}
