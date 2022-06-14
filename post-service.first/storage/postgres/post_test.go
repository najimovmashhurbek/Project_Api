package postgres

import (
	"reflect"
	"testing"

	pb "github.com/najimovmashhurbek/project-api/post-service.ozim/genproto"
)

func TestPostRepo_Create(t *testing.T) {
	tests := []struct {
		name    string
		input   *pb.Post
		want    *pb.CreatePostsRes
		wantErr bool
	}{
		{
			name: "succescase",
			input: &pb.Post{
				Name: "hello",
				//Id : "0605077a-16c0-4e66-86ec-0c6aa7f0981f",
				UserId:      "0605077a-16c0-4e66-86ec-0c6aa7f0981f",
				Description: "HI",
				Medias:      nil,
			},
			want: &pb.CreatePostsRes{
				Status: true,
			},
			wantErr: false,
		}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			{
				got, err := repoI.CreatePost(tc.input)
				if err != nil {
					t.Fatalf("%s:expected : %v,got :%v", tc.name, tc.wantErr, err)
				}
				//got.Id = ""
				//got.UserId = ""
				//got.Medias = nil
				if !reflect.DeepEqual(tc.want, got) {
					t.Fatalf("%s:expected : %v,got :%v", tc.name, tc.want, got)
				}
			}
		})

	}
}
