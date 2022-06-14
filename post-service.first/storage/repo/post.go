package repo

import (
	pb "github.com/najimovmashhurbek/project-api/post-service.ozim/genproto"
)

//UserStorageI ...
type PostStorageI interface {
	CreatePost(*pb.Post) (*pb.CreatePostsRes, error)
	DeletePost(*pb.DeleteByPostId) (*pb.DeletePostRes, error)
	UpdatePost(*pb.Post) (*pb.UpdatePostRes, error)
	Getallpost(*pb.GetAllByPostId) ([]*pb.Post, error)
	ListPosts(limit, page int64) ([]*pb.Post, int64, error)
}
