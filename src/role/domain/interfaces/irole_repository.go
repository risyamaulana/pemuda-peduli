package interfaces

import (
	"context"
	"pemuda-peduli/src/role/domain/entity"
)

type IRoleRepository interface {
	Insert(ctx context.Context, data *entity.RoleEntity) (err error)
	Update(ctx context.Context, data entity.RoleEntity, id string) (response entity.RoleEntity, err error)
	Find(ctx context.Context, data *entity.RoleQueryEntity) (response []entity.RoleEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.RoleEntity, err error)
}
