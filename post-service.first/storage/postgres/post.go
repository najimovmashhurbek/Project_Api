package postgres

import (
	//"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	//"github.com/lib/pq"
	pb "github.com/najimovmashhurbek/project-api/post-service.ozim/genproto"
	//"google.golang.org/grpc/internal/status"
)

type postRepo struct {
	db *sqlx.DB
}

//NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}

func (r *postRepo) CreatePost(post *pb.Post) (*pb.CreatePostsRes, error) {
	query := "insert into posts (id,name,description,user_id) values ($1,$2,$3,$4)"
	id := uuid.New()
	_, err := r.db.Exec(query, id, post.Name, post.Description, post.UserId)
	if err != nil {
		return nil, err
	}
	//defer r.db.Close()
	for _, media := range post.Medias {
		adid := uuid.New()
		query1 := "insert into medias (id,link,posts_id) values ($1,$2,$3)"
		_, err := r.db.Exec(query1, adid, media.Link, id)
		if err != nil {
			return nil, err
		}
	}
	return &pb.CreatePostsRes{
		Status: true,
	}, nil
}

func (r *postRepo) DeletePost(delete *pb.DeleteByPostId) (*pb.DeletePostRes, error) {
	query := "DELETE FROM posts WHERE user_id=$1"
	_, err := r.db.Exec(query, delete.Id)
	if err != nil {
		return nil, err
	}
	//time1 := time.Now()
	return &pb.DeletePostRes{
		Status: true,
	}, nil
}
func (r *postRepo) UpdatePost(post *pb.Post) (*pb.UpdatePostRes, error) {
	query := "UPDATE posts set name=$1,description=$2 where id=$3"
	query1 := "UPDATE medias set link=$1 where id=$2"
	_, err := r.db.Exec(query, post.Name, post.Description, post.Id)
	if err != nil {
		return nil, err
	}
	for _, media := range post.Medias {
		_, err := r.db.Exec(query1, media.Link, post.Id)
		if err != nil {
			return nil, err
		}
	}
	return &pb.UpdatePostRes{Status: true}, nil

}

func (r *postRepo) Getallpost(get *pb.GetAllByPostId) ([]*pb.Post, error) {
	var res pb.Post
	var resp []*pb.Post
	query1 := "select id,name,description,user_id from posts where user_id=$1"
	query := "select id,link,posts_id from medias where posts_id=$1"
	rows, err := r.db.Query(query1, get.Id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(
			&res.Id,
			&res.Name,
			&res.Description,
			&res.UserId,
		)
		if err != nil {
			return nil, err
		}

		row, err := r.db.Query(query, res.Id)
		if err != nil {
			return nil, err
		}
		for row.Next() {
			var media pb.Medias
			err = row.Scan(
				&media.Id,
				&media.Link,
				&media.PostsId,
			)
			if err != nil {
				return nil, err
			}
			res.Medias = append(res.Medias, &media)
		}
		resp = append(resp, &res)
	}
	return resp, nil
}
func (r *postRepo) ListPosts(limit, page int64) ([]*pb.Post, int64, error) {
	var (
		posts []*pb.Post
		count int64
	)
	offset := (page - 1) * limit
	query := "select  id,name,description,user_id from posts order by name OFFSET $1 LIMIT $2"
	rows, err := r.db.Query(query, offset, limit)
	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		var post pb.Post
		err := rows.Scan(
			&post.Id,
			&post.Name,
			&post.Description,
			&post.UserId,
		)
		if err != nil {
			return nil, 0, err
		}
		posts = append(posts, &post)
	}

	return posts, count, nil
}
