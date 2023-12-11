package user_repository

import (
	"context"
	userEntity "rest-api/domain/entity/user"
	"rest-api/pkg/common"
)

type UserRepository interface {
	GetInfo(ctx context.Context, id *common.ID) (*userEntity.User, error)
}
