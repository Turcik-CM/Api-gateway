package service

import (
	pbp "api-gateway/genproto/nationality"
	pb "api-gateway/genproto/post"
	pbb "api-gateway/genproto/user"
	"api-gateway/pkg/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	//pbp "api-gateway/genproto/tourism"
)

type Service interface {
	UserService() pbb.UserServiceClient
	PostService() pb.PostServiceClient
	Nationality() pbp.NationalityServiceClient
}

type service struct {
	userService pbb.UserServiceClient
	postService pb.PostServiceClient
	nationality pbp.NationalityServiceClient
}

func (s *service) Nationality() pbp.NationalityServiceClient {
	return s.nationality
}

func (s *service) UserService() pbb.UserServiceClient {
	return s.userService
}

func (s *service) PostService() pb.PostServiceClient {
	return s.postService
}

func NewService(cfg *config.Config) (Service, error) {
	userConn, err := grpc.NewClient("localhost"+cfg.USER_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	postConn, err := grpc.NewClient("localhost"+cfg.POST_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Println("localhost" + cfg.POST_SERVICE)
	if err != nil {
		return nil, err
	}

	NationalityConn, err := grpc.NewClient("localhost:", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &service{
		userService: pbb.NewUserServiceClient(userConn),
		postService: pb.NewPostServiceClient(postConn),
		nationality: pbp.NewNationalityServiceClient(NationalityConn),
	}, nil
}
