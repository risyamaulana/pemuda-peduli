package interfaces

import (
	"context"
	"pemuda-peduli/src/admin_user/domain/entity"
)

type IAdminUserRepository interface {
	Insert(ctx context.Context, data *entity.AdminUserEntity) (err error)
	Update(ctx context.Context, data entity.AdminUserEntity, id string) (response entity.AdminUserEntity, err error)

	Find(ctx context.Context, data *entity.AdminUserQueryEntity) (response []entity.AdminUserEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.AdminUserEntity, err error)

	GetByUsername(ctx context.Context, username string) (response entity.AdminUserEntity, err error)
}
