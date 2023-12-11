package mapper_user

import (
	userEntity "rest-api/domain/entity/user"
	modeluser "rest-api/internal/repository/pgsql/user/model"
	"rest-api/pkg/common"
)

func ToModelUser(usr *userEntity.User) *modeluser.User {
	return &modeluser.User{
		Id:        usr.Id().String(),
		FirstName: usr.FirstName(),
		LastName:  usr.LastName(),
		Email:     usr.Email(),
	}
}

func ToModelUsers(usrs []*userEntity.User) []*modeluser.User {
	var users []*modeluser.User

	for _, usr := range usrs {
		users = append(users, ToModelUser(usr))
	}

	return users
}

func ToDomainUser(m *modeluser.User) (*userEntity.User, error) {
	id, err := common.StringToID(m.Id)
	if err != nil {
		return nil, err
	}
	user, errDto := userEntity.NewUser(&userEntity.UserDto{
		Id:        id,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Email:     m.Email,
	})
	if errDto != nil {
		return nil, errDto
	}

	return user, nil
}

func ToDomainUsers(m []*modeluser.User) ([]*userEntity.User, error) {
	var users []*userEntity.User

	for _, user := range m {
		usr, err := ToDomainUser(user)
		if err != nil {
			return nil, err
		}
		users = append(users, usr)
	}
	return users, nil
}
