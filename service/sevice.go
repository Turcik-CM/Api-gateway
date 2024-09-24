package service

import (
	pb "api-gateway/genproto/post"
	pbb "api-gateway/genproto/user"
	"api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	//pbp "api-gateway/genproto/tourism"
)

type Service interface {
	UserService() pbb.UserServiceClient
	PostService() pb.PostServiceClient
}

type service struct {
	userService pbb.UserServiceClient
	postService pb.PostServiceClient
}

func (s *service) UserService() pbb.UserServiceClient {
	return s.userService
}

func (s *service) PostService() pb.PostServiceClient {
	return s.postService
}

func NewService(cfg *config.Config) (Service, error) {
	userConn, err := grpc.NewClient("localhost:", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	postConn, err := grpc.NewClient("localhost:", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &service{
		userService: pbb.NewUserServiceClient(userConn),
		postService: pb.NewPostServiceClient(postConn),
	}, nil
}
