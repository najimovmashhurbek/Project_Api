package service

import (
	"context"
	_ "errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	pb "github.com/najimovmashhurbek/Project_Api/user-service.first/genproto"
	l "github.com/najimovmashhurbek/Project_Api/user-service.first/pkg/logger"
	"github.com/najimovmashhurbek/Project_Api/user-service.first/storage"

	//"golang.org/x/vuln/client"
	cl "github.com/najimovmashhurbek/Project_Api/user-service.first/service/grpc_client"
)

//UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	client  cl.GrpcClientI
}

//NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger, client cl.GrpcClientI) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	/*id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	req.Id = id.String()*/

	user, err := s.storage.User().CreateUser(req)
	if err != nil {
		return nil, err
	}
	// if req.Adress != nil {
	// 	for _, adres := range req.Adress {
	// 		adres.UserId = req.Id
	// 	}
	// }
	fmt.Println(req.Id)
	if req.Post != nil {
		for _, post := range req.Post {
			post.UserId = user.Id
			_, err := s.client.PostService().CreatePost(ctx, post)
			if err != nil {
				return nil, err
			}
		}
	}
	return user, err
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteById) (*pb.DeleteUserRes, error) {
	user, err := s.storage.User().DeleteUser(req)
	if err != nil {
		return nil, err
	}
	_, err = s.client.PostService().DeletePost(ctx, &pb.DeleteByPostId{
		Id: req.Id,
	})
	return user, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.User) (*pb.UpdateUserRes, error) {
	user, err := s.storage.User().UpdateUser(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *UserService) GetAllUser(ctx context.Context, req *pb.GetAllById) (*pb.User, error) {
	user, err := s.storage.User().GetAllUser(req)
	if err != nil {
		return nil, err
	}
	postS, err := s.client.PostService().Getallpost(ctx, &pb.GetAllByPostId{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	user.Post = postS.Response

	return user, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.GetUsersReq) (*pb.GetUsersRes, error) {
	users, count, err := s.storage.User().ListUsers(req.Limit, req.Page)
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		post, err := s.client.PostService().Getallpost(ctx, &pb.GetAllByPostId{Id: user.Id})
		if err != nil {
			return nil, err
		}
		user.Post = post.Response
	}

	return &pb.GetUsersRes{
		Users: users,
		Count: count,
	}, nil
}
func (s *UserService) CheckUniquess(ctx context.Context, req *pb.CheckUniqReq) (*pb.CheckUniqResp, error) {
	exists, err := s.storage.User().CheckUniquess(req.Field, req.Value)
	if err != nil {
		return nil, err
	}
	return &pb.CheckUniqResp{
		IsExist: exists,
	}, nil
}
