package client

import (
	"ClubManager/shared/proto"

	"google.golang.org/grpc"
)

func NewUserServiceClient(target string) (proto.UserServiceClient, error) {
  conn, err := grpc.NewClient(target)
  if err != nil {
    return nil, err
  }

  return proto.NewUserServiceClient(conn), nil
}
