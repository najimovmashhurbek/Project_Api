package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/najimovmashhurbek/project-api/post-service.ozim/storage/postgres"
	"github.com/najimovmashhurbek/project-api/post-service.ozim/storage/repo"
)

//IStorage ...
type IStorage interface {
	Post() repo.PostStorageI
}

type storagePg struct {
	db       *sqlx.DB
	postRepo repo.PostStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		postRepo: postgres.NewUserRepo(db),
	}
}

func (s storagePg) Post() repo.PostStorageI {
	return s.postRepo
}
