package services

import (
	"fmt"

	"github.com/najimovmashhurbek/Project_Api/api-gateway_first/config"
	pb "github.com/najimovmashhurbek/Project_Api/api-gateway_first/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	UserService() pb.UserServiceClient
}

type serviceManager struct {
	userService pb.UserServiceClient
}

func (s *serviceManager) UserService() pb.UserServiceClient {
	return s.userService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		userService: pb.NewUserServiceClient(connUser),
	}

	return serviceManager, nil
}
