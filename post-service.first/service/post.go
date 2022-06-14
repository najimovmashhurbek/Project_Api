package service

import (
	"context"

	"github.com/jmoiron/sqlx"
	pb "github.com/najimovmashhurbek/project-api/post-service.ozim/genproto"
	l "github.com/najimovmashhurbek/project-api/post-service.ozim/pkg/logger"
	"github.com/najimovmashhurbek/project-api/post-service.ozim/storage"
)

//UserService ...
type PostService struct {
	storage storage.IStorage
	logger  l.Logger
}

//NewUserService ...
func NewPostService(db *sqlx.DB, log l.Logger) *PostService {
	return &PostService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *PostService) CreatePost(ctx context.Context, req *pb.Post) (*pb.CreatePostsRes, error) {
	post, err := s.storage.Post().CreatePost(req)
	if err != nil {
		return nil, err
	}
	return post, err
}
func (s *PostService) DeletePost(ctx context.Context, req *pb.DeleteByPostId) (*pb.DeletePostRes, error) {
	post, err := s.storage.Post().DeletePost(req)
	if err != nil {
		return nil, err
	}
	return post, nil
}
func (s *PostService) UpdatePost(ctx context.Context, req *pb.Post) (*pb.UpdatePostRes, error) {
	post, err := s.storage.Post().UpdatePost(req)
	if err != nil {
		return nil, err
	}
	return post, nil
}
func (s *PostService) Getallpost(ctx context.Context, req *pb.GetAllByPostId) (*pb.Post1, error) {
	user, err := s.storage.Post().Getallpost(req)
	user1 := pb.Post1{}
	if err != nil {
		return nil, err
	}
	user1.Response = user
	return &user1, nil

}
func (s *PostService) ListPosts(ctx context.Context, req *pb.GetPostsReq) (*pb.GetPostsRes, error) {
	posts, count, err := s.storage.Post().ListPosts(req.Limit, req.Page)
	if err != nil {
		return nil, err
	}

	return &pb.GetPostsRes{
		Posts: posts,
		Count: count,
	}, nil
}
