package interfaces

import (
	"context"
	"pemuda-peduli/src/user/domain/entity"
)

type IUserRepository interface {
	Insert(ctx context.Context, data *entity.UserEntity) (err error)

	Update(ctx context.Context, data entity.UserEntity) (response entity.UserEntity, err error)

	Get(ctx context.Context, id string) (response entity.UserEntity, err error)

	GetForLogin(ctx context.Context, username string) (response entity.UserEntity, err error)

	GetDuplicateCheck(ctx context.Context, username, phoneNumber, email string) (response entity.UserEntity, err error)
}
