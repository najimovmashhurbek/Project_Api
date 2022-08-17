package grpcClient

import (
	"github.com/najimovmashhurbek/Project_Api/post-service.first/config"
)

//GrpcClientI ...
type GrpcClientI interface {
}

//GrpcClient ...
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

//New ...
func New(cfg config.Config) (*GrpcClient, error) {
	return &GrpcClient{
		cfg:         cfg,
		connections: map[string]interface{}{},
	}, nil
}
