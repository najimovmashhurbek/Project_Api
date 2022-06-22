package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/najimovmashhurbek/Project_Api/post-service.first/config"
	"github.com/najimovmashhurbek/Project_Api/post-service.first/pkg/db"
	"github.com/najimovmashhurbek/Project_Api/post-service.first/pkg/logger"
)

var repoI *postRepo

func TestMain(m *testing.M) {
	cfg := config.Load()
	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}
	repoI = NewUserRepo(connDB)
	os.Exit(m.Run())

}
