package postgres

import (
	"testing"

	"github.com/najimovmashhurbek/Project_Api/post-service.first/config"
	pb "github.com/najimovmashhurbek/Project_Api/post-service.first/genproto"
	"github.com/najimovmashhurbek/Project_Api/post-service.first/pkg/db"
	"github.com/najimovmashhurbek/Project_Api/post-service.first/storage/repo"
	"github.com/stretchr/testify/suite"
)

type PostRepositoryTestSuite struct {
	suite.Suite
	CleanupFunc func()
	Repository  repo.PostStorageI
}

func (suite *PostRepositoryTestSuite) SetupSuite() {
	pgPool, cleanup := db.ConnectDBForSuite(config.Load())

	suite.Repository = NewUserRepo(pgPool)
	suite.CleanupFunc = cleanup
}
func (suite *PostRepositoryTestSuite) TestCreateCRUD(t *testing.T) {
	postt := pb.Post{
		Id: "",
		Name: "suite",
		Description: "test suite",
		UserId: "",
	}
	//create test
	createdPost,err:=suite.Repository.CreatePost(postt.Name,postt.Description)
	suite.Nil(err)
	suite.NotNil(createdPost)

	//getbyid test
	getpost,err:=suite.Repository.Getallpost(createdPost.Id)
	suite.Nil(err)
	suite.NotNil(getpost)
	
	
	
}