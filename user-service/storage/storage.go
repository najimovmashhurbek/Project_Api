package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/najimovmashhurbek/Project_Api/user-service.first/storage/postgres"
	"github.com/najimovmashhurbek/Project_Api/user-service.first/storage/repo"
)

//IStorage ...
type IStorage interface {
	User() repo.UserStorageI
}

type storagePg struct {
	db       *sqlx.DB
	userRepo repo.UserStorageI
}

//NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		userRepo: postgres.NewUserRepo(db),
	}
}

func (s storagePg) User() repo.UserStorageI {
	return s.userRepo
}
