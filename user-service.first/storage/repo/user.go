package repo

import (
	pb "github.com/najimovmashhurbek/Project_Api/user-service.first/genproto"
)

//UserStorageI ...
type UserStorageI interface {
	CreateUser(*pb.User) (*pb.User, error)
	DeleteUser(*pb.DeleteById) (*pb.DeleteUserRes, error)
	UpdateUser(*pb.User) (*pb.UpdateUserRes, error)
	GetAllUser(*pb.GetAllById) (*pb.User, error)
	ListUsers(limit, page int64) ([]*pb.User, int64, error)
	CheckUniquess(field, value string) (bool, error)
}
