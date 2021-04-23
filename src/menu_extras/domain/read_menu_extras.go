package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/menu_extras/domain/entity"
	"pemuda-peduli/src/menu_extras/domain/interfaces"
)

func FindMenuExtras(ctx context.Context, repo interfaces.IMenuExtrasRepository, data *entity.MenuExtrasQueryEntity) (response []entity.MenuExtrasEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetMenuExtras(ctx context.Context, repo interfaces.IMenuExtrasRepository, id string) (response entity.MenuExtrasEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
