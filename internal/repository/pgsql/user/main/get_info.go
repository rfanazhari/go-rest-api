package pgsql_user

import (
	"context"
	userEntity "rest-api/domain/entity/user"
	"rest-api/pkg/common"
)

func (p pgsqlUserMain) GetInfo(ctx context.Context, id *common.ID) (*userEntity.User, error) {
	//TODO implement me
	panic("implement me")
}
