package domain

import (
	"context"
	"errors"
	"pemuda-peduli/src/role/domain/entity"
	"pemuda-peduli/src/role/domain/interfaces"
)

func FindRole(ctx context.Context, repo interfaces.IRoleRepository, data *entity.RoleQueryEntity) (response []entity.RoleEntity, count int, err error) {
	response, count, err = repo.Find(ctx, data)
	return
}

func GetRole(ctx context.Context, repo interfaces.IRoleRepository, id string) (response entity.RoleEntity, err error) {
	response, err = repo.Get(ctx, id)
	if err != nil {
		err = errors.New("Data not found")
		return
	}
	return
}
