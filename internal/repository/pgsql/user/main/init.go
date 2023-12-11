package pgsql_user

import (
	"gorm.io/gorm"
	userrepository "rest-api/domain/repository/user"
)

type pgsqlUserMain struct {
	db *gorm.DB
}

func NewPgsqlUserMainRepository(db *gorm.DB) userrepository.UserRepository {
	return &pgsqlUserMain{db: db}
}
