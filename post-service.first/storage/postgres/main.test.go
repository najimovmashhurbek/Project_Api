package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/najimovmashhurbek/project-api/post-service.ozim/config"
	"github.com/najimovmashhurbek/project-api/post-service.ozim/pkg/db"
	"github.com/najimovmashhurbek/project-api/post-service.ozim/pkg/logger"
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
