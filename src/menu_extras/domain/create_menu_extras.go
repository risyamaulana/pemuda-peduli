package domain

import (
	"context"
	"pemuda-peduli/src/menu_extras/domain/entity"
	"pemuda-peduli/src/menu_extras/domain/interfaces"
)

func CreateMenuExtras(ctx context.Context, repo interfaces.IMenuExtrasRepository, data *entity.MenuExtrasEntity) (err error) {
	err = repo.Insert(ctx, data)
	return
}
