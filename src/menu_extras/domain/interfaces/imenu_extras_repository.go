package interfaces

import (
	"context"
	"pemuda-peduli/src/menu_extras/domain/entity"
)

type IMenuExtrasRepository interface {
	Insert(ctx context.Context, data *entity.MenuExtrasEntity) (err error)

	Update(ctx context.Context, data entity.MenuExtrasEntity, id string) (response entity.MenuExtrasEntity, err error)

	Find(ctx context.Context, data *entity.MenuExtrasQueryEntity) (company []entity.MenuExtrasEntity, count int, err error)
	Get(ctx context.Context, id string) (response entity.MenuExtrasEntity, err error)
}
